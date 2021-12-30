package models

const RouterAccessListPath = "router/access-list/"

type RouterAccessList struct {
	Comments *string                 `json:"comments,omitempty"`
	Name     *string                 `json:"name,omitempty"`
	Rule     *[]RouterAccessListRule `json:"rule,omitempty"`
}

type RouterAccessListRule struct {
	Action     *string `json:"action,omitempty"`
	ExactMatch *string `json:"exact-match,omitempty"`
	Flags      *int64  `json:"flags,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	Prefix     *string `json:"prefix,omitempty"`
	Wildcard   *string `json:"wildcard,omitempty"`
}
