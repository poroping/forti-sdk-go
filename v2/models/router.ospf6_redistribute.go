package models

const Routerospf6RedistributePath = "router/ospf6/redistribute/"

type Routerospf6Redistribute struct {
	Metric     *int64  `json:"metric,omitempty"`
	MetricType *string `json:"metric-type,omitempty"`
	Name       *string `json:"name,omitempty"`
	Routemap   *string `json:"routemap,omitempty"`
	Status     *string `json:"status,omitempty"`
}
