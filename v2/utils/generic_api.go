package utils

import (
	"encoding/json"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

type UtilsGenericApi struct {
	HTTPMethod string `json:"-"`
	Json       string `json:"json"`
	Path       string `json:"-"`
}

func (c *Client) GenericApi(payload *UtilsGenericApi, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = payload.HTTPMethod
	req.Payload = body
	req.Path = payload.Path
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
