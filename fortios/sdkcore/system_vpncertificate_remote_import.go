package forticlient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type JSONSystemVpnCertificateRemoteImport struct {
	Certificate string `json:"file_content"`
	Scope       string `json:"scope"`
}

// Uploads certificate to Fortigate
func (c *FortiSDKClient) CreateSystemVpnCertificateRemoteImport(data *JSONSystemVpnCertificateRemoteImport, vdomparam string, batch int) (res string, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/monitor/vpn-certificate/remote/import"
	data.Certificate = base64.StdEncoding.EncodeToString([]byte(data.Certificate))
	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return
	}

	var before []string
	before, err = c.GetRemoteCertName(vdomparam, batch)
	if err != nil {
		err = fmt.Errorf("cannot get list of existing cert names: %s", err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes, batch)
	err = req.Send3(vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	if result["status"] == "error" {
		err = fortiAPIErrorFormat(result, string(body))
		if err != nil {
			err = fmt.Errorf("error importing certificate: %s", err)
			return
		}
	}

	var after []string
	after, err = c.GetRemoteCertName(vdomparam, batch)
	if err != nil {
		err = fmt.Errorf("cannot get updated list of cert names: %s", err)
		return
	}

	diff := difference(after, before)

	if len(diff) == 0 {
		err = fmt.Errorf("error finding new cert")
		return
	}

	res = diff[0]

	return
}

func (c *FortiSDKClient) GetRemoteCertName(vdomparam string, batch int) (name []string, err error) {
	HTTPMethod := "GET"
	path := "api/v2/cmdb/vpn.certificate/remote"
	urlparams := make(map[string][]string)
	urlparams["filter"] = []string{"name=@REMOTE_Cert_"}
	urlparams["format"] = []string{"name"}

	req := c.NewRequest(HTTPMethod, path, nil, nil, batch)
	err = req.Send3(vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	if len(result["results"].([]interface{})) == 0 {
		name = make([]string, 0)
		return
	}

	for _, v := range result["results"].([]interface{}) {
		s := v.(map[string]interface{})["name"].(string)
		name = append(name, s)
	}

	return
}
func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
