package models

const WebProxyWispPath = "web-proxy/wisp/"

type WebProxyWisp struct {
	Comment        *string  `json:"comment,omitempty"`
	MaxConnections *float64 `json:"max-connections,omitempty"`
	Name           *string  `json:"name,omitempty"`
	OutgoingIp     *string  `json:"outgoing-ip,omitempty"`
	ServerIp       *string  `json:"server-ip,omitempty"`
	ServerPort     *float64 `json:"server-port,omitempty"`
	Timeout        *float64 `json:"timeout,omitempty"`
}
