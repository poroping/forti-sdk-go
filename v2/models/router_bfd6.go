package models

const RouterBfd6Path = "router/bfd6/"

type RouterBfd6 struct {
	Neighbor *[]RouterBfd6Neighbor `json:"neighbor,omitempty"`
}

const RouterBfd6NeighborPath = "router/bfd6/neighbor/"

type RouterBfd6Neighbor struct {
	Interface  *string `json:"interface,omitempty"`
	Ip6Address *string `json:"ip6-address,omitempty"`
}
