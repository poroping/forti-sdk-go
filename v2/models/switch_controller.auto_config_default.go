package models

const SwitchControllerautoConfigDefaultPath = "switch-controller.auto-config/default/"

type SwitchControllerautoConfigDefault struct {
	FgtPolicy *string `json:"fgt-policy,omitempty"`
	IclPolicy *string `json:"icl-policy,omitempty"`
	IslPolicy *string `json:"isl-policy,omitempty"`
}
