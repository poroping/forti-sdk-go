package models

const SystemArpTablePath = "system/arp-table/"

type SystemArpTable struct {
	Fosid     *int64  `json:"fosid,omitempty"`
	Interface *string `json:"interface,omitempty"`
	Ip        *string `json:"ip,omitempty"`
	Mac       *string `json:"mac,omitempty"`
}