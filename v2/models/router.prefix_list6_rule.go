package models

const RouterprefixList6RulePath = "router/prefix-list6/rule/"

type RouterprefixList6Rule struct {
	Action  *string `json:"action,omitempty"`
	Flags   *int64  `json:"flags,omitempty"`
	Ge      *int64  `json:"ge,omitempty"`
	Fosid   *int64  `json:"fosid,omitempty"`
	Le      *int64  `json:"le,omitempty"`
	Prefix6 *string `json:"prefix6,omitempty"`
}
