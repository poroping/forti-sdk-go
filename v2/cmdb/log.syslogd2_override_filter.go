package cmdb

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateLogsyslogd2OverrideFilter(payload *models.Logsyslogd2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.Logsyslogd2OverrideFilterPath)
	return c.UpdateLogsyslogd2OverrideFilter("", payload, params)
}

func (c *Client) ReadLogsyslogd2OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.Logsyslogd2OverrideFilter, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.Logsyslogd2OverrideFilterPath + mkey + "/"
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
		v := models.Logsyslogd2OverrideFilter{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	err = errors.New("unable to parse API response results")

	return nil, err
}

func (c *Client) UpdateLogsyslogd2OverrideFilter(mkey string, payload *models.Logsyslogd2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.Logsyslogd2OverrideFilterPath + mkey + "/"
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogsyslogd2OverrideFilter(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.Logsyslogd2OverrideFilter{}
	_, err := c.UpdateLogsyslogd2OverrideFilter("", payload, params)
	return err
}
