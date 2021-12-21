package models

const RouteraccessList6RulePath = "router.access-list6/rule/"

type RouteraccessList6Rule struct {
	Action     *string `json:"action,omitempty"`
	ExactMatch *string `json:"exact-match,omitempty"`
	Flags      *int64  `json:"flags,omitempty"`
	Fosid      *int64  `json:"fosid,omitempty"`
	Prefix6    *string `json:"prefix6,omitempty"`
}
