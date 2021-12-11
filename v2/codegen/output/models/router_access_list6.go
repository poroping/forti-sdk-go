package models

const RouterAccessList6Path = "router/access-list6/"

type RouterAccessList6 struct {
	Comments *string                 `json:"comments,omitempty"`
	Name     *string                 `json:"name,omitempty"`
	Rule     []RouterAccessList6Rule `json:"rule,omitempty"`
}

type RouterAccessList6Rule struct {
	Action     *string `json:"action,omitempty"`
	ExactMatch *string `json:"exact-match,omitempty"`
	Flags      *int64  `json:"flags,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	Prefix6    *string `json:"prefix6,omitempty"`
}
