package models

const RouterospfRedistributePath = "router.ospf/redistribute/"

type RouterospfRedistribute struct {
	Metric     *int64  `json:"metric,omitempty"`
	MetricType *string `json:"metric-type,omitempty"`
	Name       *string `json:"name,omitempty"`
	Routemap   *string `json:"routemap,omitempty"`
	Status     *string `json:"status,omitempty"`
	Tag        *int64  `json:"tag,omitempty"`
}
