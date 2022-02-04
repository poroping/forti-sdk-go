package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateFirewallWildcardFqdnGroup(payload *models.FirewallWildcardFqdnGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	mkey := ""
	if payload.Name != nil && *params.AllowAppend {
		mkey = *payload.Name
		read, err := c.ReadFirewallWildcardFqdnGroup(mkey, params)
		if err != nil {
			return nil, err
		}
		if read != nil {
			log.Printf("[WARN] Resource at path %q with mkey %q detected upon CREATE with flag set to to overwrite. Changing to UPDATE.", models.FirewallWildcardFqdnGroupPath, mkey)
			return c.UpdateFirewallWildcardFqdnGroup(mkey, payload, params)
		}
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "POST"
	req.Payload = body
	req.Path = models.CmdbBasePath + models.FirewallWildcardFqdnGroupPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) ReadFirewallWildcardFqdnGroup(mkey string, params *models.CmdbRequestParams) (*models.FirewallWildcardFqdnGroup, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.FirewallWildcardFqdnGroupPath + url.PathEscape(mkey) + "/"
	req.Params = *params

	res, err := request.Read(c.config, req)
	if err != nil {
		return nil, err
	}

	// marshal/unmarshal results

	if tmp, ok := res.Results.([]interface{}); ok {
		jsontmp, err := json.Marshal(tmp[0])
		if err != nil {
			return nil, err
		}
		v := models.FirewallWildcardFqdnGroup{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateFirewallWildcardFqdnGroup(mkey string, payload *models.FirewallWildcardFqdnGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.FirewallWildcardFqdnGroupPath + url.PathEscape(mkey) + "/"
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteFirewallWildcardFqdnGroup(mkey string, params *models.CmdbRequestParams) error {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "DELETE"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.FirewallWildcardFqdnGroupPath + url.PathEscape(mkey) + "/"
	req.Params = *params

	err := request.Delete(c.config, req)
	return err
}

func (c *Client) ListFirewallWildcardFqdnGroup(params *models.CmdbRequestParams) (*[]models.FirewallWildcardFqdnGroup, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.FirewallWildcardFqdnGroupPath
	req.Params = *params

	res, err := request.Read(c.config, req)
	if err != nil {
		return nil, err
	}

	// marshal/unmarshal results

	if tmp, ok := res.Results.([]interface{}); ok {
		jsontmp, err := json.Marshal(tmp)
		if err != nil {
			return nil, err
		}
		v := []models.FirewallWildcardFqdnGroup{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}
