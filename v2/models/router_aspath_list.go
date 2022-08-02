package models

const RouterAspathListPath = "router/aspath-list/"

type RouterAspathList struct {
	Name *string                 `json:"name,omitempty"`
	Rule *[]RouterAspathListRule `json:"rule,omitempty"`
}

const RouterAspathListRulePath = "router/aspath-list/rule/"

type RouterAspathListRule struct {
	Action *string `json:"action,omitempty"`
	Id     *int64  `json:"id,omitempty"`
	Regexp *string `json:"regexp,omitempty"`
}
