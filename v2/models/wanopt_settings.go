package models

const WanoptSettingsPath = "wanopt/settings/"

type WanoptSettings struct {
	AutoDetectAlgorithm *string `json:"auto-detect-algorithm,omitempty"`
	HostId              *string `json:"host-id,omitempty"`
	TunnelOptimization  *string `json:"tunnel-optimization,omitempty"`
	TunnelSslAlgorithm  *string `json:"tunnel-ssl-algorithm,omitempty"`
}
