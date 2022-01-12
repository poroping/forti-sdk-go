package monitor

import (
	"encoding/json"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

const MonitorRouterIPv4Path = "router/ipv4/"

type MonitorRouterIPv4Results struct {
	Results []MonitorRouterIPv4Result `json:"results,omitempty"`
}

type MonitorRouterIPv4Result struct {
	IPVersion     *int64  `json:"ip_version,omitempty"`
	Type          *string `json:"type,omitempty"`
	SubType       *string `json:"subtype,omitempty"`
	IPMask        *string `json:"ip_mask,omitempty"`
	Distance      *int64  `json:"distance,omitempty"`
	Metric        *int64  `json:"metric,omitempty"`
	Priority      *int64  `json:"priority,omitempty"`
	Vrf           *int64  `json:"vrf,omitempty"`
	Gateway       *string `json:"gateway,omitempty"`
	Interface     *string `json:"interface,omitempty"`
	IsTunnelRoute *bool   `json:"is_tunnel_route,omitempty"`
	TunnelParent  *string `json:"tunnel_parent,omitempty"`
	InstallDate   *int64  `json:"install_date,omitempty"`
}

type MonitorRouterIPv4Filter struct {
	Vdom      *string
	IPMask    *string
	Gateway   *string
	Type      *string
	Interface *string
}

func (c *Client) MonitorRouterIPv4() (*[]MonitorRouterIPv4Result, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Path = models.MonitorBasePath + MonitorRouterIPv4Path

	res, err := request.Read(c.config, req)
	if err != nil {
		return nil, err
	}

	if tmp, ok := res.Results.([]interface{}); ok {
		jsontmp, err := json.Marshal(tmp)
		if err != nil {
			return nil, err
		}
		v := []MonitorRouterIPv4Result{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}
