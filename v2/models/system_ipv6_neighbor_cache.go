package models

const SystemIpv6NeighborCachePath = "system/ipv6-neighbor-cache/"

type SystemIpv6NeighborCache struct {
	Id        *int64  `json:"id,omitempty"`
	Interface *string `json:"interface,omitempty"`
	Ipv6      *string `json:"ipv6,omitempty"`
	Mac       *string `json:"mac,omitempty"`
}
