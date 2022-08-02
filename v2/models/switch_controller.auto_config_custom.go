package models

const SwitchControllerAutoConfigCustomPath = "switch-controller.auto-config/custom/"

type SwitchControllerAutoConfigCustom struct {
	Name          *string                                          `json:"name,omitempty"`
	SwitchBinding *[]SwitchControllerAutoConfigCustomSwitchBinding `json:"switch-binding,omitempty"`
}

const SwitchControllerAutoConfigCustomSwitchBindingPath = "switch-controller.auto-config/custom/switch-binding/"

type SwitchControllerAutoConfigCustomSwitchBinding struct {
	Policy   *string `json:"policy,omitempty"`
	SwitchId *string `json:"switch-id,omitempty"`
}
