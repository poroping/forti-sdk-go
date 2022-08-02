package models

const SwitchControllerRemoteLogPath = "switch-controller/remote-log/"

type SwitchControllerRemoteLog struct {
	Csv      *string `json:"csv,omitempty"`
	Facility *string `json:"facility,omitempty"`
	Name     *string `json:"name,omitempty"`
	Port     *int64  `json:"port,omitempty"`
	Server   *string `json:"server,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status   *string `json:"status,omitempty"`
}
