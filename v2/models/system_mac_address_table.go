package models

const SystemMacAddressTablePath = "system/mac-address-table/"

type SystemMacAddressTable struct {
	Interface       *string `json:"interface,omitempty"`
	Mac             *string `json:"mac,omitempty"`
	ReplySubstitute *string `json:"reply-substitute,omitempty"`
}
