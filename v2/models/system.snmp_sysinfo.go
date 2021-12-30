package models

const SystemsnmpSysinfoPath = "system/snmp/sysinfo/"

type SystemsnmpSysinfo struct {
	ContactInfo            *string `json:"contact-info,omitempty"`
	Description            *string `json:"description,omitempty"`
	EngineId               *string `json:"engine-id,omitempty"`
	EngineIdType           *string `json:"engine-id-type,omitempty"`
	Location               *string `json:"location,omitempty"`
	Status                 *string `json:"status,omitempty"`
	TrapHighCpuThreshold   *int64  `json:"trap-high-cpu-threshold,omitempty"`
	TrapLogFullThreshold   *int64  `json:"trap-log-full-threshold,omitempty"`
	TrapLowMemoryThreshold *int64  `json:"trap-low-memory-threshold,omitempty"`
}
