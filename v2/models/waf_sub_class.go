package models

const WafSubClassPath = "waf/sub-class/"

type WafSubClass struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
