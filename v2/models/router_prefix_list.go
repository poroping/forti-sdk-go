package models

const RouterPrefixListPath = "router/prefix-list/"

type RouterPrefixList struct {
	Comments *string                `json:"comments,omitempty"`
	Name     *string                `json:"name,omitempty"`
	Rule     []RouterPrefixListRule `json:"rule,omitempty"`
}

type RouterPrefixListRule struct {
	Action *string `json:"action,omitempty"`
	Ge     *int64  `json:"ge,omitempty"`
	Id     *int64  `json:"id,omitempty"`
	Le     *int64  `json:"le,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
}
