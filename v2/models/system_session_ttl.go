package models

const SystemSessionTtlPath = "system/session-ttl/"

type SystemSessionTtl struct {
	Default *string                 `json:"default,omitempty"`
	Port    *[]SystemSessionTtlPort `json:"port,omitempty"`
}

const SystemSessionTtlPortPath = "system/session-ttl/port/"

type SystemSessionTtlPort struct {
	EndPort          *int64  `json:"end-port,omitempty"`
	Id               *int64  `json:"id,omitempty"`
	Protocol         *int64  `json:"protocol,omitempty"`
	RefreshDirection *string `json:"refresh-direction,omitempty"`
	StartPort        *int64  `json:"start-port,omitempty"`
	Timeout          *string `json:"timeout,omitempty"`
}
