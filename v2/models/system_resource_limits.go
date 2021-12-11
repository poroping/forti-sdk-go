package models

const SystemResourceLimitsPath = "system/resource-limits/"

type SystemResourceLimits struct {
	CustomService        *int64 `json:"custom-service,omitempty"`
	DialupTunnel         *int64 `json:"dialup-tunnel,omitempty"`
	FirewallAddress      *int64 `json:"firewall-address,omitempty"`
	FirewallAddrgrp      *int64 `json:"firewall-addrgrp,omitempty"`
	FirewallPolicy       *int64 `json:"firewall-policy,omitempty"`
	IpsecPhase1          *int64 `json:"ipsec-phase1,omitempty"`
	IpsecPhase1Interface *int64 `json:"ipsec-phase1-interface,omitempty"`
	IpsecPhase2          *int64 `json:"ipsec-phase2,omitempty"`
	IpsecPhase2Interface *int64 `json:"ipsec-phase2-interface,omitempty"`
	LogDiskQuota         *int64 `json:"log-disk-quota,omitempty"`
	OnetimeSchedule      *int64 `json:"onetime-schedule,omitempty"`
	Proxy                *int64 `json:"proxy,omitempty"`
	RecurringSchedule    *int64 `json:"recurring-schedule,omitempty"`
	ServiceGroup         *int64 `json:"service-group,omitempty"`
	Session              *int64 `json:"session,omitempty"`
	Sslvpn               *int64 `json:"sslvpn,omitempty"`
	User                 *int64 `json:"user,omitempty"`
	UserGroup            *int64 `json:"user-group,omitempty"`
}
