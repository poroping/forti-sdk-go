package forticlient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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

func (c *FortiSDKClient) MonitorRouterBgpPaths() (*MonitorRouterBgpPathsResults, error) {
	HTTPMethod := "GET"
	path := "/api/v2/monitor/router/bgp/paths"

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

	var result MonitorRouterBgpPathsResults
	json.Unmarshal([]byte(string(body)), &result)
	// check for errors

	return &result, err
}
