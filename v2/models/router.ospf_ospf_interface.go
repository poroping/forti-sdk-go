package models

const RouterospfOspfInterfacePath = "router.ospf/ospf-interface/"

type RouterospfOspfInterface struct {
	Authentication     *string                           `json:"authentication,omitempty"`
	AuthenticationKey  *string                           `json:"authentication-key,omitempty"`
	Bfd                *string                           `json:"bfd,omitempty"`
	Comments           *string                           `json:"comments,omitempty"`
	Cost               *float64                          `json:"cost,omitempty"`
	DatabaseFilterOut  *string                           `json:"database-filter-out,omitempty"`
	DeadInterval       *float64                          `json:"dead-interval,omitempty"`
	HelloInterval      *float64                          `json:"hello-interval,omitempty"`
	HelloMultiplier    *float64                          `json:"hello-multiplier,omitempty"`
	Interface          *string                           `json:"interface,omitempty"`
	Ip                 *string                           `json:"ip,omitempty"`
	Keychain           *string                           `json:"keychain,omitempty"`
	Md5Keychain        *string                           `json:"md5-keychain,omitempty"`
	Md5Keys            *[]RouterospfOspfInterfaceMd5Keys `json:"md5-keys,omitempty"`
	Mtu                *float64                          `json:"mtu,omitempty"`
	MtuIgnore          *string                           `json:"mtu-ignore,omitempty"`
	Name               *string                           `json:"name,omitempty"`
	NetworkType        *string                           `json:"network-type,omitempty"`
	PrefixLength       *float64                          `json:"prefix-length,omitempty"`
	Priority           *float64                          `json:"priority,omitempty"`
	ResyncTimeout      *float64                          `json:"resync-timeout,omitempty"`
	RetransmitInterval *float64                          `json:"retransmit-interval,omitempty"`
	Status             *string                           `json:"status,omitempty"`
	TransmitDelay      *float64                          `json:"transmit-delay,omitempty"`
}

type RouterospfOspfInterfaceMd5Keys struct {
	Id        *float64 `json:"id,omitempty"`
	KeyString *string  `json:"key-string,omitempty"`
}
