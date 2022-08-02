package models

const SwitchControllerSwitchLogPath = "switch-controller/switch-log/"

type SwitchControllerSwitchLog struct {
	Severity *string `json:"severity,omitempty"`
	Status   *string `json:"status,omitempty"`
}
