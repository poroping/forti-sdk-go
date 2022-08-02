package models

const SwitchControllerCustomCommandPath = "switch-controller/custom-command/"

type SwitchControllerCustomCommand struct {
	Command     *string `json:"command,omitempty"`
	CommandName *string `json:"command-name,omitempty"`
	Description *string `json:"description,omitempty"`
}
