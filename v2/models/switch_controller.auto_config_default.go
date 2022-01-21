package models

const SwitchControllerAutoConfigDefaultPath = "switch-controller.auto-config/default/"

type SwitchControllerAutoConfigDefault struct {
	FgtPolicy *string `json:"fgt-policy,omitempty"`
	IclPolicy *string `json:"icl-policy,omitempty"`
	IslPolicy *string `json:"isl-policy,omitempty"`
}
