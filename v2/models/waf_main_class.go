package models

const WafMainClassPath = "waf/main-class/"

type WafMainClass struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
