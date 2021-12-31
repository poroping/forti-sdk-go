package cmdb

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateLogfortianalyzerOverrideFilter(payload *models.LogfortianalyzerOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.LogfortianalyzerOverrideFilterPath)
	return c.UpdateLogfortianalyzerOverrideFilter("", payload, params)
}

func (c *Client) ReadLogfortianalyzerOverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogfortianalyzerOverrideFilter, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.LogfortianalyzerOverrideFilterPath
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
		v := models.LogfortianalyzerOverrideFilter{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	err = errors.New("unable to parse API response results")

	return nil, err
}

func (c *Client) UpdateLogfortianalyzerOverrideFilter(mkey string, payload *models.LogfortianalyzerOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.LogfortianalyzerOverrideFilterPath + mkey + "/"
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogfortianalyzerOverrideFilter(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.LogfortianalyzerOverrideFilter{}
	_, err := c.UpdateLogfortianalyzerOverrideFilter("", payload, params)
	return err
}
