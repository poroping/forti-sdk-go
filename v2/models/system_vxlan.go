package models

const SystemVxlanPath = "system/vxlan/"

type SystemVxlan struct {
	Dstport      *int64                  `json:"dstport,omitempty"`
	Interface    *string                 `json:"interface,omitempty"`
	IpVersion    *string                 `json:"ip-version,omitempty"`
	MulticastTtl *int64                  `json:"multicast-ttl,omitempty"`
	Name         *string                 `json:"name,omitempty"`
	RemoteIp     *[]SystemVxlanRemoteIp  `json:"remote-ip,omitempty"`
	RemoteIp6    *[]SystemVxlanRemoteIp6 `json:"remote-ip6,omitempty"`
	Vni          *int64                  `json:"vni,omitempty"`
}

const SystemVxlanRemoteIpPath = "system/vxlan/remote-ip/"

type SystemVxlanRemoteIp struct {
	Ip *string `json:"ip,omitempty"`
}

const SystemVxlanRemoteIp6Path = "system/vxlan/remote-ip6/"

type SystemVxlanRemoteIp6 struct {
	Ip6 *string `json:"ip6,omitempty"`
}
