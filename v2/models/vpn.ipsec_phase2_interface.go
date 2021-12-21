package models

const VpnipsecPhase2InterfacePath = "vpn.ipsec/phase2-interface/"

type VpnipsecPhase2Interface struct {
	AddRoute               *string  `json:"add-route,omitempty"`
	AutoDiscoveryForwarder *string  `json:"auto-discovery-forwarder,omitempty"`
	AutoDiscoverySender    *string  `json:"auto-discovery-sender,omitempty"`
	AutoNegotiate          *string  `json:"auto-negotiate,omitempty"`
	Comments               *string  `json:"comments,omitempty"`
	DhcpIpsec              *string  `json:"dhcp-ipsec,omitempty"`
	Dhgrp                  *string  `json:"dhgrp,omitempty"`
	Diffserv               *string  `json:"diffserv,omitempty"`
	Diffservcode           *string  `json:"diffservcode,omitempty"`
	DstAddrType            *string  `json:"dst-addr-type,omitempty"`
	DstEndIp               *string  `json:"dst-end-ip,omitempty"`
	DstEndIp6              *string  `json:"dst-end-ip6,omitempty"`
	DstName                *string  `json:"dst-name,omitempty"`
	DstName6               *string  `json:"dst-name6,omitempty"`
	DstPort                *float64 `json:"dst-port,omitempty"`
	DstStartIp             *string  `json:"dst-start-ip,omitempty"`
	DstStartIp6            *string  `json:"dst-start-ip6,omitempty"`
	DstSubnet              *string  `json:"dst-subnet,omitempty"`
	DstSubnet6             *string  `json:"dst-subnet6,omitempty"`
	Encapsulation          *string  `json:"encapsulation,omitempty"`
	InitiatorTsNarrow      *string  `json:"initiator-ts-narrow,omitempty"`
	Ipv4Df                 *string  `json:"ipv4-df,omitempty"`
	Keepalive              *string  `json:"keepalive,omitempty"`
	KeylifeType            *string  `json:"keylife-type,omitempty"`
	Keylifekbs             *float64 `json:"keylifekbs,omitempty"`
	Keylifeseconds         *float64 `json:"keylifeseconds,omitempty"`
	L2tp                   *string  `json:"l2tp,omitempty"`
	Name                   *string  `json:"name,omitempty"`
	Pfs                    *string  `json:"pfs,omitempty"`
	Phase1name             *string  `json:"phase1name,omitempty"`
	Proposal               *string  `json:"proposal,omitempty"`
	Protocol               *float64 `json:"protocol,omitempty"`
	Replay                 *string  `json:"replay,omitempty"`
	RouteOverlap           *string  `json:"route-overlap,omitempty"`
	SingleSource           *string  `json:"single-source,omitempty"`
	SrcAddrType            *string  `json:"src-addr-type,omitempty"`
	SrcEndIp               *string  `json:"src-end-ip,omitempty"`
	SrcEndIp6              *string  `json:"src-end-ip6,omitempty"`
	SrcName                *string  `json:"src-name,omitempty"`
	SrcName6               *string  `json:"src-name6,omitempty"`
	SrcPort                *float64 `json:"src-port,omitempty"`
	SrcStartIp             *string  `json:"src-start-ip,omitempty"`
	SrcStartIp6            *string  `json:"src-start-ip6,omitempty"`
	SrcSubnet              *string  `json:"src-subnet,omitempty"`
	SrcSubnet6             *string  `json:"src-subnet6,omitempty"`
}
