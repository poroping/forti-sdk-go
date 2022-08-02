package models

const SwitchControllerSnmpTrapThresholdPath = "switch-controller/snmp-trap-threshold/"

type SwitchControllerSnmpTrapThreshold struct {
	TrapHighCpuThreshold   *int64 `json:"trap-high-cpu-threshold,omitempty"`
	TrapLogFullThreshold   *int64 `json:"trap-log-full-threshold,omitempty"`
	TrapLowMemoryThreshold *int64 `json:"trap-low-memory-threshold,omitempty"`
}
