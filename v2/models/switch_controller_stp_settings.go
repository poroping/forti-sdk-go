package models

const SwitchControllerStpSettingsPath = "switch-controller/stp-settings/"

type SwitchControllerStpSettings struct {
	ForwardTime  *int64  `json:"forward-time,omitempty"`
	HelloTime    *int64  `json:"hello-time,omitempty"`
	MaxAge       *int64  `json:"max-age,omitempty"`
	MaxHops      *int64  `json:"max-hops,omitempty"`
	Name         *string `json:"name,omitempty"`
	PendingTimer *int64  `json:"pending-timer,omitempty"`
	Revision     *int64  `json:"revision,omitempty"`
}
