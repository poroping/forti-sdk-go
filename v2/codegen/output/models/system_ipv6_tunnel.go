package models

const SystemIpv6TunnelPath = "system/ipv6-tunnel/"

type SystemIpv6Tunnel struct {
	AutoAsicOffload *string `json:"auto-asic-offload,omitempty"`
	Destination     *string `json:"destination,omitempty"`
	Interface       *string `json:"interface,omitempty"`
	Name            *string `json:"name,omitempty"`
	Source          *string `json:"source,omitempty"`
	UseSdwan        *string `json:"use-sdwan,omitempty"`
}
