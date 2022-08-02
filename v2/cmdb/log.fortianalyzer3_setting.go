package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateLogFortianalyzer3Setting(payload *models.LogFortianalyzer3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.LogFortianalyzer3SettingPath)
	return c.UpdateLogFortianalyzer3Setting("", payload, params)
}

func (c *Client) ReadLogFortianalyzer3Setting(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzer3Setting, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.LogFortianalyzer3SettingPath
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
		v := models.LogFortianalyzer3Setting{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateLogFortianalyzer3Setting(mkey string, payload *models.LogFortianalyzer3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.LogFortianalyzer3SettingPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogFortianalyzer3Setting(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.LogFortianalyzer3Setting{}
	_, err := c.UpdateLogFortianalyzer3Setting("", payload, params)
	return err
}
