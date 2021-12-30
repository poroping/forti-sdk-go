package models

const WebProxyWispPath = "web-proxy/wisp/"

type WebProxyWisp struct {
	Comment        *string `json:"comment,omitempty"`
	MaxConnections *int64  `json:"max-connections,omitempty"`
	Name           *string `json:"name,omitempty"`
	OutgoingIp     *string `json:"outgoing-ip,omitempty"`
	ServerIp       *string `json:"server-ip,omitempty"`
	ServerPort     *int64  `json:"server-port,omitempty"`
	Timeout        *int64  `json:"timeout,omitempty"`
}
