package models

const SwitchControllerSnmpCommunityPath = "switch-controller/snmp-community/"

type SwitchControllerSnmpCommunity struct {
	Events         *string                               `json:"events,omitempty"`
	Hosts          *[]SwitchControllerSnmpCommunityHosts `json:"hosts,omitempty"`
	Fosid          *float64                              `json:"fosid,omitempty"`
	Name           *string                               `json:"name,omitempty"`
	QueryV1Port    *float64                              `json:"query-v1-port,omitempty"`
	QueryV1Status  *string                               `json:"query-v1-status,omitempty"`
	QueryV2cPort   *float64                              `json:"query-v2c-port,omitempty"`
	QueryV2cStatus *string                               `json:"query-v2c-status,omitempty"`
	Status         *string                               `json:"status,omitempty"`
	TrapV1Lport    *float64                              `json:"trap-v1-lport,omitempty"`
	TrapV1Rport    *float64                              `json:"trap-v1-rport,omitempty"`
	TrapV1Status   *string                               `json:"trap-v1-status,omitempty"`
	TrapV2cLport   *float64                              `json:"trap-v2c-lport,omitempty"`
	TrapV2cRport   *float64                              `json:"trap-v2c-rport,omitempty"`
	TrapV2cStatus  *string                               `json:"trap-v2c-status,omitempty"`
}

type SwitchControllerSnmpCommunityHosts struct {
	Id *float64 `json:"id,omitempty"`
	Ip *string  `json:"ip,omitempty"`
}
