package models

const RouterospfNeighborPath = "router/ospf/neighbor/"

type RouterospfNeighbor struct {
	Cost         *int64  `json:"cost,omitempty"`
	Id           *int64  `json:"id,omitempty"`
	Ip           *string `json:"ip,omitempty"`
	PollInterval *int64  `json:"poll-interval,omitempty"`
	Priority     *int64  `json:"priority,omitempty"`
}
