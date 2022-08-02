package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateLogTacacsaccounting2Filter(payload *models.LogTacacsaccounting2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.LogTacacsaccounting2FilterPath)
	return c.UpdateLogTacacsaccounting2Filter("", payload, params)
}

func (c *Client) ReadLogTacacsaccounting2Filter(mkey string, params *models.CmdbRequestParams) (*models.LogTacacsaccounting2Filter, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.LogTacacsaccounting2FilterPath
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
		v := models.LogTacacsaccounting2Filter{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateLogTacacsaccounting2Filter(mkey string, payload *models.LogTacacsaccounting2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.LogTacacsaccounting2FilterPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogTacacsaccounting2Filter(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.LogTacacsaccounting2Filter{}
	_, err := c.UpdateLogTacacsaccounting2Filter("", payload, params)
	return err
}
