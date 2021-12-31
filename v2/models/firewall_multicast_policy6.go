package models

const FirewallMulticastPolicy6Path = "firewall/multicast-policy6/"

type FirewallMulticastPolicy6 struct {
	Action          *string                            `json:"action,omitempty"`
	AutoAsicOffload *string                            `json:"auto-asic-offload,omitempty"`
	Comments        *string                            `json:"comments,omitempty"`
	Dstaddr         *[]FirewallMulticastPolicy6Dstaddr `json:"dstaddr,omitempty"`
	Dstintf         *string                            `json:"dstintf,omitempty"`
	EndPort         *int64                             `json:"end-port,omitempty"`
	Id              *int64                             `json:"id,omitempty"`
	Logtraffic      *string                            `json:"logtraffic,omitempty"`
	Name            *string                            `json:"name,omitempty"`
	Protocol        *int64                             `json:"protocol,omitempty"`
	Srcaddr         *[]FirewallMulticastPolicy6Srcaddr `json:"srcaddr,omitempty"`
	Srcintf         *string                            `json:"srcintf,omitempty"`
	StartPort       *int64                             `json:"start-port,omitempty"`
	Status          *string                            `json:"status,omitempty"`
	Uuid            *string                            `json:"uuid,omitempty"`
}

type FirewallMulticastPolicy6Dstaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallMulticastPolicy6Srcaddr struct {
	Name *string `json:"name,omitempty"`
}
