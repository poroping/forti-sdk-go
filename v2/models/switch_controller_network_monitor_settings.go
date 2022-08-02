package models

const SwitchControllerNetworkMonitorSettingsPath = "switch-controller/network-monitor-settings/"

type SwitchControllerNetworkMonitorSettings struct {
	NetworkMonitoring *string `json:"network-monitoring,omitempty"`
}
