package models

const FirewallIpv6EhFilterPath = "firewall/ipv6-eh-filter/"

type FirewallIpv6EhFilter struct {
	Auth        *string `json:"auth,omitempty"`
	DestOpt     *string `json:"dest-opt,omitempty"`
	Fragment    *string `json:"fragment,omitempty"`
	HdoptType   *int64  `json:"hdopt-type,omitempty"`
	HopOpt      *string `json:"hop-opt,omitempty"`
	NoNext      *string `json:"no-next,omitempty"`
	Routing     *string `json:"routing,omitempty"`
	RoutingType *int64  `json:"routing-type,omitempty"`
}
