package cmdb

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateVpnIpsecConcentrator(payload *models.VpnIpsecConcentrator, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	mkey := ""
	if payload.Id != nil && *params.AllowAppend {
		mkey = strconv.Itoa(int(*payload.Id))
		read, err := c.ReadVpnIpsecConcentrator(mkey, params)
		if err != nil {
			return nil, err
		}
		if read != nil {
			log.Printf("[WARN] Resource at path %q with mkey %q detected upon CREATE with flag set to to overwrite. Changing to UPDATE.", models.VpnIpsecConcentratorPath, mkey)
			return c.UpdateVpnIpsecConcentrator(mkey, payload, params)
		}
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "POST"
	req.Payload = body
	req.Path = models.CmdbBasePath + models.VpnIpsecConcentratorPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) ReadVpnIpsecConcentrator(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecConcentrator, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.VpnIpsecConcentratorPath + mkey + "/"
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
		v := models.VpnIpsecConcentrator{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) UpdateVpnIpsecConcentrator(mkey string, payload *models.VpnIpsecConcentrator, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.VpnIpsecConcentratorPath + mkey + "/"
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteVpnIpsecConcentrator(mkey string, params *models.CmdbRequestParams) error {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "DELETE"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.VpnIpsecConcentratorPath + mkey + "/"
	req.Params = *params

	err := request.Delete(c.config, req)
	return err
}

func (c *Client) ListVpnIpsecConcentrator(mkey string, params *models.CmdbRequestParams) (*[]models.VpnIpsecConcentrator, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.VpnIpsecConcentratorPath
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
		v := []models.VpnIpsecConcentrator{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}
