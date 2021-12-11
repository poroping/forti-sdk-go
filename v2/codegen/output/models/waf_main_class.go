package models

const WafMainClassPath = "waf/main-class/"

type WafMainClass struct {
	Fosid *int64  `json:"fosid,omitempty"`
	Name  *string `json:"name,omitempty"`
}
