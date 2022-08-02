package models

const MonitoringNpuHpePath = "monitoring/npu-hpe/"

type MonitoringNpuHpe struct {
	Interval    *int64  `json:"interval,omitempty"`
	Multipliers *string `json:"multipliers,omitempty"`
	Status      *string `json:"status,omitempty"`
}
