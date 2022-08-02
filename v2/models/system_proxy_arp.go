package models

const SystemProxyArpPath = "system/proxy-arp/"

type SystemProxyArp struct {
	EndIp     *string `json:"end-ip,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	Interface *string `json:"interface,omitempty"`
	Ip        *string `json:"ip,omitempty"`
}
