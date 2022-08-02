package models

const SwitchControllerVirtualPortPoolPath = "switch-controller/virtual-port-pool/"

type SwitchControllerVirtualPortPool struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}
