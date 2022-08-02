package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/tools/imports"
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

type t struct {
	name   string
	prefix string
	suffix string
}

var renderedTemplates = []t{
	{
		name:   "models",
		prefix: "../../models/",
		suffix: ".go",
	},
	{
		name:   "cmdb",
		prefix: "../../cmdb/",
		suffix: ".go",
	},
}

// Map of templating functions
var tfuncs = template.FuncMap{
	"mixedCase":  mixedCase,
	"replace":    replace,
	"typeLookup": typeLookup,
	// "flatten":     flatten,
	"valilookup":  valilookup,
	"title":       title,
	"subcategory": subcategory,
	"difflookup":  diffLookup,
	"readExample": readExample,
	"debug":       debugx,
	"deref":       deref,
	"urlPath":     urlPath,
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

type cmdbRes struct {
	Path     string
	Complex  bool
	Category string
}

func renderTemplate(templateName, outputPath string, t *template.Template, data interface{}) {
	log.Println(outputPath)
	// create output file
	outputFile := filepath.Join(outputPath)
	os.MkdirAll(filepath.Dir(outputFile), 0755)
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer f.Close()

	output := new(bytes.Buffer)
	err = t.ExecuteTemplate(output, templateName, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// lint
	output2, err := format.Source(output.Bytes())
	if err != nil {
		f.Write(output.Bytes())
		fmt.Println(outputPath)
		panic(err)
	}
	output2, err = imports.Process("", output2, nil)
	if err != nil {
		panic(err)
	}

	f.Write(output2)

}

func main() {
	f, err := os.ReadFile("./merged.json")
	if err != nil {
		log.Fatal(err)
	}

	schema := &FullSchema{}
	err = json.Unmarshal(f, schema)
	if err != nil {
		log.Fatal(err)
	}

	schema.augmentSchema()

	// load templates
	templates := template.Must(template.New("main").Funcs(tfuncs).ParseGlob("./templates/*.gotmpl"))

	cmdbResources := []cmdbRes{}

	for _, v := range schema.Results {
		// exclude specific paths
		if slices.Contains(excludedFullPaths, mixedCase(&v.Schema.FullPath)) {
			continue
		}
		// ignore test paths
		if *v.Schema.Path == "test" {
			continue
		}

		for _, o := range renderedTemplates {
			renderTemplate(o.name, fmt.Sprintf("%s%s%s", o.prefix, flatten(v.Schema.FullPath), o.suffix), templates, v.Schema)
		}

		for _, o2 := range v.Schema.Children {
			if deref(o2.Category) == "table" && slices.Contains(includedFullPaths, mixedCase(&o2.FullPath)) {
				renderTemplate("cmdb", fmt.Sprintf("%s%s%s", "../../cmdb/", flatten(o2.FullPath), ".go"), templates, o2)
			}
		}

		tmp := cmdbRes{
			Path:     v.Schema.FullPath,
			Complex:  isComplex(*v.Schema.Category),
			Category: *v.Schema.Category,
		}
		cmdbResources = append(cmdbResources, tmp)
	}
	renderTemplate("cmdbClient", "../../cmdb/cmdb.go", templates, cmdbResources)

}

func isComplex(category string) bool {
	return category == "complex"
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

/////////////////////////////////

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func flatten(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	// s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, " ", "_")
	return s
}

func urlPath(sp *string) string {
	if sp == nil {
		return ""
	}
	s := strings.ToLower(*sp)
	s = strings.ReplaceAll(s, " ", "/")
	return s
}

func typeLookup(s string) string {
	m := map[string]string{
		"string":                 "*string",
		"option":                 "*string",
		"ipv4-address":           "*string",
		"ipv4-address-any":       "*string",
		"ipv4-classnet":          "*string",
		"ipv4-classnet-host":     "*string",
		"ipv4-classnet-any":      "*string",
		"ipv4-netmask":           "*string",
		"ipv4-netmask-any":       "*string",
		"ipv4-address-multicast": "*string",
		"ipv6-address":           "*string",
		"ipv6-prefix":            "*string",
		"ipv6-network":           "*string",
		"var-string":             "*string",
		"password":               "*string",
		"integer":                "*int64",
		"user":                   "*string",
		"password-2":             "*string",
		"password-3":             "*string",
		"varlen_password":        "*string",
		"mac-address":            "*string",
		"password_aes256":        "*string",
		"uuid":                   "*string",
		"ether-type":             "*string",
		"datatime":               "*string",
	}
	s, ok := m[s]
	if !ok {
		// s = "FIXMEEEEEE"
		// return s
		log.Fatalf("Type not found %s", s)
	}
	return s
}

func valilookup(values map[string]interface{}) string {
	vtype := values["type"].(string)
	size, _ := values["size"].(float64)
	multi_val, _ := values["multiple_values"].(bool)
	o, _ := values["options"].([]interface{})
	min, _ := values["min-value"].(float64)
	max, _ := values["max-value"].(float64)
	m := map[string]string{
		"string":             valiStringLen(size),
		"option":             valiOptions(o, multi_val),
		"ipv4-address":       "validation.IsIPv4Address",
		"ipv4-address-any":   "validation.IsIPv4Address",
		"ipv4-classnet":      "fortiValidateIPv4Classnet",
		"ipv4-classnet-host": "fortiValidateIPv4ClassnetHost",
		"ipv4-classnet-any":  "fortiValidateIPv4ClassnetAny",
		"ipv4-netmask":       "fortiValidateIPv4Netmask",
		"ipv6-address":       "validation.IsIPv6Address",
		"ipv6-prefix":        "fortiValidateIPv6Prefix",
		"ipv6-network":       "fortiValidateIPv6Network",
		"var-string":         valiStringLen(size),
		"password":           "",
		"integer":            valiInt(int(min), int(max)),
		"user":               "",
		"password-3":         "",
		"mac-address":        "",
		"password_aes256":    "",
		"uuid":               "",
		"ether-type":         "",
	}
	s, ok := m[vtype]
	if !ok {
		s = "VALI-ERROR"
	}
	if s == "" {
		return ""
	} else {
		return fmt.Sprintf("ValidateFunc: %s,", s)
	}
}

func diffLookup(values map[string]interface{}) string {
	vtype := values["type"].(string)
	multi_val, _ := values["multiple_values"].(bool)
	m := map[string]string{
		"string":             "",
		"option":             diffOptions(multi_val),
		"ipv4-address":       "",
		"ipv4-address-any":   "",
		"ipv4-classnet":      "",
		"ipv4-classnet-host": "",
		"ipv4-classnet-any":  "",
		"ipv4-netmask":       "",
		"ipv6-address":       "diffIPEqual",
		"ipv6-prefix":        "diffCidrEqual",
		"ipv6-network":       "diffCidrEqual",
		"var-string":         "",
		"password":           "",
		"integer":            "",
		"user":               "",
		"password-3":         "",
		"mac-address":        "",
		"password_aes256":    "",
		"uuid":               "",
		"ether-type":         "",
	}
	s, ok := m[vtype]
	if !ok {
		s = "DIFF-LOOKUP-ERROR"
	}
	if s == "" {
		return ""
	} else {
		return fmt.Sprintf("DiffSuppressFunc: %s,", s)
	}
}

func diffOptions(multi_val bool) string {
	if multi_val {
		return "diffFakeListEqual"
	} else {
		return ""
	}
}

func valiStringLen(v float64) string {
	i := int(v)
	return fmt.Sprintf("validation.StringLenBetween(0, %d)", i)
}

func valiInt(min, max int) string {
	if max > 2147483647 {
		return ""
	}
	return fmt.Sprintf("validation.IntBetween(%d, %d)", min, max)
}

func valiOptions(opts []interface{}, multi_val bool) string {
	if multi_val {
		return ""
	}
	l := make([]string, 0)
	for _, v := range opts {
		l = append(l, fmt.Sprintf("%q", v.(map[string]interface{})["name"].(string)))
	}
	s := strings.Join(l, ", ")
	return fmt.Sprintf("validation.StringInSlice([]string{%s})", s)
}

func title(v string) string {
	return strings.Title(v)
}

func subcategory(input string) string {
	s := strings.Split(input, ".")
	return s[0]
}

func debugx(v interface{}) string {
	fmt.Println(v)
	return ""
}

func readExample(name, typ string) string {
	file, err := os.Open(fmt.Sprintf("./examples/%s.%s.txt", name, typ))
	if err != nil {
		return ""
	}
	example, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(example)
}
