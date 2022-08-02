package models

const SwitchControllerStormControlPolicyPath = "switch-controller/storm-control-policy/"

type SwitchControllerStormControlPolicy struct {
	Broadcast        *string `json:"broadcast,omitempty"`
	Description      *string `json:"description,omitempty"`
	Name             *string `json:"name,omitempty"`
	Rate             *int64  `json:"rate,omitempty"`
	StormControlMode *string `json:"storm-control-mode,omitempty"`
	UnknownMulticast *string `json:"unknown-multicast,omitempty"`
	UnknownUnicast   *string `json:"unknown-unicast,omitempty"`
}
