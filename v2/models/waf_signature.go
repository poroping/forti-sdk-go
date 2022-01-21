package models

const WafSignaturePath = "waf/signature/"

type WafSignature struct {
	Desc *string `json:"desc,omitempty"`
	Id   *int64  `json:"id,omitempty"`
}
