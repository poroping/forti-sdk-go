package models

const RouterospfNetworkPath = "router/ospf/network/"

type RouterospfNetwork struct {
	Area     *string `json:"area,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Prefix   *string `json:"prefix,omitempty"`
}
