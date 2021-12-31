package cmdb

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateLogfortianalyzer3OverrideSetting(payload *models.Logfortianalyzer3OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.Logfortianalyzer3OverrideSettingPath)
	return c.UpdateLogfortianalyzer3OverrideSetting("", payload, params)
}

func (c *Client) ReadLogfortianalyzer3OverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.Logfortianalyzer3OverrideSetting, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.Logfortianalyzer3OverrideSettingPath
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
		v := models.Logfortianalyzer3OverrideSetting{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	err = errors.New("unable to parse API response results")

	return nil, err
}

func (c *Client) UpdateLogfortianalyzer3OverrideSetting(mkey string, payload *models.Logfortianalyzer3OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.Logfortianalyzer3OverrideSettingPath + mkey + "/"
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogfortianalyzer3OverrideSetting(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.Logfortianalyzer3OverrideSetting{}
	_, err := c.UpdateLogfortianalyzer3OverrideSetting("", payload, params)
	return err
}
