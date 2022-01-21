package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateSystemAutoupdateTunneling(payload *models.SystemAutoupdateTunneling, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.SystemAutoupdateTunnelingPath)
	return c.UpdateSystemAutoupdateTunneling("", payload, params)
}

func (c *Client) ReadSystemAutoupdateTunneling(mkey string, params *models.CmdbRequestParams) (*models.SystemAutoupdateTunneling, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.SystemAutoupdateTunnelingPath
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
		v := models.SystemAutoupdateTunneling{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateSystemAutoupdateTunneling(mkey string, payload *models.SystemAutoupdateTunneling, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.SystemAutoupdateTunnelingPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteSystemAutoupdateTunneling(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.SystemAutoupdateTunneling{}
	_, err := c.UpdateSystemAutoupdateTunneling("", payload, params)
	return err
}
