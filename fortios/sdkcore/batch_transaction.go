package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func (c *FortiSDKClient) StartBatch(timeout int, vdomparam string) (err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb"

	params := make(map[string][]string)
	params["action"] = []string{"transaction-start"}

	data := map[string]int{
		"timeout": timeout,
	}

	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest(HTTPMethod, path, &params, bytes, 0)
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
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

func (c *FortiSDKClient) CommitBatch(vdomparam string, batch int) (err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb"

	params := make(map[string][]string)
	params["action"] = []string{"transaction-commit"}

	data := map[string]int{
		"transaction-id": batch,
	}

	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest(HTTPMethod, path, &params, bytes, 0)
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
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

func (c *FortiSDKClient) AbortBatch(vdomparam string, batch int) (err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb"

	params := make(map[string][]string)
	params["action"] = []string{"transaction-abort"}

	data := map[string]int{
		"transaction-id": batch,
	}

	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest(HTTPMethod, path, &params, bytes, 0)
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
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

func (c *FortiSDKClient) DummyBatch() (res string, err error) {
	return
}
