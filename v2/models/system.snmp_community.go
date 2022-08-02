package models

const SystemSnmpCommunityPath = "system.snmp/community/"

type SystemSnmpCommunity struct {
	Events         *string                      `json:"events,omitempty"`
	Hosts          *[]SystemSnmpCommunityHosts  `json:"hosts,omitempty"`
	Hosts6         *[]SystemSnmpCommunityHosts6 `json:"hosts6,omitempty"`
	Id             *int64                       `json:"id,omitempty"`
	Name           *string                      `json:"name,omitempty"`
	QueryV1Port    *int64                       `json:"query-v1-port,omitempty"`
	QueryV1Status  *string                      `json:"query-v1-status,omitempty"`
	QueryV2cPort   *int64                       `json:"query-v2c-port,omitempty"`
	QueryV2cStatus *string                      `json:"query-v2c-status,omitempty"`
	Status         *string                      `json:"status,omitempty"`
	TrapV1Lport    *int64                       `json:"trap-v1-lport,omitempty"`
	TrapV1Rport    *int64                       `json:"trap-v1-rport,omitempty"`
	TrapV1Status   *string                      `json:"trap-v1-status,omitempty"`
	TrapV2cLport   *int64                       `json:"trap-v2c-lport,omitempty"`
	TrapV2cRport   *int64                       `json:"trap-v2c-rport,omitempty"`
	TrapV2cStatus  *string                      `json:"trap-v2c-status,omitempty"`
}

const SystemSnmpCommunityHostsPath = "system.snmp/community/hosts/"

type SystemSnmpCommunityHosts struct {
	HaDirect *string `json:"ha-direct,omitempty"`
	HostType *string `json:"host-type,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Ip       *string `json:"ip,omitempty"`
	SourceIp *string `json:"source-ip,omitempty"`
}

const SystemSnmpCommunityHosts6Path = "system.snmp/community/hosts6/"

type SystemSnmpCommunityHosts6 struct {
	HaDirect   *string `json:"ha-direct,omitempty"`
	HostType   *string `json:"host-type,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	Ipv6       *string `json:"ipv6,omitempty"`
	SourceIpv6 *string `json:"source-ipv6,omitempty"`
}
