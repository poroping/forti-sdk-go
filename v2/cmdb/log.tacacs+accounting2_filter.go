package cmdb

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateLogtacacsaccounting2Filter(payload *models.Logtacacsaccounting2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.Logtacacsaccounting2FilterPath)
	return c.UpdateLogtacacsaccounting2Filter("", payload, params)
}

func (c *Client) ReadLogtacacsaccounting2Filter(mkey string, params *models.CmdbRequestParams) (*models.Logtacacsaccounting2Filter, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.Logtacacsaccounting2FilterPath + mkey + "/"
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
		v := models.Logtacacsaccounting2Filter{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	err = errors.New("unable to parse API response results")

	return nil, err
}

func (c *Client) UpdateLogtacacsaccounting2Filter(mkey string, payload *models.Logtacacsaccounting2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.Logtacacsaccounting2FilterPath + mkey + "/"
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogtacacsaccounting2Filter(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.Logtacacsaccounting2Filter{}
	_, err := c.UpdateLogtacacsaccounting2Filter("", payload, params)
	return err
}
