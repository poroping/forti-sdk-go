package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateLogSyslogd3OverrideFilter(payload *models.LogSyslogd3OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.LogSyslogd3OverrideFilterPath)
	return c.UpdateLogSyslogd3OverrideFilter("", payload, params)
}

func (c *Client) ReadLogSyslogd3OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd3OverrideFilter, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.LogSyslogd3OverrideFilterPath
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
		v := models.LogSyslogd3OverrideFilter{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateLogSyslogd3OverrideFilter(mkey string, payload *models.LogSyslogd3OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.LogSyslogd3OverrideFilterPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogSyslogd3OverrideFilter(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.LogSyslogd3OverrideFilter{}
	_, err := c.UpdateLogSyslogd3OverrideFilter("", payload, params)
	return err
}
