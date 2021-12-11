package forticlient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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

func (c *FortiSDKClient) MonitorRouterBgpNeighbors() (*MonitorRouterBgpNeighborsResults, error) {
	HTTPMethod := "GET"
	path := "/api/v2/monitor/router/bgp/neighbors"

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

	var result MonitorRouterBgpNeighborsResults
	json.Unmarshal([]byte(string(body)), &result)
	// check for errors

	return &result, err
}
