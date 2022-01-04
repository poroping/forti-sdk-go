package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateLogSyslogd3Filter(payload *models.LogSyslogd3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.LogSyslogd3FilterPath)
	return c.UpdateLogSyslogd3Filter("", payload, params)
}

func (c *Client) ReadLogSyslogd3Filter(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd3Filter, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.LogSyslogd3FilterPath
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
		v := models.LogSyslogd3Filter{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateLogSyslogd3Filter(mkey string, payload *models.LogSyslogd3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.LogSyslogd3FilterPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogSyslogd3Filter(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.LogSyslogd3Filter{}
	_, err := c.UpdateLogSyslogd3Filter("", payload, params)
	return err
}
