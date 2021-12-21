package models

const RouterAspathListPath = "router/aspath-list/"

type RouterAspathList struct {
	Name *string                 `json:"name,omitempty"`
	Rule *[]RouterAspathListRule `json:"rule,omitempty"`
}

type RouterAspathListRule struct {
	Action *string  `json:"action,omitempty"`
	Id     *float64 `json:"id,omitempty"`
	Regexp *string  `json:"regexp,omitempty"`
}
