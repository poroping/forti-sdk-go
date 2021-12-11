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
	paths := getPaths("./paths.json")

	for _, v := range paths {
		version, filePath, schema, err := getSchema(c, v)
		if err != nil {
			fmt.Print(err)
		}
		jsonSchema, err := json.Marshal(schema)
		if err != nil {
			fmt.Print(err)
		}
		os.MkdirAll(fmt.Sprintf("./%s", version), os.FileMode(int(0755)))
		os.WriteFile(fmt.Sprintf("./%s/%s.json", version, filePath), jsonSchema, os.FileMode(int(0755)))
	}

}

func fclient() (*client.FortiSDKClient, error) {
	godotenv.Load()
	hostname := os.Getenv("HOSTNAME")
	token := os.Getenv("TOKEN")
	print(token)
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

func getPaths(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	paths := &[]string{}
	err = json.Unmarshal(data, paths)
	if err != nil {
		fmt.Print(err)
	}
	return *paths
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
