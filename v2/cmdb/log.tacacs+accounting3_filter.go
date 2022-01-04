package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateLogTacacsaccounting3Filter(payload *models.LogTacacsaccounting3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.LogTacacsaccounting3FilterPath)
	return c.UpdateLogTacacsaccounting3Filter("", payload, params)
}

func (c *Client) ReadLogTacacsaccounting3Filter(mkey string, params *models.CmdbRequestParams) (*models.LogTacacsaccounting3Filter, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.LogTacacsaccounting3FilterPath
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
		v := models.LogTacacsaccounting3Filter{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateLogTacacsaccounting3Filter(mkey string, payload *models.LogTacacsaccounting3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.LogTacacsaccounting3FilterPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogTacacsaccounting3Filter(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.LogTacacsaccounting3Filter{}
	_, err := c.UpdateLogTacacsaccounting3Filter("", payload, params)
	return err
}
