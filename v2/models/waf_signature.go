package models

const WafSignaturePath = "waf/signature/"

type WafSignature struct {
	Desc  *string `json:"desc,omitempty"`
	Fosid *int64  `json:"fosid,omitempty"`
}
