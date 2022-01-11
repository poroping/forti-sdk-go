package monitor

import (
	"encoding/json"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

const MonitorRouterBgpPathsPath = "router/bgp/paths/"

type MonitorRouterBgpPathsResults struct {
	Results []MonitorRouterBgpPathsResult `json:"results,omitempty"`
}

type MonitorRouterBgpPathsResult struct {
	NlriPrefix    *string `json:"nlri_prefix,omitempty"`
	NlriPrefixLen *int64  `json:"nlri_prefix_len,omitempty"`
	LearnedFrom   *string `json:"learned_from,omitempty"`
	NextHop       *string `json:"next_hop,omitempty"`
	Origin        *string `json:"origin,omitempty"`
	IsBest        *bool   `json:"is_best,omitempty"`
}

type MonitorRouterBgpPathsFilter struct {
	Vdom *string
}

func (c *Client) MonitorRouterBgpPaths() (*[]MonitorRouterBgpPathsResult, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Path = models.MonitorBasePath + MonitorRouterBgpPathsPath

	res, err := request.Read(c.config, req)
	if err != nil {
		return nil, err
	}

	if tmp, ok := res.Results.([]interface{}); ok {
		jsontmp, err := json.Marshal(tmp[0])
		if err != nil {
			return nil, err
		}
		v := []MonitorRouterBgpPathsResult{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}
