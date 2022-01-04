package models

const SystemSnmpUserPath = "system.snmp/user/"

type SystemSnmpUser struct {
	AuthProto     *string `json:"auth-proto,omitempty"`
	AuthPwd       *string `json:"auth-pwd,omitempty"`
	Events        *string `json:"events,omitempty"`
	HaDirect      *string `json:"ha-direct,omitempty"`
	Name          *string `json:"name,omitempty"`
	NotifyHosts   *string `json:"notify-hosts,omitempty"`
	NotifyHosts6  *string `json:"notify-hosts6,omitempty"`
	PrivProto     *string `json:"priv-proto,omitempty"`
	PrivPwd       *string `json:"priv-pwd,omitempty"`
	Queries       *string `json:"queries,omitempty"`
	QueryPort     *int64  `json:"query-port,omitempty"`
	SecurityLevel *string `json:"security-level,omitempty"`
	SourceIp      *string `json:"source-ip,omitempty"`
	SourceIpv6    *string `json:"source-ipv6,omitempty"`
	Status        *string `json:"status,omitempty"`
	TrapLport     *int64  `json:"trap-lport,omitempty"`
	TrapRport     *int64  `json:"trap-rport,omitempty"`
	TrapStatus    *string `json:"trap-status,omitempty"`
}
