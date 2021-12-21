package models

const RouterospfOspfInterfacePath = "router.ospf/ospf-interface/"

type RouterospfOspfInterface struct {
	Authentication     *string                           `json:"authentication,omitempty"`
	AuthenticationKey  *string                           `json:"authentication-key,omitempty"`
	Bfd                *string                           `json:"bfd,omitempty"`
	Comments           *string                           `json:"comments,omitempty"`
	Cost               *int64                            `json:"cost,omitempty"`
	DatabaseFilterOut  *string                           `json:"database-filter-out,omitempty"`
	DeadInterval       *int64                            `json:"dead-interval,omitempty"`
	HelloInterval      *int64                            `json:"hello-interval,omitempty"`
	HelloMultiplier    *int64                            `json:"hello-multiplier,omitempty"`
	Interface          *string                           `json:"interface,omitempty"`
	Ip                 *string                           `json:"ip,omitempty"`
	Keychain           *string                           `json:"keychain,omitempty"`
	Md5Keychain        *string                           `json:"md5-keychain,omitempty"`
	Md5Keys            *[]RouterospfOspfInterfaceMd5Keys `json:"md5-keys,omitempty"`
	Mtu                *int64                            `json:"mtu,omitempty"`
	MtuIgnore          *string                           `json:"mtu-ignore,omitempty"`
	Name               *string                           `json:"name,omitempty"`
	NetworkType        *string                           `json:"network-type,omitempty"`
	PrefixLength       *int64                            `json:"prefix-length,omitempty"`
	Priority           *int64                            `json:"priority,omitempty"`
	ResyncTimeout      *int64                            `json:"resync-timeout,omitempty"`
	RetransmitInterval *int64                            `json:"retransmit-interval,omitempty"`
	Status             *string                           `json:"status,omitempty"`
	TransmitDelay      *int64                            `json:"transmit-delay,omitempty"`
}

type RouterospfOspfInterfaceMd5Keys struct {
	Id        *int64  `json:"id,omitempty"`
	KeyString *string `json:"key-string,omitempty"`
}
