package models

const FirewallserviceCustomPath = "firewall.service/custom/"

type FirewallserviceCustom struct {
	AppCategory       *[]FirewallserviceCustomAppCategory `json:"app-category,omitempty"`
	AppServiceType    *string                             `json:"app-service-type,omitempty"`
	Application       *[]FirewallserviceCustomApplication `json:"application,omitempty"`
	Category          *string                             `json:"category,omitempty"`
	CheckResetRange   *string                             `json:"check-reset-range,omitempty"`
	Color             *float64                            `json:"color,omitempty"`
	Comment           *string                             `json:"comment,omitempty"`
	FabricObject      *string                             `json:"fabric-object,omitempty"`
	Fqdn              *string                             `json:"fqdn,omitempty"`
	Helper            *string                             `json:"helper,omitempty"`
	Icmpcode          *float64                            `json:"icmpcode,omitempty"`
	Icmptype          *float64                            `json:"icmptype,omitempty"`
	Iprange           *string                             `json:"iprange,omitempty"`
	Name              *string                             `json:"name,omitempty"`
	Protocol          *string                             `json:"protocol,omitempty"`
	ProtocolNumber    *float64                            `json:"protocol-number,omitempty"`
	Proxy             *string                             `json:"proxy,omitempty"`
	SctpPortrange     *string                             `json:"sctp-portrange,omitempty"`
	SessionTtl        *string                             `json:"session-ttl,omitempty"`
	TcpHalfcloseTimer *float64                            `json:"tcp-halfclose-timer,omitempty"`
	TcpHalfopenTimer  *float64                            `json:"tcp-halfopen-timer,omitempty"`
	TcpPortrange      *string                             `json:"tcp-portrange,omitempty"`
	TcpRstTimer       *float64                            `json:"tcp-rst-timer,omitempty"`
	TcpTimewaitTimer  *float64                            `json:"tcp-timewait-timer,omitempty"`
	UdpIdleTimer      *float64                            `json:"udp-idle-timer,omitempty"`
	UdpPortrange      *string                             `json:"udp-portrange,omitempty"`
	Visibility        *string                             `json:"visibility,omitempty"`
}

type FirewallserviceCustomAppCategory struct {
	Id *float64 `json:"id,omitempty"`
}

type FirewallserviceCustomApplication struct {
	Id *float64 `json:"id,omitempty"`
}
