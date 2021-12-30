package models

const SystemHaMonitorPath = "system/ha-monitor/"

type SystemHaMonitor struct {
	MonitorVlan         *string `json:"monitor-vlan,omitempty"`
	VlanHbInterval      *int64  `json:"vlan-hb-interval,omitempty"`
	VlanHbLostThreshold *int64  `json:"vlan-hb-lost-threshold,omitempty"`
}
