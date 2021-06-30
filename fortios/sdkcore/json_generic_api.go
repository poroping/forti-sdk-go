package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

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
	path += "?"
	path += params.Specialparams

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

		req = c.NewRequest(HTTPMethod, path, nil, bytes, batch)
	} else {
		req = c.NewRequest(HTTPMethod, path, nil, nil, batch)
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
	path += specialparams
	req := c.NewRequest("GET", path, nil, nil, batch)
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

	if fortiAPIHttpStatus404Checking(result) == true {
		mapTmp = nil
		return
	}

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		mapTmp = result["results"].([]interface{})
	}

	return
}
