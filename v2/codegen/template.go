package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	schemaLoc = "./apischema/auto-merge/"
)

func main() {
	schemas, _ := ioutil.ReadDir(schemaLoc)
	for _, schema := range schemas {
		fileName := schema.Name()
		log.Printf("[DEBUG] processing: %s", fileName)
		funcMap := template.FuncMap{
			"replace":     replace,
			"mixedCase":   mixedCase,
			"typeLookup":  typeLookup,
			"flatten":     flatten,
			"valilookup":  valilookup,
			"title":       title,
			"subcategory": subcategory,
			"difflookup":  diffLookup,
			"readExample": readExample,
			"debug":       debugx,
		}
		t := template.Must(template.New("main").Funcs(funcMap).ParseGlob("./templates/*.gotmpl"))

		m := map[string]interface{}{}
		jsondata := getSchema(fileName)
		if err := json.Unmarshal([]byte(jsondata), &m); err != nil {
			panic(err)
		}
		// Add relative path to json struct for use in templates.
		n := addPaths(m)
		// Add extra values to use in templates.
		r := addResourceInfo(n)

		render("models", fileName, t, r)
		render("cmdb", fileName, t, r)
		// d := addDataSourceInfo(n)

		// var buf2 bytes.Buffer
		// sname = "data_source"
		// if err := t.ExecuteTemplate(&buf2, sname, d); err != nil {
		// 	panic(err)
		// }

		// // lint
		// f, err = format.Source(buf2.Bytes())
		// if err != nil {
		// 	panic(err)
		// }

		// dsname := fname
		// dsname = sname + "_" + dsname
		// dsname = strings.TrimSuffix(dsname, ".json")
		// dsname += ".go"
		// os.WriteFile("./output/"+dsname, f, os.FileMode(perm))

		// var buf3 bytes.Buffer
		// sname = "data_source_docs"
		// if err := t.ExecuteTemplate(&buf3, sname, d); err != nil {
		// 	panic(err)
		// }

		// f = buf3.Bytes()

		// dsdocsname := fname
		// dsdocsname = "fortios_" + dsdocsname
		// dsdocsname = strings.TrimSuffix(dsdocsname, ".json")
		// dsdocsname += ".html.markdown"
		// os.WriteFile("./output/d/"+dsdocsname, f, os.FileMode(perm))

		// var buf4 bytes.Buffer
		// sname = "resource_docs"
		// if err := t.ExecuteTemplate(&buf4, sname, d); err != nil {
		// 	panic(err)
		// }

		// f = buf4.Bytes()

		// resdocsname := fname
		// resdocsname = "fortios_" + resdocsname
		// resdocsname = strings.TrimSuffix(resdocsname, ".json")
		// resdocsname += ".html.markdown"
		// os.WriteFile("./output/r/"+resdocsname, f, os.FileMode(perm))
	}
}

func debugx(v interface{}) string {
	fmt.Println(v)
	return ""
}

func render(templateName, fileName string, t *template.Template, m map[string]interface{}) {
	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, templateName, m); err != nil {
		panic(err)
	}

	// lint
	// f := buf.Bytes()
	f, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	// Debug
	// f := buf.Bytes()
	// fmt.Println(f)

	perm := int(0755)
	resname := strings.TrimSuffix(fileName, ".json")
	resname += ".go"
	os.WriteFile("../"+templateName+"/"+resname, f, os.FileMode(perm))
}

func addPaths(m map[string]interface{}) map[string]interface{} {
	path := mixedCase(m["path"].(string))
	if _, ok := m["child_path"].(string); ok {
		name, _ := m["name"].(string)
		path += name
	}
	resource := mixedCase(m["results"].(map[string]interface{})["name"].(string))
	tmp := m["results"]
	m["fpath"] = path + resource
	m["results"].(map[string]interface{})["fpath"] = path + resource
	recursePaths(tmp, path+resource)
	return m
}

func recursePaths(v interface{}, pre string) interface{} {
	child, ok := v.(map[string]interface{})["children"].(map[string]interface{})
	if ok {
		for _, v := range child {
			if v2, ok := v.(map[string]interface{}); ok {
				name := mixedCase(v2["name"].(string))
				v2["fpath"] = pre + name
				recursePaths(v, pre+name)
			}
		}
	} else {
		return v
	}
	return v
}

func addSensitive(m map[string]interface{}) map[string]interface{} {
	tmp := m["results"].(map[string]interface{})
	recurseSensitive(tmp)
	return m
}

func recurseSensitive(m map[string]interface{}) map[string]interface{} {
	if child, ok := m["children"].(map[string]interface{}); ok {
		for _, v := range child {
			if v2, ok := v.(map[string]interface{}); ok {
				if s, ok := v2["type"].(string); ok {
					if strings.Contains(s, "password") {
						v2["sensitive"] = true
					}
				}
				recurseSensitive(v2)
			}
		}
	} else {
		return m
	}
	return m
}

