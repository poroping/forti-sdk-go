package cmdb

import (
	"encoding/json"
	"errors"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

func (c *Client) CreateVideofilterProfile(payload *models.VideofilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "POST"
	req.Payload = body
	req.Path = models.CmdbBasePath + models.VideofilterProfilePath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) ReadVideofilterProfile(mkey string, params *models.CmdbRequestParams) (*models.VideofilterProfile, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.VideofilterProfilePath + mkey + "/"
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
		v := models.VideofilterProfile{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	err = errors.New("unable to parse API response results")

	return nil, err
}

func (c *Client) UpdateVideofilterProfile(mkey string, payload *models.VideofilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.VideofilterProfilePath + mkey + "/"
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteVideofilterProfile(mkey string, params *models.CmdbRequestParams) error {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "DELETE"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.VideofilterProfilePath + mkey + "/"
	req.Params = *params

	err := request.Delete(c.config, req)
	return err
}