package models

const SystemStpPath = "system/stp/"

type SystemStp struct {
	ForwardDelay   *int64  `json:"forward-delay,omitempty"`
	HelloTime      *int64  `json:"hello-time,omitempty"`
	MaxAge         *int64  `json:"max-age,omitempty"`
	MaxHops        *int64  `json:"max-hops,omitempty"`
	SwitchPriority *string `json:"switch-priority,omitempty"`
}
