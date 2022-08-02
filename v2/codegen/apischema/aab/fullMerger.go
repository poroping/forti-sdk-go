package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/imdario/mergo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FullSchema struct {
	HTTPMethod *string  `json:"http_method,omitempty"`
	Revision   *string  `json:"revision,omitempty"`
	Vdom       *string  `json:"vdom,omitempty"`
	Action     *string  `json:"action,omitempty"`
	Status     *string  `json:"status,omitempty"`
	HTTPStatus *int64   `json:"http_status,omitempty"`
	Serial     *string  `json:"serial,omitempty"`
	Version    *string  `json:"version,omitempty"`
	Build      *int64   `json:"build,omitempty"`
	Results    []*Paths `json:"results,omitempty"`
}

func (o *FullSchema) augmentSchema() {
	for _, p := range o.Results {
		addFullPath(p.Schema)
		complexParent(p.Schema)
		dynamicTable(p.Schema)
		for _, v := range p.Schema.Children {
			attributeSensitive(v)
			nameBeginWithInt(v)
			v.setGoName()
		}
	}
}

type Paths struct {
	Path   *string `json:"path,omitempty"`
	Name   *string `json:"name,omitempty"`
	Schema *Schema `json:"schema,omitempty"`
}

type Schema struct {
	Name                *string                      `json:"name,omitempty"`
	Category            *string                      `json:"category,omitempty"`
	Help                *string                      `json:"help,omitempty"`
	Mkey                *string                      `json:"mkey,omitempty"`
	MkeyType            *string                      `json:"mkey_type,omitempty"`
	MaxTableSizeVdom    *int64                       `json:"max_table_size_vdom,omitempty"`
	MaxTableSizeGlobal  *int64                       `json:"max_table_size_global,omitempty"`
	MaxTableSizeItem    *int64                       `json:"max_table_size_item,omitempty"`
	MemberTable         *bool                        `json:"member_table,omitempty"`
	Path                *string                      `json:"path,omitempty"`
	ReadOnly            *bool                        `json:"readonly,omitempty"`
	QType               *int64                       `json:"q_type,omitempty"`
	AccessGroup         *string                      `json:"access_group,omitempty"`
	Children            map[string]*SchemaAttributes `json:"children,omitempty"`
	ComplexParent       bool                         `json:"-"`
	DynamicSortSubtable bool                         `json:"-"`
	FullPath            string                       `json:"-"`
	AddedIn             string                       `json:"added_in"`
	RemovedIn           string                       `json:"removed_in"`
}

func (o *Schema) setComplexParent() {
	o.ComplexParent = true
}

func (o *Schema) setDynamicSortSubtable() {
	o.DynamicSortSubtable = true
}

type SchemaAttributes struct {
	Name                     *string                      `json:"name,omitempty"`
	Category                 *string                      `json:"category,omitempty"`
	Type                     *string                      `json:"type,omitempty"`
	Help                     *string                      `json:"help,omitempty"`
	Mkey                     *string                      `json:"mkey,omitempty"`
	MkeyType                 *string                      `json:"mkey_type,omitempty"`
	MinValue                 *int64                       `json:"min-value,omitempty"`
	MaxValue                 *int64                       `json:"max-value,omitempty"`
	Default                  interface{}                  `json:"default,omitempty"`
	MultipleValues           *bool                        `json:"multiple_values,omitempty"`
	Options                  []SchemaAttributesOptions    `json:"options,omitempty"`
	Children                 map[string]*SchemaAttributes `json:"children,omitempty"`
	Datasource               []string                     `json:"datasource,omitempty"`
	Size                     *int64                       `json:"size,omitempty"`
	ReadOnly                 *bool                        `json:"readonly,omitempty"`
	FullPath                 string                       `json:"-"`
	GoName                   string                       `json:"-"`
	Sensitive                bool                         `json:"-"`
	NameBeginWithInt         bool                         `json:"-"`
	NameReserved             bool                         `json:"-"`
	DataSourceSchemaRequired bool                         `json:"-"`
	ResourceSchemaRequired   bool                         `json:"-"`
	ResourceForceNew         bool                         `json:"-"`
	AddedIn                  string                       `json:"added_in"`
	RemovedIn                string                       `json:"removed_in"`
}

func (o *SchemaAttributes) setSensitive() {
	o.Sensitive = true
}

func (o *SchemaAttributes) setNameBeginWithInt() {
	o.NameBeginWithInt = true
}

func (o *SchemaAttributes) setFullPath(s string) {
	o.FullPath = s
}

func (o *SchemaAttributes) setGoName() {
	g := o.GoName
	o.GoName = mixedCase(&g)
}

