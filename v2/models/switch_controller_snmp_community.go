package models

const SwitchControllerSnmpCommunityPath = "switch-controller/snmp-community/"

type SwitchControllerSnmpCommunity struct {
	Events         *string                               `json:"events,omitempty"`
	Hosts          *[]SwitchControllerSnmpCommunityHosts `json:"hosts,omitempty"`
	Id             *int64                                `json:"id,omitempty"`
	Name           *string                               `json:"name,omitempty"`
	QueryV1Port    *int64                                `json:"query-v1-port,omitempty"`
	QueryV1Status  *string                               `json:"query-v1-status,omitempty"`
	QueryV2cPort   *int64                                `json:"query-v2c-port,omitempty"`
	QueryV2cStatus *string                               `json:"query-v2c-status,omitempty"`
	Status         *string                               `json:"status,omitempty"`
	TrapV1Lport    *int64                                `json:"trap-v1-lport,omitempty"`
	TrapV1Rport    *int64                                `json:"trap-v1-rport,omitempty"`
	TrapV1Status   *string                               `json:"trap-v1-status,omitempty"`
	TrapV2cLport   *int64                                `json:"trap-v2c-lport,omitempty"`
	TrapV2cRport   *int64                                `json:"trap-v2c-rport,omitempty"`
	TrapV2cStatus  *string                               `json:"trap-v2c-status,omitempty"`
}

const SwitchControllerSnmpCommunityHostsPath = "switch-controller/snmp-community/hosts/"

type SwitchControllerSnmpCommunityHosts struct {
	Id *int64  `json:"id,omitempty"`
	Ip *string `json:"ip,omitempty"`
}
