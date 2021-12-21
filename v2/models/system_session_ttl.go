package models

const SystemSessionTtlPath = "system/session-ttl/"

type SystemSessionTtl struct {
	Default *string                 `json:"default,omitempty"`
	Port    *[]SystemSessionTtlPort `json:"port,omitempty"`
}

type SystemSessionTtlPort struct {
	EndPort   *float64 `json:"end-port,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	Protocol  *float64 `json:"protocol,omitempty"`
	StartPort *float64 `json:"start-port,omitempty"`
	Timeout   *string  `json:"timeout,omitempty"`
}
