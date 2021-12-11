package models

const WirelessControllerSnmpPath = "wireless-controller/snmp/"

type WirelessControllerSnmp struct {
	Community            []WirelessControllerSnmpCommunity `json:"community,omitempty"`
	ContactInfo          *string                           `json:"contact-info,omitempty"`
	EngineId             *string                           `json:"engine-id,omitempty"`
	TrapHighCpuThreshold *int64                            `json:"trap-high-cpu-threshold,omitempty"`
	TrapHighMemThreshold *int64                            `json:"trap-high-mem-threshold,omitempty"`
	User                 []WirelessControllerSnmpUser      `json:"user,omitempty"`
}

type WirelessControllerSnmpCommunity struct {
	Hosts          []WirelessControllerSnmpCommunityHosts `json:"hosts,omitempty"`
	Id             *int64                                 `json:"id,omitempty"`
	Name           *string                                `json:"name,omitempty"`
	QueryV1Status  *string                                `json:"query-v1-status,omitempty"`
	QueryV2cStatus *string                                `json:"query-v2c-status,omitempty"`
	Status         *string                                `json:"status,omitempty"`
	TrapV1Status   *string                                `json:"trap-v1-status,omitempty"`
	TrapV2cStatus  *string                                `json:"trap-v2c-status,omitempty"`
}

type WirelessControllerSnmpCommunityHosts struct {
	Id *int64  `json:"id,omitempty"`
	Ip *string `json:"ip,omitempty"`
}

type WirelessControllerSnmpUser struct {
	AuthProto     *string `json:"auth-proto,omitempty"`
	AuthPwd       *string `json:"auth-pwd,omitempty"`
	Name          *string `json:"name,omitempty"`
	NotifyHosts   *string `json:"notify-hosts,omitempty"`
	PrivProto     *string `json:"priv-proto,omitempty"`
	PrivPwd       *string `json:"priv-pwd,omitempty"`
	Queries       *string `json:"queries,omitempty"`
	SecurityLevel *string `json:"security-level,omitempty"`
	Status        *string `json:"status,omitempty"`
	TrapStatus    *string `json:"trap-status,omitempty"`
}
