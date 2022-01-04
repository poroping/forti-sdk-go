package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateFtpProxyExplicit(payload *models.FtpProxyExplicit, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.FtpProxyExplicitPath)
	return c.UpdateFtpProxyExplicit("", payload, params)
}

func (c *Client) ReadFtpProxyExplicit(mkey string, params *models.CmdbRequestParams) (*models.FtpProxyExplicit, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.FtpProxyExplicitPath
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
		v := models.FtpProxyExplicit{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateFtpProxyExplicit(mkey string, payload *models.FtpProxyExplicit, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.FtpProxyExplicitPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteFtpProxyExplicit(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.FtpProxyExplicit{}
	_, err := c.UpdateFtpProxyExplicit("", payload, params)
	return err
}
