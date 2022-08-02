package models

const SwitchControllerSnmpSysinfoPath = "switch-controller/snmp-sysinfo/"

type SwitchControllerSnmpSysinfo struct {
	ContactInfo *string `json:"contact-info,omitempty"`
	Description *string `json:"description,omitempty"`
	EngineId    *string `json:"engine-id,omitempty"`
	Location    *string `json:"location,omitempty"`
	Status      *string `json:"status,omitempty"`
}
