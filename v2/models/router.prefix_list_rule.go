package models

const RouterprefixListRulePath = "router.prefix-list/rule/"

type RouterprefixListRule struct {
	Action *string  `json:"action,omitempty"`
	Flags  *float64 `json:"flags,omitempty"`
	Ge     *float64 `json:"ge,omitempty"`
	Fosid  *float64 `json:"fosid,omitempty"`
	Le     *float64 `json:"le,omitempty"`
	Prefix *string  `json:"prefix,omitempty"`
}
