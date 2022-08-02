package models

const SwitchControllerStormControlPath = "switch-controller/storm-control/"

type SwitchControllerStormControl struct {
	Broadcast        *string `json:"broadcast,omitempty"`
	Rate             *int64  `json:"rate,omitempty"`
	UnknownMulticast *string `json:"unknown-multicast,omitempty"`
	UnknownUnicast   *string `json:"unknown-unicast,omitempty"`
}
