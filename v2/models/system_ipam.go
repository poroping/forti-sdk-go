package models

const SystemIpamPath = "system/ipam/"

type SystemIpam struct {
	PoolSubnet *string `json:"pool-subnet,omitempty"`
	ServerType *string `json:"server-type,omitempty"`
	Status     *string `json:"status,omitempty"`
}
