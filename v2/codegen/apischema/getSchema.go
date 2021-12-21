package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/poroping/forti-sdk-go/v2/auth"
	"github.com/poroping/forti-sdk-go/v2/client"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func main() {
	c, err := fclient()
	if err != nil {
		fmt.Print(err)

	}
	paths := getPaths("./paths")

	for _, v := range paths {
		version, filePath, schema, err := getSchema(c, v)
		if err != nil {
			fmt.Print(err)
		}
		jsonSchema, err := json.Marshal(schema)
		if err != nil {
			fmt.Print(err)
			continue
		}
		os.MkdirAll(fmt.Sprintf("./%s", version), os.FileMode(int(0755)))
		os.WriteFile(fmt.Sprintf("./%s/%s.json", version, filePath), jsonSchema, os.FileMode(int(0755)))
	}

}

func fclient() (*client.FortiSDKClient, error) {
	godotenv.Load()
	hostname := os.Getenv("HOSTNAME")
	token := os.Getenv("TOKEN")
	if hostname == "" || token == "" {
		return nil, errors.New("hostname and token must be set")
	}

	auth := &auth.Auth{
		Hostname: hostname,
		Token:    token,
		Vdom:     "root",
		Insecure: true,
	}

	fc, _ := client.NewClientBase(auth)

	return fc, nil
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

func getSchema(c *client.FortiSDKClient, path string) (string, string, interface{}, error) {
	params := &models.CmdbRequestParams{}
	params.Action = "schema"
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Path = models.CmdbBasePath + path
	req.Params = *params
	resp, err := request.Read(c.Config, req)
	if err != nil {
		return "", "", err, nil
	}
	filePath := buildPath(resp)
	return *resp.Version, filePath, resp, nil
}

func buildPath(r *models.CmdbResponse) string {
	path := *r.Path
	path = strings.ReplaceAll(path, "-", "_")
	name := *r.Name
	name = strings.ReplaceAll(name, "-", "_")
	l := []string{path, name}
	if r.ChildPath != nil {
		cp := *r.ChildPath
		cp = strings.ReplaceAll(cp, "-", "_")
		l = append(l, cp)
	}
	return strings.Join(l, "_")
}
