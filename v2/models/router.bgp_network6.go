package models

const RouterbgpNetwork6Path = "router/bgp/network6/"

type RouterbgpNetwork6 struct {
	Backdoor *string `json:"backdoor,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Prefix6  *string `json:"prefix6,omitempty"`
	RouteMap *string `json:"route-map,omitempty"`
}
