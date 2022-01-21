package models

const SystemSessionHelperPath = "system/session-helper/"

type SystemSessionHelper struct {
	Id       *int64  `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Port     *int64  `json:"port,omitempty"`
	Protocol *int64  `json:"protocol,omitempty"`
}
