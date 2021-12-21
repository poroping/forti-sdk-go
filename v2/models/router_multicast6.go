package models

const RouterMulticast6Path = "router/multicast6/"

type RouterMulticast6 struct {
	Interface        *[]RouterMulticast6Interface   `json:"interface,omitempty"`
	MulticastPmtu    *string                        `json:"multicast-pmtu,omitempty"`
	MulticastRouting *string                        `json:"multicast-routing,omitempty"`
	PimSmGlobal      *[]RouterMulticast6PimSmGlobal `json:"pim-sm-global,omitempty"`
}

type RouterMulticast6Interface struct {
	HelloHoldtime *float64 `json:"hello-holdtime,omitempty"`
	HelloInterval *float64 `json:"hello-interval,omitempty"`
	Name          *string  `json:"name,omitempty"`
}

type RouterMulticast6PimSmGlobal struct {
	RegisterRateLimit *float64                                `json:"register-rate-limit,omitempty"`
	RpAddress         *[]RouterMulticast6PimSmGlobalRpAddress `json:"rp-address,omitempty"`
}

type RouterMulticast6PimSmGlobalRpAddress struct {
	Id         *float64 `json:"id,omitempty"`
	Ip6Address *string  `json:"ip6-address,omitempty"`
}
