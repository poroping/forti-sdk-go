package models

const RouterospfNeighborPath = "router.ospf/neighbor/"

type RouterospfNeighbor struct {
	Cost         *float64 `json:"cost,omitempty"`
	Fosid        *float64 `json:"fosid,omitempty"`
	Ip           *string  `json:"ip,omitempty"`
	PollInterval *float64 `json:"poll-interval,omitempty"`
	Priority     *float64 `json:"priority,omitempty"`
}
