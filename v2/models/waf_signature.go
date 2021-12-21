package models

const WafSignaturePath = "waf/signature/"

type WafSignature struct {
	Desc  *string  `json:"desc,omitempty"`
	Fosid *float64 `json:"fosid,omitempty"`
}
