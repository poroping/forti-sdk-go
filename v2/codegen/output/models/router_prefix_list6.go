package models

const RouterPrefixList6Path = "router/prefix-list6/"

type RouterPrefixList6 struct {
	Comments *string                 `json:"comments,omitempty"`
	Name     *string                 `json:"name,omitempty"`
	Rule     []RouterPrefixList6Rule `json:"rule,omitempty"`
}

type RouterPrefixList6Rule struct {
	Action  *string `json:"action,omitempty"`
	Flags   *int64  `json:"flags,omitempty"`
	Ge      *int64  `json:"ge,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	Le      *int64  `json:"le,omitempty"`
	Prefix6 *string `json:"prefix6,omitempty"`
}
