package monitor

import (
	"encoding/json"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

const MonitorRouterBgpNeighborsPath = "router/bgp/neighbors/"

type MonitorRouterBgpNeighborsResults struct {
	Results []MonitorRouterBgpNeighborsResult `json:"results,omitempty"`
}

type MonitorRouterBgpNeighborsResult struct {
	NeighborIP  *string `json:"neighbor_ip,omitempty"`
	LocalIP     *string `json:"local_ip,omitempty"`
	RemoteAs    *int64  `json:"remote_as,omitempty"`
	AdminStatus *bool   `json:"admin_status,omitempty"`
	State       *string `json:"state,omitempty"`
	Type        *string `json:"type,omitempty"`
}

type MonitorRouterBgpNeighborsFilter struct {
	Vdom *string
}

func (c *Client) MonitorRouterBgpNeighbors() (*[]MonitorRouterBgpNeighborsResult, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Path = models.MonitorBasePath + MonitorRouterBgpNeighborsPath

	res, err := request.Read(c.config, req)
	if err != nil {
		return nil, err
	}

	if tmp, ok := res.Results.([]interface{}); ok {
		jsontmp, err := json.Marshal(tmp)
		if err != nil {
			return nil, err
		}
		v := []MonitorRouterBgpNeighborsResult{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}
