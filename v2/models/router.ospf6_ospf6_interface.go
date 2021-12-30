package models

const Routerospf6Ospf6InterfacePath = "router/ospf6/ospf6-interface/"

type Routerospf6Ospf6Interface struct {
	AreaId              *string                               `json:"area-id,omitempty"`
	Authentication      *string                               `json:"authentication,omitempty"`
	Bfd                 *string                               `json:"bfd,omitempty"`
	Cost                *int64                                `json:"cost,omitempty"`
	DeadInterval        *int64                                `json:"dead-interval,omitempty"`
	HelloInterval       *int64                                `json:"hello-interval,omitempty"`
	Interface           *string                               `json:"interface,omitempty"`
	IpsecAuthAlg        *string                               `json:"ipsec-auth-alg,omitempty"`
	IpsecEncAlg         *string                               `json:"ipsec-enc-alg,omitempty"`
	IpsecKeys           *[]Routerospf6Ospf6InterfaceIpsecKeys `json:"ipsec-keys,omitempty"`
	KeyRolloverInterval *int64                                `json:"key-rollover-interval,omitempty"`
	Mtu                 *int64                                `json:"mtu,omitempty"`
	MtuIgnore           *string                               `json:"mtu-ignore,omitempty"`
	Name                *string                               `json:"name,omitempty"`
	Neighbor            *[]Routerospf6Ospf6InterfaceNeighbor  `json:"neighbor,omitempty"`
	NetworkType         *string                               `json:"network-type,omitempty"`
	Priority            *int64                                `json:"priority,omitempty"`
	RetransmitInterval  *int64                                `json:"retransmit-interval,omitempty"`
	Status              *string                               `json:"status,omitempty"`
	TransmitDelay       *int64                                `json:"transmit-delay,omitempty"`
}

type Routerospf6Ospf6InterfaceIpsecKeys struct {
	AuthKey *string `json:"auth-key,omitempty"`
	EncKey  *string `json:"enc-key,omitempty"`
	Spi     *int64  `json:"spi,omitempty"`
}

type Routerospf6Ospf6InterfaceNeighbor struct {
	Cost         *int64  `json:"cost,omitempty"`
	Ip6          *string `json:"ip6,omitempty"`
	PollInterval *int64  `json:"poll-interval,omitempty"`
	Priority     *int64  `json:"priority,omitempty"`
}