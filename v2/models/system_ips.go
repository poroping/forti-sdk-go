package models

const SystemIpsPath = "system/ips/"

type SystemIps struct {
	OverrideSignatureHoldById *string `json:"override-signature-hold-by-id,omitempty"`
	SignatureHoldTime         *string `json:"signature-hold-time,omitempty"`
}
