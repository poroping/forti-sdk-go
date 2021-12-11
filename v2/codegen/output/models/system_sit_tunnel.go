package models

const SystemSitTunnelPath = "system/sit-tunnel/"

type SystemSitTunnel struct {
	AutoAsicOffload *string `json:"auto-asic-offload,omitempty"`
	Destination     *string `json:"destination,omitempty"`
	Interface       *string `json:"interface,omitempty"`
	Ip6             *string `json:"ip6,omitempty"`
	Name            *string `json:"name,omitempty"`
	Source          *string `json:"source,omitempty"`
	UseSdwan        *string `json:"use-sdwan,omitempty"`
}
