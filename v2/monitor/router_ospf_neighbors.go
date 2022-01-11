package monitor

import (
	"encoding/json"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

const MonitorRouterOspfNeighborsPath = "router/ospf/neighbors/"

type MonitorRouterOspfNeighborsResults struct {
	Results []MonitorRouterOspfNeighborsResult `json:"results,omitempty"`
}

type MonitorRouterOspfNeighborsResult struct {
	NeighborIP *string `json:"neighbor_ip,omitempty"`
	RouterId   *string `json:"router_id,omitempty"`
	State      *string `json:"state,omitempty"`
	Priority   *string `json:"priority,omitempty"`
}

type MonitorRouterOspfNeighborsFilter struct {
	Vdom *string
}

func (c *Client) MonitorRouterOspfNeighbors() (*[]MonitorRouterOspfNeighborsResult, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Path = models.MonitorBasePath + MonitorRouterOspfNeighborsPath

	res, err := request.Read(c.config, req)
	if err != nil {
		return nil, err
	}

	if tmp, ok := res.Results.([]interface{}); ok {
		jsontmp, err := json.Marshal(tmp[0])
		if err != nil {
			return nil, err
		}
		v := []MonitorRouterOspfNeighborsResult{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}
