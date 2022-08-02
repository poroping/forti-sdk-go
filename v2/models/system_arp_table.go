package models

const SystemArpTablePath = "system/arp-table/"

type SystemArpTable struct {
	Id        *int64  `json:"id,omitempty"`
	Interface *string `json:"interface,omitempty"`
	Ip        *string `json:"ip,omitempty"`
	Mac       *string `json:"mac,omitempty"`
}