// mark if name begins with number cause Go struct no likey and need to prepend
func addNumberName(m map[string]interface{}) map[string]interface{} {
	tmp := m["results"].(map[string]interface{})
	recurseNumberName(tmp)
	return m
}

func recurseNumberName(m map[string]interface{}) map[string]interface{} {
	if child, ok := m["children"].(map[string]interface{}); ok {
		for _, v := range child {
			if v2, ok := v.(map[string]interface{}); ok {
				if s, ok := v2["name"].(string); ok {
					if s[0] >= '0' && s[0] <= '9' {
						v2["name_begin_with_int"] = true
					}
				}
				recurseNumberName(v2)
			}
		}
	} else {
		return m
	}
	return m
}

func addDataSourceInfo(m map[string]interface{}) map[string]interface{} {
	m = replaceTopLevelId(m)
	m = addDataSourceSchemaRequired(m)
	m = addSensitive(m)
	m = addNumberName(m)
	return m
}

func addDataSourceSchemaRequired(m map[string]interface{}) map[string]interface{} {
	mkey, ok := m["results"].(map[string]interface{})["mkey"].(string)
	if !ok {
		return m
	}
	child, ok := m["results"].(map[string]interface{})["children"].(map[string]interface{})
	if ok {
		for _, v := range child {
			name := v.(map[string]interface{})["name"].(string)
			if name == mkey {
				v.(map[string]interface{})["schema_required"] = true
			}

		}
	}
	return m
}

func addResourceSchemaRequired(m map[string]interface{}) map[string]interface{} {
	mkey, ok := m["results"].(map[string]interface{})["mkey"].(string)
	if !ok {
		return m
	}
	child, ok := m["results"].(map[string]interface{})["children"].(map[string]interface{})
	if ok {
		for _, v := range child {
			name := v.(map[string]interface{})["name"].(string)
			if name == "seq-num" || name == "id" {
				v.(map[string]interface{})["schema_opt_force"] = true
			} else if name == mkey {
				v.(map[string]interface{})["schema_required"] = true
			}

		}
	}
	return m
}

func addResourceInfo(m map[string]interface{}) map[string]interface{} {
	m = addDynSortTable(m)
	m = replaceTopLevelId(m)
	m = addResourceSchemaRequired(m)
	m = addSensitive(m)
	m = complexEmptyOnDestroy(m)
	m = addNumberName(m)
	return m
}

func replaceTopLevelId(m map[string]interface{}) map[string]interface{} {
	child, ok := m["results"].(map[string]interface{})["children"].(map[string]interface{})
	if ok {
		for _, v := range child {
			name := v.(map[string]interface{})["name"].(string)
			if name == "id" {
				v.(map[string]interface{})["fosid"] = true
				v.(map[string]interface{})["name"] = "fosid"
			}

		}
	}
	return m
}

func addDynSortTable(m map[string]interface{}) map[string]interface{} {
	child, ok := m["results"].(map[string]interface{})["children"].(map[string]interface{})
	if ok {
		dyn := false
		for _, v := range child {
			if category, ok := v.(map[string]interface{})["category"].(string); ok {
				if category == "table" {
					dyn = true
				}
			}
			m["results"].(map[string]interface{})["dynamic_sort_table"] = dyn
		}
	}
	return m
}

func complexEmptyOnDestroy(m map[string]interface{}) map[string]interface{} {
	category := m["results"].(map[string]interface{})["category"].(string)
	if category == "complex" {
		if child, ok := m["results"].(map[string]interface{})["children"].(map[string]interface{}); ok {
			complex_parent := false
			for _, v := range child {
				if category, ok := v.(map[string]interface{})["category"].(string); ok {
					if category == "table" {
						complex_parent = true
					}
				}
				m["results"].(map[string]interface{})["complex_parent"] = complex_parent
			}
		}
	}
	return m
}

func getSchema(name string) (content []byte) {
	file := schemaLoc
	file += name
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func flatten(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, ".", "")
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
	}
	s, ok := m[s]
	if !ok {
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
	if s == `"enable", "disable"` || s == `"disable", "enable"` {
		return "fortiValidateEnableDisable()"
	}
	return fmt.Sprintf("fortiValidateEnum([]string{%s})", s)
}

func mixedCase(v string) string {
	v = strings.ReplaceAll(v, ".", "")
	v = strings.ReplaceAll(v, "-", " ")
	v = strings.ReplaceAll(v, "+", "")
	v = strings.ReplaceAll(v, "(", "")
	v = strings.ReplaceAll(v, ")", "")
	return strings.ReplaceAll(strings.Title(v), " ", "")
}

func title(v string) string {
	return strings.Title(v)
}

func subcategory(input string) string {
	s := strings.Split(input, ".")
	return s[0]
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

// Entire swagger schema
// https://fndn.fortinet.net/index.php?/fortiapi/1-fortios/&do=downloadApiJSONFile&handlerId=974&path=&productId=1&csrfKey=c10ec59a59015e6c2538adec6796f1ca
