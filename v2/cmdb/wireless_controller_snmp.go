package cmdb

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateWirelessControllerSnmp(payload *models.WirelessControllerSnmp, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.WirelessControllerSnmpPath)
	return c.UpdateWirelessControllerSnmp("", payload, params)
}

func (c *Client) ReadWirelessControllerSnmp(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerSnmp, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.WirelessControllerSnmpPath + mkey + "/"
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
		v := models.WirelessControllerSnmp{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	err = errors.New("unable to parse API response results")

	return nil, err
}

func (c *Client) UpdateWirelessControllerSnmp(mkey string, payload *models.WirelessControllerSnmp, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.WirelessControllerSnmpPath + mkey + "/"
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteWirelessControllerSnmp(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.WirelessControllerSnmp{}
	_, err := c.UpdateWirelessControllerSnmp("", payload, params)
	return err
}
