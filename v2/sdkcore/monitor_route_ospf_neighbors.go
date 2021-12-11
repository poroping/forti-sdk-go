package forticlient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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

func (c *FortiSDKClient) MonitorRouterOspfNeighbors() (*MonitorRouterOspfNeighborsResults, error) {
	HTTPMethod := "GET"
	path := "/api/v2/monitor/router/ospf/neighbors"

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

	var result MonitorRouterOspfNeighborsResults
	json.Unmarshal([]byte(string(body)), &result)
	// check for errors

	return &result, err
}
