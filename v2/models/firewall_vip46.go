package models

const FirewallVip46Path = "firewall/vip46/"

type FirewallVip46 struct {
	ArpReply      *string                       `json:"arp-reply,omitempty"`
	Color         *int64                        `json:"color,omitempty"`
	Comment       *string                       `json:"comment,omitempty"`
	Extip         *string                       `json:"extip,omitempty"`
	Extport       *string                       `json:"extport,omitempty"`
	Id            *int64                        `json:"id,omitempty"`
	LdbMethod     *string                       `json:"ldb-method,omitempty"`
	Mappedip      *string                       `json:"mappedip,omitempty"`
	Mappedport    *string                       `json:"mappedport,omitempty"`
	Monitor       *[]FirewallVip46Monitor       `json:"monitor,omitempty"`
	Name          *string                       `json:"name,omitempty"`
	Portforward   *string                       `json:"portforward,omitempty"`
	Protocol      *string                       `json:"protocol,omitempty"`
	Realservers   *[]FirewallVip46Realservers   `json:"realservers,omitempty"`
	ServerType    *string                       `json:"server-type,omitempty"`
	SrcFilter     *[]FirewallVip46SrcFilter     `json:"src-filter,omitempty"`
	SrcintfFilter *[]FirewallVip46SrcintfFilter `json:"srcintf-filter,omitempty"`
	Type          *string                       `json:"type,omitempty"`
	Uuid          *string                       `json:"uuid,omitempty"`
}

type FirewallVip46Monitor struct {
	Name *string `json:"name,omitempty"`
}

type FirewallVip46Realservers struct {
	ClientIp         *string                            `json:"client-ip,omitempty"`
	Healthcheck      *string                            `json:"healthcheck,omitempty"`
	HolddownInterval *int64                             `json:"holddown-interval,omitempty"`
	Id               *int64                             `json:"id,omitempty"`
	Ip               *string                            `json:"ip,omitempty"`
	MaxConnections   *int64                             `json:"max-connections,omitempty"`
	Monitor          *[]FirewallVip46RealserversMonitor `json:"monitor,omitempty"`
	Port             *int64                             `json:"port,omitempty"`
	Status           *string                            `json:"status,omitempty"`
	Weight           *int64                             `json:"weight,omitempty"`
}

type FirewallVip46RealserversMonitor struct {
	Name *string `json:"name,omitempty"`
}

type FirewallVip46SrcFilter struct {
	Range *string `json:"range,omitempty"`
}

type FirewallVip46SrcintfFilter struct {
	InterfaceName *string `json:"interface-name,omitempty"`
}
