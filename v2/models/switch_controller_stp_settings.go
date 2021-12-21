package models

const SwitchControllerStpSettingsPath = "switch-controller/stp-settings/"

type SwitchControllerStpSettings struct {
	ForwardTime  *float64 `json:"forward-time,omitempty"`
	HelloTime    *float64 `json:"hello-time,omitempty"`
	MaxAge       *float64 `json:"max-age,omitempty"`
	MaxHops      *float64 `json:"max-hops,omitempty"`
	Name         *string  `json:"name,omitempty"`
	PendingTimer *float64 `json:"pending-timer,omitempty"`
	Revision     *float64 `json:"revision,omitempty"`
}
