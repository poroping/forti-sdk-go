package models

const FirewallCentralSnatMapPath = "firewall/central-snat-map/"

type FirewallCentralSnatMap struct {
	Comments   *string                             `json:"comments,omitempty"`
	DstAddr    *[]FirewallCentralSnatMapDstAddr    `json:"dst-addr,omitempty"`
	DstAddr6   *[]FirewallCentralSnatMapDstAddr6   `json:"dst-addr6,omitempty"`
	Dstintf    *[]FirewallCentralSnatMapDstintf    `json:"dstintf,omitempty"`
	Nat        *string                             `json:"nat,omitempty"`
	NatIppool  *[]FirewallCentralSnatMapNatIppool  `json:"nat-ippool,omitempty"`
	NatIppool6 *[]FirewallCentralSnatMapNatIppool6 `json:"nat-ippool6,omitempty"`
	NatPort    *string                             `json:"nat-port,omitempty"`
	Nat46      *string                             `json:"nat46,omitempty"`
	Nat64      *string                             `json:"nat64,omitempty"`
	OrigAddr   *[]FirewallCentralSnatMapOrigAddr   `json:"orig-addr,omitempty"`
	OrigAddr6  *[]FirewallCentralSnatMapOrigAddr6  `json:"orig-addr6,omitempty"`
	OrigPort   *string                             `json:"orig-port,omitempty"`
	Policyid   *float64                            `json:"policyid,omitempty"`
	Protocol   *float64                            `json:"protocol,omitempty"`
	Srcintf    *[]FirewallCentralSnatMapSrcintf    `json:"srcintf,omitempty"`
	Status     *string                             `json:"status,omitempty"`
	Type       *string                             `json:"type,omitempty"`
	Uuid       *string                             `json:"uuid,omitempty"`
}

type FirewallCentralSnatMapDstAddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallCentralSnatMapDstAddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallCentralSnatMapDstintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallCentralSnatMapNatIppool struct {
	Name *string `json:"name,omitempty"`
}

type FirewallCentralSnatMapNatIppool6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallCentralSnatMapOrigAddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallCentralSnatMapOrigAddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallCentralSnatMapSrcintf struct {
	Name *string `json:"name,omitempty"`
}
