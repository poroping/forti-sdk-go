package models

const SystemSessionHelperPath = "system/session-helper/"

type SystemSessionHelper struct {
	Fosid    *float64 `json:"fosid,omitempty"`
	Name     *string  `json:"name,omitempty"`
	Port     *float64 `json:"port,omitempty"`
	Protocol *float64 `json:"protocol,omitempty"`
}
