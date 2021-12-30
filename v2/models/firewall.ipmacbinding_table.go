package models

const FirewallipmacbindingTablePath = "firewall/ipmacbinding/table/"

type FirewallipmacbindingTable struct {
	Ip     *string `json:"ip,omitempty"`
	Mac    *string `json:"mac,omitempty"`
	Name   *string `json:"name,omitempty"`
	SeqNum *int64  `json:"seq-num,omitempty"`
	Status *string `json:"status,omitempty"`
}
