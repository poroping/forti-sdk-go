package models

const SwitchControllerSnmpUserPath = "switch-controller/snmp-user/"

type SwitchControllerSnmpUser struct {
	AuthProto     *string `json:"auth-proto,omitempty"`
	AuthPwd       *string `json:"auth-pwd,omitempty"`
	Name          *string `json:"name,omitempty"`
	PrivProto     *string `json:"priv-proto,omitempty"`
	PrivPwd       *string `json:"priv-pwd,omitempty"`
	Queries       *string `json:"queries,omitempty"`
	QueryPort     *int64  `json:"query-port,omitempty"`
	SecurityLevel *string `json:"security-level,omitempty"`
}
