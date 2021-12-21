package models

const RouteraccessListRulePath = "router.access-list/rule/"

type RouteraccessListRule struct {
	Action     *string  `json:"action,omitempty"`
	ExactMatch *string  `json:"exact-match,omitempty"`
	Flags      *float64 `json:"flags,omitempty"`
	Fosid      *float64 `json:"fosid,omitempty"`
	Prefix     *string  `json:"prefix,omitempty"`
	Wildcard   *string  `json:"wildcard,omitempty"`
}
