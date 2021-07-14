package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/poroping/forti-sdk-go/fortios/request"
)

// JSONJSONGenericAPI contains the parameters for Create API function
type JSONJSONGenericAPI struct {
	Path          string `json:"path"`
	Method        string `json:"method"`
	Specialparams string `json:"specialparams"`
	Json          string `json:"json"`
	Response      string `json:"response"`
}

// CreateJSONGenericAPI API operation for FortiOS sends request to FortiGate/FortiOS APIs.
// Returns the response from FortiGate or FortiOS .
func (c *FortiSDKClient) CreateJSONGenericAPI(params *JSONJSONGenericAPI, vdomparam string, batch int) (res string, err error) {
	HTTPMethod := params.Method
	path := params.Path
	urlparams := make(map[string][]string)
	if params.Specialparams != "" {
		urlparams, err = url.ParseQuery(params.Specialparams)
		if err != nil {
			err = fmt.Errorf("error parsing url parameters %q", params.Specialparams)
			return
		}
	}

	var req *request.Request

	if params.Json != "" {
		locJSON := []byte(params.Json)

		var js json.RawMessage
		if json.Unmarshal([]byte(params.Json), &js) != nil {
			err = fmt.Errorf("JSON format Error")
			return
		}

		log.Printf("FOS-fortios resquest1: %s", locJSON)
		bytes := bytes.NewBuffer(locJSON)

		req = c.NewRequest(HTTPMethod, path, &urlparams, bytes, batch)
	} else {
		req = c.NewRequest(HTTPMethod, path, &urlparams, nil, batch)
	}

	err = req.Send3(vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}
	log.Printf("FOS-fortios response1: %s", string(body))

	res = string(body)

	return
}

// UpdateJSONGenericAPI API operation for FortiOS
func (c *FortiSDKClient) UpdateJSONGenericAPI(params *JSONJSONGenericAPI, mkey string) (res string, err error) {
	return
}

// DeleteJSONGenericAPI API operation for FortiOS
func (c *FortiSDKClient) DeleteJSONGenericAPI(mkey string) (err error) {
	return
}

// ReadJSONGenericAPI API operation for FortiOS
func (c *FortiSDKClient) ReadJSONGenericAPI(mkey string) (output *JSONJSONGenericAPI, err error) {
	return
}

// GenericGroupRead API operation for FortiOS, Read Generic Group
func (c *FortiSDKClient) GenericGroupRead(path, specialparams, vdomparam string, batch int) (mapTmp []interface{}, err error) {
	urlparams := make(map[string][]string)
	if specialparams != "" {
		urlparams, err = url.ParseQuery(specialparams)
		if err != nil {
			err = fmt.Errorf("error parsing url parameters %q", specialparams)
			return
		}
	}
	req := c.NewRequest("GET", path, &urlparams, nil, batch)
	err = req.Send3(vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if fortiAPIHttpStatus404Checking(result) {
		mapTmp = nil
		return
	}

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		mapTmp = result["results"].([]interface{})
	}

	return
}
