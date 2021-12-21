package models

const SystemFssoPollingPath = "system/fsso-polling/"

type SystemFssoPolling struct {
	AuthPassword   *string  `json:"auth-password,omitempty"`
	Authentication *string  `json:"authentication,omitempty"`
	ListeningPort  *float64 `json:"listening-port,omitempty"`
	Status         *string  `json:"status,omitempty"`
}
