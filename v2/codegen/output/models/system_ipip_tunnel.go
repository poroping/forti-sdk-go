package models

const SystemIpipTunnelPath = "system/ipip-tunnel/"

type SystemIpipTunnel struct {
	AutoAsicOffload *string `json:"auto-asic-offload,omitempty"`
	Interface       *string `json:"interface,omitempty"`
	LocalGw         *string `json:"local-gw,omitempty"`
	Name            *string `json:"name,omitempty"`
	RemoteGw        *string `json:"remote-gw,omitempty"`
	UseSdwan        *string `json:"use-sdwan,omitempty"`
}
