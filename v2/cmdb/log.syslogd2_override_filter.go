package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateLogSyslogd2OverrideFilter(payload *models.LogSyslogd2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.LogSyslogd2OverrideFilterPath)
	return c.UpdateLogSyslogd2OverrideFilter("", payload, params)
}

func (c *Client) ReadLogSyslogd2OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd2OverrideFilter, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.LogSyslogd2OverrideFilterPath
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
		v := models.LogSyslogd2OverrideFilter{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateLogSyslogd2OverrideFilter(mkey string, payload *models.LogSyslogd2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.LogSyslogd2OverrideFilterPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogSyslogd2OverrideFilter(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.LogSyslogd2OverrideFilter{}
	_, err := c.UpdateLogSyslogd2OverrideFilter("", payload, params)
	return err
}
