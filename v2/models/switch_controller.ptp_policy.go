package models

const SwitchControllerPtpPolicyPath = "switch-controller.ptp/policy/"

type SwitchControllerPtpPolicy struct {
	Name   *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
}
