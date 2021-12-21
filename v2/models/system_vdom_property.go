package models

const SystemVdomPropertyPath = "system/vdom-property/"

type SystemVdomProperty struct {
	CustomService        *string  `json:"custom-service,omitempty"`
	Description          *string  `json:"description,omitempty"`
	DialupTunnel         *string  `json:"dialup-tunnel,omitempty"`
	FirewallAddress      *string  `json:"firewall-address,omitempty"`
	FirewallAddrgrp      *string  `json:"firewall-addrgrp,omitempty"`
	FirewallPolicy       *string  `json:"firewall-policy,omitempty"`
	IpsecPhase1          *string  `json:"ipsec-phase1,omitempty"`
	IpsecPhase1Interface *string  `json:"ipsec-phase1-interface,omitempty"`
	IpsecPhase2          *string  `json:"ipsec-phase2,omitempty"`
	IpsecPhase2Interface *string  `json:"ipsec-phase2-interface,omitempty"`
	LogDiskQuota         *string  `json:"log-disk-quota,omitempty"`
	Name                 *string  `json:"name,omitempty"`
	OnetimeSchedule      *string  `json:"onetime-schedule,omitempty"`
	Proxy                *string  `json:"proxy,omitempty"`
	RecurringSchedule    *string  `json:"recurring-schedule,omitempty"`
	ServiceGroup         *string  `json:"service-group,omitempty"`
	Session              *string  `json:"session,omitempty"`
	SnmpIndex            *float64 `json:"snmp-index,omitempty"`
	Sslvpn               *string  `json:"sslvpn,omitempty"`
	User                 *string  `json:"user,omitempty"`
	UserGroup            *string  `json:"user-group,omitempty"`
}
