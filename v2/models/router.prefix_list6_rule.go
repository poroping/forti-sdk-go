package models

const RouterprefixList6RulePath = "router.prefix-list6/rule/"

type RouterprefixList6Rule struct {
	Action  *string  `json:"action,omitempty"`
	Flags   *float64 `json:"flags,omitempty"`
	Ge      *float64 `json:"ge,omitempty"`
	Fosid   *float64 `json:"fosid,omitempty"`
	Le      *float64 `json:"le,omitempty"`
	Prefix6 *string  `json:"prefix6,omitempty"`
}
