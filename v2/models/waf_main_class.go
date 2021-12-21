package models

const WafMainClassPath = "waf/main-class/"

type WafMainClass struct {
	Fosid *float64 `json:"fosid,omitempty"`
	Name  *string  `json:"name,omitempty"`
}
