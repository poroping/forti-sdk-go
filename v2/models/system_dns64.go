package models

const SystemDns64Path = "system/dns64/"

type SystemDns64 struct {
	AlwaysSynthesizeAaaaRecord *string `json:"always-synthesize-aaaa-record,omitempty"`
	Dns64Prefix                *string `json:"dns64-prefix,omitempty"`
	Status                     *string `json:"status,omitempty"`
}
