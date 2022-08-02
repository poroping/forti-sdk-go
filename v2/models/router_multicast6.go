package models

const RouterMulticast6Path = "router/multicast6/"

type RouterMulticast6 struct {
	Interface        *[]RouterMulticast6Interface `json:"interface,omitempty"`
	MulticastPmtu    *string                      `json:"multicast-pmtu,omitempty"`
	MulticastRouting *string                      `json:"multicast-routing,omitempty"`
	PimSmGlobal      *RouterMulticast6PimSmGlobal `json:"pim-sm-global,omitempty"`
}

const RouterMulticast6InterfacePath = "router/multicast6/interface/"

type RouterMulticast6Interface struct {
	HelloHoldtime *int64  `json:"hello-holdtime,omitempty"`
	HelloInterval *int64  `json:"hello-interval,omitempty"`
	Name          *string `json:"name,omitempty"`
}

const RouterMulticast6PimSmGlobalPath = "router/multicast6/pim-sm-global/"

type RouterMulticast6PimSmGlobal struct {
	RegisterRateLimit *int64                                  `json:"register-rate-limit,omitempty"`
	RpAddress         *[]RouterMulticast6PimSmGlobalRpAddress `json:"rp-address,omitempty"`
}

const RouterMulticast6PimSmGlobalRpAddressPath = "router/multicast6/pim-sm-global/rp-address/"

type RouterMulticast6PimSmGlobalRpAddress struct {
	Id         *int64  `json:"id,omitempty"`
	Ip6Address *string `json:"ip6-address,omitempty"`
}
