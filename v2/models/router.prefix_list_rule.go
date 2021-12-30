package models

const RouterprefixListRulePath = "router/prefix-list/rule/"

type RouterprefixListRule struct {
	Action *string `json:"action,omitempty"`
	Flags  *int64  `json:"flags,omitempty"`
	Ge     *int64  `json:"ge,omitempty"`
	Fosid  *int64  `json:"fosid,omitempty"`
	Le     *int64  `json:"le,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
}