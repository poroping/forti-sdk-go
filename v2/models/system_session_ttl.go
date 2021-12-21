package models

const SystemSessionTtlPath = "system/session-ttl/"

type SystemSessionTtl struct {
	Default *string                 `json:"default,omitempty"`
	Port    *[]SystemSessionTtlPort `json:"port,omitempty"`
}

type SystemSessionTtlPort struct {
	EndPort   *int64  `json:"end-port,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	Protocol  *int64  `json:"protocol,omitempty"`
	StartPort *int64  `json:"start-port,omitempty"`
	Timeout   *string `json:"timeout,omitempty"`
}
