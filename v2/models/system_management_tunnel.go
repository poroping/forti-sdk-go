package models

const SystemManagementTunnelPath = "system/management-tunnel/"

type SystemManagementTunnel struct {
	AllowCollectStatistics *string `json:"allow-collect-statistics,omitempty"`
	AllowConfigRestore     *string `json:"allow-config-restore,omitempty"`
	AllowPushConfiguration *string `json:"allow-push-configuration,omitempty"`
	AllowPushFirmware      *string `json:"allow-push-firmware,omitempty"`
	AuthorizedManagerOnly  *string `json:"authorized-manager-only,omitempty"`
	SerialNumber           *string `json:"serial-number,omitempty"`
	Status                 *string `json:"status,omitempty"`
}
