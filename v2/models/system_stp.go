package models

const SystemStpPath = "system/stp/"

type SystemStp struct {
	ForwardDelay   *float64 `json:"forward-delay,omitempty"`
	HelloTime      *float64 `json:"hello-time,omitempty"`
	MaxAge         *float64 `json:"max-age,omitempty"`
	MaxHops        *float64 `json:"max-hops,omitempty"`
	SwitchPriority *string  `json:"switch-priority,omitempty"`
}
