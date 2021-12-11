package models

const SystemSessionHelperPath = "system/session-helper/"

type SystemSessionHelper struct {
	Fosid    *int64  `json:"fosid,omitempty"`
	Name     *string `json:"name,omitempty"`
	Port     *int64  `json:"port,omitempty"`
	Protocol *int64  `json:"protocol,omitempty"`
}
