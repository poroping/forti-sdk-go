package models

const SystemIpamPath = "system/ipam/"

type SystemIpam struct {
	PoolSubnet *string            `json:"pool-subnet,omitempty"`
	Pools      *[]SystemIpamPools `json:"pools,omitempty"`
	Rules      *[]SystemIpamRules `json:"rules,omitempty"`
	ServerType *string            `json:"server-type,omitempty"`
	Status     *string            `json:"status,omitempty"`
}

type SystemIpamPools struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	Subnet      *string `json:"subnet,omitempty"`
}

type SystemIpamRules struct {
	Description *string                     `json:"description,omitempty"`
	Device      *[]SystemIpamRulesDevice    `json:"device,omitempty"`
	Dhcp        *string                     `json:"dhcp,omitempty"`
	Interface   *[]SystemIpamRulesInterface `json:"interface,omitempty"`
	Name        *string                     `json:"name,omitempty"`
	Pool        *[]SystemIpamRulesPool      `json:"pool,omitempty"`
	Role        *string                     `json:"role,omitempty"`
}

type SystemIpamRulesDevice struct {
	Name *string `json:"name,omitempty"`
}

type SystemIpamRulesInterface struct {
	Name *string `json:"name,omitempty"`
}

type SystemIpamRulesPool struct {
	Name *string `json:"name,omitempty"`
}
