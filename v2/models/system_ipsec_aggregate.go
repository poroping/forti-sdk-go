package models

const SystemIpsecAggregatePath = "system/ipsec-aggregate/"

type SystemIpsecAggregate struct {
	Algorithm *string                       `json:"algorithm,omitempty"`
	Member    *[]SystemIpsecAggregateMember `json:"member,omitempty"`
	Name      *string                       `json:"name,omitempty"`
}

type SystemIpsecAggregateMember struct {
	TunnelName *string `json:"tunnel-name,omitempty"`
}