type SchemaAttributesOptions struct {
	Name *string `json:"name,omitempty"`
	Help *string `json:"help,omitempty"`
}

const (
	templateDir = "../templates"
)

// These paths are dumb/not used/troublesome
var excludedFullPaths = []string{
	"SystemPerformanceStatus",
	"SystemPerformanceTop",
	"WirelessControllerStatus",
	"WirelessControllerVapStatus",
	"WirelessControllerSpectralInfo",
	"FirewallIprope",
	"SystemSessionHelperInfoList",
	"VpnIkeGateway",
}

// Explicitly included sub tables
var includedFullPaths = []string{
	"RouterBgpNeighborAdminDistance",
	"RouterBgpNeighborAggregateAddress",
	"RouterBgpNeighborAggregateAddress6",
	"RouterBgpNeighborGroup",
	"RouterBgpNeighborRange",
	"RouterBgpNeighbor",
	"RouterBgpNetwork",
	"RouterBgpNetwork6",
	"RouterBgpRedistribute",
	"RouterBgpRedistribute6",
}

func mixedCase(vp *string) string {
	if vp == nil {
		return ""
	}
	c := cases.Title(language.English)
	v := *vp
	v = strings.ReplaceAll(v, ".", " ")
	v = strings.ReplaceAll(v, "-", " ")
	v = strings.ReplaceAll(v, "+", "")
	v = strings.ReplaceAll(v, "<", "")
	v = strings.ReplaceAll(v, ">", "")
	v = strings.ReplaceAll(v, "(", "")
	v = strings.ReplaceAll(v, ")", "")
	return strings.ReplaceAll(c.String(v), " ", "")
}

// Mark if parent is complex type
func complexParent(s *Schema) {
	if category := s.Category; category != nil {
		if *category == "complex" {
			for _, v := range s.Children {
				if cat := v.Category; cat != nil {
					if *cat == "table" {
						s.setComplexParent()
					}
				}
			}
		}
	}
}

// Mark if table can be sorted
func dynamicTable(s *Schema) {
	for _, v := range s.Children {
		if cat := v.Category; cat != nil {
			if *cat == "table" {
				s.setDynamicSortSubtable()
			}
		}
	}
}

// Mark if name begins with number cause Go struct no likey and need to prepend
func nameBeginWithInt(s *SchemaAttributes) {
	name := *s.Name
	if name[0] >= '0' && name[0] <= '9' {
		s.setNameBeginWithInt()
	}
	for key := range s.Children {
		nameBeginWithInt(s.Children[key])
	}
}

// Mark if the attribute is considered sensitive
func attributeSensitive(s *SchemaAttributes) {
	if t := s.Type; t != nil {
		if strings.Contains(*t, "password") {
			s.setSensitive()
		}
	}
	for key := range s.Children {
		attributeSensitive(s.Children[key])
	}
}

// Add the full path to the schema
func addFullPath(s *Schema) {
	path := *s.Path
	resource := *s.Name
	pre := path + " " + resource
	s.FullPath = pre
	for key := range s.Children {
		addFullPathSchemaAttributes(s.Children[key], pre)
	}
}

// Add the full path to the schema attributes
func addFullPathSchemaAttributes(s *SchemaAttributes, pre string) {
	fp := pre + " " + *s.Name
	s.setFullPath(fp)
	for key := range s.Children {
		addFullPathSchemaAttributes(s.Children[key], fp)
	}
}

