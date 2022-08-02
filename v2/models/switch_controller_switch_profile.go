package models

const SwitchControllerSwitchProfilePath = "switch-controller/switch-profile/"

type SwitchControllerSwitchProfile struct {
	LoginPasswd         *string `json:"login-passwd,omitempty"`
	LoginPasswdOverride *string `json:"login-passwd-override,omitempty"`
	Name                *string `json:"name,omitempty"`
}
