package models

const RouterbgpNeighborRangePath = "router.bgp/neighbor-range/"

type RouterbgpNeighborRange struct {
	Fosid          *float64 `json:"fosid,omitempty"`
	MaxNeighborNum *float64 `json:"max-neighbor-num,omitempty"`
	NeighborGroup  *string  `json:"neighbor-group,omitempty"`
	Prefix         *string  `json:"prefix,omitempty"`
}
