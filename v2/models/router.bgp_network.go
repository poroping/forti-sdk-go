package models

const RouterbgpNetworkPath = "router.bgp/network/"

type RouterbgpNetwork struct {
	Backdoor *string `json:"backdoor,omitempty"`
	Fosid    *int64  `json:"fosid,omitempty"`
	Prefix   *string `json:"prefix,omitempty"`
	RouteMap *string `json:"route-map,omitempty"`
}
