package models

const RouterospfNetworkPath = "router.ospf/network/"

type RouterospfNetwork struct {
	Area     *string  `json:"area,omitempty"`
	Comments *string  `json:"comments,omitempty"`
	Fosid    *float64 `json:"fosid,omitempty"`
	Prefix   *string  `json:"prefix,omitempty"`
}
