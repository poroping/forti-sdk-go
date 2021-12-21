package models

const SystemResourceLimitsPath = "system/resource-limits/"

type SystemResourceLimits struct {
	CustomService        *float64 `json:"custom-service,omitempty"`
	DialupTunnel         *float64 `json:"dialup-tunnel,omitempty"`
	FirewallAddress      *float64 `json:"firewall-address,omitempty"`
	FirewallAddrgrp      *float64 `json:"firewall-addrgrp,omitempty"`
	FirewallPolicy       *float64 `json:"firewall-policy,omitempty"`
	IpsecPhase1          *float64 `json:"ipsec-phase1,omitempty"`
	IpsecPhase1Interface *float64 `json:"ipsec-phase1-interface,omitempty"`
	IpsecPhase2          *float64 `json:"ipsec-phase2,omitempty"`
	IpsecPhase2Interface *float64 `json:"ipsec-phase2-interface,omitempty"`
	LogDiskQuota         *float64 `json:"log-disk-quota,omitempty"`
	OnetimeSchedule      *float64 `json:"onetime-schedule,omitempty"`
	Proxy                *float64 `json:"proxy,omitempty"`
	RecurringSchedule    *float64 `json:"recurring-schedule,omitempty"`
	ServiceGroup         *float64 `json:"service-group,omitempty"`
	Session              *float64 `json:"session,omitempty"`
	Sslvpn               *float64 `json:"sslvpn,omitempty"`
	User                 *float64 `json:"user,omitempty"`
	UserGroup            *float64 `json:"user-group,omitempty"`
}
