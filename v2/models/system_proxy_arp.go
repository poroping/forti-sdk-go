package models

const SystemProxyArpPath = "system/proxy-arp/"

type SystemProxyArp struct {
	EndIp     *string `json:"end-ip,omitempty"`
	Fosid     *int64  `json:"fosid,omitempty"`
	Interface *string `json:"interface,omitempty"`
	Ip        *string `json:"ip,omitempty"`
}
