package models

const WanoptSettingsPath = "wanopt/settings/"

type WanoptSettings struct {
	AutoDetectAlgorithm *string `json:"auto-detect-algorithm,omitempty"`
	HostId              *string `json:"host-id,omitempty"`
	TunnelSslAlgorithm  *string `json:"tunnel-ssl-algorithm,omitempty"`
}
