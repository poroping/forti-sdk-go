package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateFirewallIpropeAppctrlStatus(payload *models.FirewallIpropeAppctrlStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.FirewallIpropeAppctrlStatusPath)
	return c.UpdateFirewallIpropeAppctrlStatus("", payload, params)
}

func (c *Client) ReadFirewallIpropeAppctrlStatus(mkey string, params *models.CmdbRequestParams) (*models.FirewallIpropeAppctrlStatus, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.FirewallIpropeAppctrlStatusPath
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
		v := models.FirewallIpropeAppctrlStatus{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateFirewallIpropeAppctrlStatus(mkey string, payload *models.FirewallIpropeAppctrlStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.FirewallIpropeAppctrlStatusPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteFirewallIpropeAppctrlStatus(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.FirewallIpropeAppctrlStatus{}
	_, err := c.UpdateFirewallIpropeAppctrlStatus("", payload, params)
	return err
}