func main() {
	// paths := getPaths("../paths")
	versions := getVersions()
	fmt.Println(versions)

	files := getFiles(versions)

	final := mergeFiles(files)

	data, err := json.MarshalIndent(final, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(fmt.Sprintf("../full/merged.json"), data, os.FileMode(int(0755)))
}

func getPaths(directory string) []string {
	d, _ := os.ReadDir(directory)
	var dir []string
	for _, v := range d {
		if strings.HasSuffix(v.Name(), ".json") {
			dir = append(dir, v.Name())
		}
	}

	paths := []string{}
	fmt.Println(dir)
	for _, file := range dir {
		data, err := os.ReadFile(fmt.Sprintf("%s/%s", directory, file))
		if err != nil {
			fmt.Print(err)
		}
		tmp := &[]string{}
		err = json.Unmarshal(data, tmp)
		if err != nil {
			fmt.Print(err)
		}
		paths = append(paths, *tmp...)
	}

	return removeDuplicateStr(paths)
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func mergeFiles(files []FullSchema) *FullSchema {
	merged := &FullSchema{}
	var err error
	for i := range files {
		if len(files) == 1 {
			merged = &files[i]
			continue
		}
		if i == 0 {
			new := files[i+1]
			old := files[i]
			merged, err = mergeFullSchema(&new, &old)
			if err != nil {
				log.Fatal(err)
			}
			continue
		}
		if i == len(files)-1 {
			break
		}
		new := files[i+1]
		old := *merged
		merged, err = mergeFullSchema(&new, &old)
		if err != nil {
			log.Fatal(err)
		}
	}
	return merged
}

func mergeFullSchema(new, old *FullSchema) (*FullSchema, error) {
	versionNew := *new.Version
	versionOld := *old.Version

	log.Println(*new.Version, len(new.Results), *old.Version, len(old.Results))

	pathMapNew := make(map[string]*Paths)
	pathMapOld := make(map[string]*Paths)

	for _, v := range new.Results {
		path := v.Path
		name := v.Name
		k := *path + *name
		pathMapNew[k] = v
	}

	for _, v := range old.Results {
		// if v == nil {
		// 	continue
		// }
		path := v.Path
		name := v.Name
		k := *path + *name
		pathMapOld[k] = v
	}

	pathsMerged, _ := mergePaths(pathMapNew, pathMapOld, versionNew, versionOld)

	pathsUnmerged := make([]*Paths, 0)
	for _, v := range pathsMerged {
		pathsUnmerged = append(pathsUnmerged, v)
	}

	new.Results = pathsUnmerged

	return new, nil
}

func mergePaths(new, old map[string]*Paths, versionNew, versionOld string) (map[string]*Paths, error) {
	for oldKey, oldValue := range old {
		if _, ok := new[oldKey]; ok {
			new[oldKey].Schema = mergeSchema(new[oldKey].Schema, old[oldKey].Schema, versionNew, versionOld)
		} else {
			new[oldKey] = oldValue
			log.Println("old path removed:", oldKey)
			if oldValue.Schema.RemovedIn != "" {
				new[oldKey].Schema.RemovedIn = versionNew
			}
		}
	}

	for newKey := range new {
		if _, ok := old[newKey]; !ok {
			if new[newKey].Schema.AddedIn == "" {
				log.Println("new key added:", newKey)
				new[newKey].Schema.AddedIn = versionNew
			}
		}
	}

	return new, nil

}

func mergeSchema(new, old *Schema, versionNew, versionOld string) *Schema {
	tmp, err := mergeSchemaAttributes(new.Children, old.Children, versionNew, versionOld)
	if err != nil {
		log.Fatal(err)
	}
	new.Children = tmp
	return new
}

// merge merges the two JSON-marshalable values new and old,
// preferring new over old except where new and old are
// JSON objects, in which case the keys from both objects
// are included and their values merged recursively.
//

func mergeSchemaAttributes(new, old map[string]*SchemaAttributes, versionNew, versionOld string) (map[string]*SchemaAttributes, error) {
	for old_key, v2 := range old {
		if v1, ok := new[old_key]; ok {
			tmp := mergeSchemaAttribute(v1, v2)
			new[old_key] = tmp
		} else { // removed in new schema
			new[old_key] = v2
			if m, ok := new[old_key]; ok {
				if m.RemovedIn != "" {
					m.RemovedIn = versionNew
				}

			}
		}
	}
	for new_key, v2 := range new {
		if _, ok := old[new_key]; ok {
		} else { // added in new schema
			new[new_key] = v2
			if m, ok := new[new_key]; ok {
				if m.AddedIn != "" {
					if new_key != "children" { // for category changes, deal with this better later.
						m.AddedIn = versionNew
					}
				}
			}
		}
	}
	return new, nil
}

func mergeSchemaAttribute(new, old *SchemaAttributes) *SchemaAttributes {
	err := mergo.Merge(old, new, mergo.WithOverride)
	if err != nil {
		log.Fatal(err)
	}

	return old
}

func getVersions() []string {
	d, _ := os.ReadDir("../full")
	var dir []string
	for _, v := range d {
		if !v.IsDir() {
			if strings.HasPrefix(v.Name(), "v") {
				dir = append(dir, v.Name())
				log.Println("found:", v.Name())
			}
		}
	}
	return dir
}

func getFiles(dir []string) []FullSchema {
	files := []FullSchema{}
	for _, v := range dir {
		filePath := fmt.Sprintf("../full/%s", v)
		x1, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("[INFO] resource %s not present", filePath)
		}
		j1 := &FullSchema{}
		err = json.Unmarshal(x1, j1)
		if err != nil {
			log.Fatal(err)
		}
		files = append(files, *j1)
	}
	return files
}
