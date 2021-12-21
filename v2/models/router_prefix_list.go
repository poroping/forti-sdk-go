package models

const RouterPrefixListPath = "router/prefix-list/"

type RouterPrefixList struct {
	Comments *string                 `json:"comments,omitempty"`
	Name     *string                 `json:"name,omitempty"`
	Rule     *[]RouterPrefixListRule `json:"rule,omitempty"`
}

type RouterPrefixListRule struct {
	Action *string  `json:"action,omitempty"`
	Flags  *float64 `json:"flags,omitempty"`
	Ge     *float64 `json:"ge,omitempty"`
	Id     *float64 `json:"id,omitempty"`
	Le     *float64 `json:"le,omitempty"`
	Prefix *string  `json:"prefix,omitempty"`
}
