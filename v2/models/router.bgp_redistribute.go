package models

const RouterbgpRedistributePath = "router/bgp/redistribute/"

type RouterbgpRedistribute struct {
	Name     *string `json:"name,omitempty"`
	RouteMap *string `json:"route-map,omitempty"`
	Status   *string `json:"status,omitempty"`
}
