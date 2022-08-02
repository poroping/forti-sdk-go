package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateFirewallIpv6EhFilter(payload *models.FirewallIpv6EhFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.FirewallIpv6EhFilterPath)
	return c.UpdateFirewallIpv6EhFilter("", payload, params)
}

func (c *Client) ReadFirewallIpv6EhFilter(mkey string, params *models.CmdbRequestParams) (*models.FirewallIpv6EhFilter, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.FirewallIpv6EhFilterPath
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
		v := models.FirewallIpv6EhFilter{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateFirewallIpv6EhFilter(mkey string, payload *models.FirewallIpv6EhFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.FirewallIpv6EhFilterPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteFirewallIpv6EhFilter(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.FirewallIpv6EhFilter{}
	_, err := c.UpdateFirewallIpv6EhFilter("", payload, params)
	return err
}
