package forticlient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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

func (m MonitorRouterIPv4Results) HasDefault() bool {
	defaultRoute := "0.0.0.0/0"
	for _, route := range m.Results {
		if *route.IPMask == defaultRoute {
			return true
		}
	}
	return false
}

func (c *FortiSDKClient) MonitorRouterIPv4() (*MonitorRouterIPv4Results, error) {
	HTTPMethod := "GET"
	path := "/api/v2/monitor/router/ipv4"

	params := make(map[string][]string)
	req := c.NewRequest(HTTPMethod, path, &params, nil, 0)
	err := req.Send3("root")
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return nil, err
	}
	// log.Printf("FOS-fortios reading response: %s", string(body))

	var result MonitorRouterIPv4Results
	json.Unmarshal([]byte(string(body)), &result)
	// check for errors

	return &result, err
}
