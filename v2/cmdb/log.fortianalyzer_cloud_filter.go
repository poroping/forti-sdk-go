package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateLogFortianalyzerCloudFilter(payload *models.LogFortianalyzerCloudFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.LogFortianalyzerCloudFilterPath)
	return c.UpdateLogFortianalyzerCloudFilter("", payload, params)
}

func (c *Client) ReadLogFortianalyzerCloudFilter(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzerCloudFilter, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.LogFortianalyzerCloudFilterPath
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
		v := models.LogFortianalyzerCloudFilter{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateLogFortianalyzerCloudFilter(mkey string, payload *models.LogFortianalyzerCloudFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.LogFortianalyzerCloudFilterPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogFortianalyzerCloudFilter(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.LogFortianalyzerCloudFilter{}
	_, err := c.UpdateLogFortianalyzerCloudFilter("", payload, params)
	return err
}
