package models

const SwitchControllerSflowPath = "switch-controller/sflow/"

type SwitchControllerSflow struct {
	CollectorIp   *string `json:"collector-ip,omitempty"`
	CollectorPort *int64  `json:"collector-port,omitempty"`
}
