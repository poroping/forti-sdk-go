package models

const RouterBfdPath = "router/bfd/"

type RouterBfd struct {
	Neighbor *[]RouterBfdNeighbor `json:"neighbor,omitempty"`
}

const RouterBfdNeighborPath = "router/bfd/neighbor/"

type RouterBfdNeighbor struct {
	Interface *string `json:"interface,omitempty"`
	Ip        *string `json:"ip,omitempty"`
}
