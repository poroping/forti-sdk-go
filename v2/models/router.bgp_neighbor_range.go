package models

const RouterbgpNeighborRangePath = "router.bgp/neighbor-range/"

type RouterbgpNeighborRange struct {
	Fosid          *int64  `json:"fosid,omitempty"`
	MaxNeighborNum *int64  `json:"max-neighbor-num,omitempty"`
	NeighborGroup  *string `json:"neighbor-group,omitempty"`
	Prefix         *string `json:"prefix,omitempty"`
}
