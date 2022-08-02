package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateSwitchControllerSnmpSysinfo(payload *models.SwitchControllerSnmpSysinfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.SwitchControllerSnmpSysinfoPath)
	return c.UpdateSwitchControllerSnmpSysinfo("", payload, params)
}

func (c *Client) ReadSwitchControllerSnmpSysinfo(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSnmpSysinfo, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.SwitchControllerSnmpSysinfoPath
	req.Params = *params

	res, err := request.Read(c.config, req)
	if err != nil {
		return nil, err
	}

	// marshal/unmarshal results

	if tmp, ok := res.Results.(interface{}); ok {
		jsontmp, err := json.Marshal(tmp)
		if err != nil {
			return nil, err
		}
		v := models.SwitchControllerSnmpSysinfo{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateSwitchControllerSnmpSysinfo(mkey string, payload *models.SwitchControllerSnmpSysinfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.SwitchControllerSnmpSysinfoPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteSwitchControllerSnmpSysinfo(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.SwitchControllerSnmpSysinfo{}
	_, err := c.UpdateSwitchControllerSnmpSysinfo("", payload, params)
	return err
}
