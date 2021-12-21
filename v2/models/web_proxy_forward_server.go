package models

const WebProxyForwardServerPath = "web-proxy/forward-server/"

type WebProxyForwardServer struct {
	AddrType         *string  `json:"addr-type,omitempty"`
	Comment          *string  `json:"comment,omitempty"`
	Fqdn             *string  `json:"fqdn,omitempty"`
	Healthcheck      *string  `json:"healthcheck,omitempty"`
	Ip               *string  `json:"ip,omitempty"`
	Monitor          *string  `json:"monitor,omitempty"`
	Name             *string  `json:"name,omitempty"`
	Password         *string  `json:"password,omitempty"`
	Port             *float64 `json:"port,omitempty"`
	ServerDownOption *string  `json:"server-down-option,omitempty"`
	Username         *string  `json:"username,omitempty"`
}
