package models

const WebProxyForwardServerGroupPath = "web-proxy/forward-server-group/"

type WebProxyForwardServerGroup struct {
	Affinity        *string                                `json:"affinity,omitempty"`
	GroupDownOption *string                                `json:"group-down-option,omitempty"`
	LdbMethod       *string                                `json:"ldb-method,omitempty"`
	Name            *string                                `json:"name,omitempty"`
	ServerList      []WebProxyForwardServerGroupServerList `json:"server-list,omitempty"`
}

type WebProxyForwardServerGroupServerList struct {
	Name   *string `json:"name,omitempty"`
	Weight *int64  `json:"weight,omitempty"`
}
