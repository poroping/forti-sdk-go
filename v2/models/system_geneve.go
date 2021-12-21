package models

const SystemGenevePath = "system/geneve/"

type SystemGeneve struct {
	Dstport   *float64 `json:"dstport,omitempty"`
	Interface *string  `json:"interface,omitempty"`
	IpVersion *string  `json:"ip-version,omitempty"`
	Name      *string  `json:"name,omitempty"`
	RemoteIp  *string  `json:"remote-ip,omitempty"`
	RemoteIp6 *string  `json:"remote-ip6,omitempty"`
	Type      *string  `json:"type,omitempty"`
	Vni       *float64 `json:"vni,omitempty"`
}
