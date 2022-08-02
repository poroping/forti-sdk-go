package models

const FirewallVip64Path = "firewall/vip64/"

type FirewallVip64 struct {
	ArpReply    *string                     `json:"arp-reply,omitempty"`
	Color       *int64                      `json:"color,omitempty"`
	Comment     *string                     `json:"comment,omitempty"`
	Extip       *string                     `json:"extip,omitempty"`
	Extport     *string                     `json:"extport,omitempty"`
	Id          *int64                      `json:"id,omitempty"`
	LdbMethod   *string                     `json:"ldb-method,omitempty"`
	Mappedip    *string                     `json:"mappedip,omitempty"`
	Mappedport  *string                     `json:"mappedport,omitempty"`
	Monitor     *[]FirewallVip64Monitor     `json:"monitor,omitempty"`
	Name        *string                     `json:"name,omitempty"`
	Portforward *string                     `json:"portforward,omitempty"`
	Protocol    *string                     `json:"protocol,omitempty"`
	Realservers *[]FirewallVip64Realservers `json:"realservers,omitempty"`
	ServerType  *string                     `json:"server-type,omitempty"`
	SrcFilter   *[]FirewallVip64SrcFilter   `json:"src-filter,omitempty"`
	Type        *string                     `json:"type,omitempty"`
	Uuid        *string                     `json:"uuid,omitempty"`
}

const FirewallVip64MonitorPath = "firewall/vip64/monitor/"

type FirewallVip64Monitor struct {
	Name *string `json:"name,omitempty"`
}

const FirewallVip64RealserversPath = "firewall/vip64/realservers/"

type FirewallVip64Realservers struct {
	ClientIp         *string                            `json:"client-ip,omitempty"`
	Healthcheck      *string                            `json:"healthcheck,omitempty"`
	HolddownInterval *int64                             `json:"holddown-interval,omitempty"`
	Id               *int64                             `json:"id,omitempty"`
	Ip               *string                            `json:"ip,omitempty"`
	MaxConnections   *int64                             `json:"max-connections,omitempty"`
	Monitor          *[]FirewallVip64RealserversMonitor `json:"monitor,omitempty"`
	Port             *int64                             `json:"port,omitempty"`
	Status           *string                            `json:"status,omitempty"`
	Weight           *int64                             `json:"weight,omitempty"`
}

const FirewallVip64RealserversMonitorPath = "firewall/vip64/realservers/monitor/"

type FirewallVip64RealserversMonitor struct {
	Name *string `json:"name,omitempty"`
}

const FirewallVip64SrcFilterPath = "firewall/vip64/src-filter/"

type FirewallVip64SrcFilter struct {
	Range *string `json:"range,omitempty"`
}
