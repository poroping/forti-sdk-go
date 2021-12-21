package models

const SystemVxlanPath = "system/vxlan/"

type SystemVxlan struct {
	Dstport      *float64                `json:"dstport,omitempty"`
	Interface    *string                 `json:"interface,omitempty"`
	IpVersion    *string                 `json:"ip-version,omitempty"`
	MulticastTtl *float64                `json:"multicast-ttl,omitempty"`
	Name         *string                 `json:"name,omitempty"`
	RemoteIp     *[]SystemVxlanRemoteIp  `json:"remote-ip,omitempty"`
	RemoteIp6    *[]SystemVxlanRemoteIp6 `json:"remote-ip6,omitempty"`
	Vni          *float64                `json:"vni,omitempty"`
}

type SystemVxlanRemoteIp struct {
	Ip *string `json:"ip,omitempty"`
}

type SystemVxlanRemoteIp6 struct {
	Ip6 *string `json:"ip6,omitempty"`
}
