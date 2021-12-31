package models

const RouterbgpNetworkPath = "router/bgp/network/"

type RouterbgpNetwork struct {
	Backdoor *string `json:"backdoor,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Prefix   *string `json:"prefix,omitempty"`
	RouteMap *string `json:"route-map,omitempty"`
}
