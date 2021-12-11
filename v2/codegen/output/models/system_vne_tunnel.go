package models

const SystemVneTunnelPath = "system/vne-tunnel/"

type SystemVneTunnel struct {
	AutoAsicOffload *string `json:"auto-asic-offload,omitempty"`
	BmrHostname     *string `json:"bmr-hostname,omitempty"`
	Br              *string `json:"br,omitempty"`
	Interface       *string `json:"interface,omitempty"`
	Ipv4Address     *string `json:"ipv4-address,omitempty"`
	Mode            *string `json:"mode,omitempty"`
	SslCertificate  *string `json:"ssl-certificate,omitempty"`
	Status          *string `json:"status,omitempty"`
	UpdateUrl       *string `json:"update-url,omitempty"`
}
