package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	paths := getPaths("../paths")
	versions := getVersions()
	fmt.Println(versions)

	for _, path := range paths {
		path = strings.ReplaceAll(path, "-", "_")
		path = strings.ReplaceAll(path, "/", "_")
		files, resRemovedIn, resAddedIn, mergedVersions := getFiles(versions, path)

		final := mergeFiles(files)
		if final, ok := final.(map[string]interface{}); ok {
			final["mergedVersions"] = strings.Join(mergedVersions, ",")
			if resAddedIn != "" {
				final["addedIn"] = resAddedIn
				log.Printf("[INFO] resource %s flagged as added in version: %s", path, resAddedIn)
			}
			if resRemovedIn != "" {
				final["removedIn"] = resRemovedIn
				log.Printf("[INFO] resource %s flagged as removed in version: %s", path, resRemovedIn)
			}

		}

		data, err := json.MarshalIndent(final, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		os.WriteFile(fmt.Sprintf("../auto-merge/%s.json", path), data, os.FileMode(int(0755)))
	}
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

func mergeFiles(files []interface{}) interface{} {
	var merged *interface{}
	var err error
	for i := range files {
		if len(files) == 1 {
			merged = &files[i]
			continue
		}
		if i == 0 {
			new := files[i+1]
			old := files[i]
			versionNew := new.(map[string]interface{})["version"].(string)
			versionOld := old.(map[string]interface{})["version"].(string)
			merged, err = merge(&new, &old, versionNew, versionOld)
			if err != nil {
				log.Fatal(err)
			}
		}
		if i == len(files)-1 {
			break
		}
		new := files[i+1]
		old := *merged
		versionNew := new.(map[string]interface{})["version"].(string)
		versionOld := old.(map[string]interface{})["version"].(string)
		merged, err = merge(&new, &old, versionNew, versionOld)
		if err != nil {
			log.Fatal(err)
		}
	}
	return *merged
}

// merge merges the two JSON-marshalable values new and old,
// preferring new over old except where new and old are
// JSON objects, in which case the keys from both objects
// are included and their values merged recursively.
//

func merge(newp, oldp *interface{}, versionNew, versionOld string) (*interface{}, error) {
	new := *newp
	switch new := new.(type) {
	case map[string]interface{}:
		tmp := *oldp
		old, ok := tmp.(map[string]interface{})
		if !ok { // not json obj then prefer new value
			return newp, nil
		}
		for old_key, v2 := range old {
			if v1, ok := new[old_key]; ok {
				tmp, _ := merge(&v1, &v2, versionNew, versionOld)
				new[old_key] = *tmp
			} else { // removed in new schema
				new[old_key] = v2
				if m, ok := new[old_key].(map[string]interface{}); ok {
					if _, exists := m["removedIn"].(string); !exists {
						m["removedIn"] = versionNew
					}

				}
			}
		}
		for new_key, v2 := range new {
			if _, ok := old[new_key]; ok {
			} else { // added in new schema
				old[new_key] = v2
				if m, ok := old[new_key].(map[string]interface{}); ok {
					if _, exists := m["addedIn"].(string); !exists {
						if new_key != "children" { // for category changes, deal with this better later.
							m["addedIn"] = versionNew
						}
					}
				}
			}
		}
	case nil:
		tmp := *oldp
		_, ok := tmp.(map[string]interface{})
		if ok {
			return oldp, nil
		}
	}
	return &new, nil
}

func getVersions() []string {
	d, _ := os.ReadDir("..")
	var dir []string
	for _, v := range d {
		if v.IsDir() {
			if strings.HasPrefix(v.Name(), "v") {
				dir = append(dir, v.Name())
			}
		}
	}
	return dir
}

func getFiles(dir []string, path string) ([]interface{}, string, string, []string) {
	files := make([]interface{}, 0)
	mergedVersions := []string{}
	resRemovedIn := ""
	resAddedIn := ""
	for _, v := range dir {
		filePath := fmt.Sprintf("../%s/%s.json", v, path)
		x1, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("[INFO] resource %s not present in schema version: %s", path, v)
			if resRemovedIn == "" && resAddedIn != "" {
				resRemovedIn = v
			}
			continue
		}
		var j1 interface{}
		err = json.Unmarshal(x1, &j1)
		if err != nil {
			log.Fatal(err)
		}
		if resAddedIn == "" {
			resAddedIn = v
		}
		mergedVersions = append(mergedVersions, v)
		files = append(files, j1)
	}
	return files, resRemovedIn, resAddedIn, mergedVersions
}
