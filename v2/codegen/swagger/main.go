package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	d, _ := os.ReadDir("./cmdb")
	var dir []string
	for _, v := range d {
		if strings.HasSuffix(v.Name(), ".json") {
			dir = append(dir, v.Name())
		}
	}
	for _, v := range dir {
		getPaths(v)
	}
}

func getPaths(fileName string) {
	b, err := os.ReadFile("./cmdb/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	var j1 interface{}
	err = json.Unmarshal(b, &j1)
	if err != nil {
		log.Fatal(err)
	}

	tmp := j1.(map[string]interface{})["tags"].([]interface{})
	paths := []string{}
	for _, v := range tmp {
		p := v.(map[string]interface{})["name"].(string)
		paths = append(paths, p)
	}

	data, err := json.MarshalIndent(paths, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(fmt.Sprintf("./paths/%s", fileName), data, os.FileMode(int(0755)))
}
