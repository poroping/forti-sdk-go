package models

const RouterbgpNeighborPath = "router.bgp/neighbor/"

type RouterbgpNeighbor struct {
	Activate                    *string                                   `json:"activate,omitempty"`
	Activate6                   *string                                   `json:"activate6,omitempty"`
	AdditionalPath              *string                                   `json:"additional-path,omitempty"`
	AdditionalPath6             *string                                   `json:"additional-path6,omitempty"`
	AdvAdditionalPath           *int64                                    `json:"adv-additional-path,omitempty"`
	AdvAdditionalPath6          *int64                                    `json:"adv-additional-path6,omitempty"`
	AdvertisementInterval       *int64                                    `json:"advertisement-interval,omitempty"`
	AllowasIn                   *int64                                    `json:"allowas-in,omitempty"`
	AllowasInEnable             *string                                   `json:"allowas-in-enable,omitempty"`
	AllowasInEnable6            *string                                   `json:"allowas-in-enable6,omitempty"`
	AllowasIn6                  *int64                                    `json:"allowas-in6,omitempty"`
	AsOverride                  *string                                   `json:"as-override,omitempty"`
	AsOverride6                 *string                                   `json:"as-override6,omitempty"`
	AttributeUnchanged          *string                                   `json:"attribute-unchanged,omitempty"`
	AttributeUnchanged6         *string                                   `json:"attribute-unchanged6,omitempty"`
	Bfd                         *string                                   `json:"bfd,omitempty"`
	CapabilityDefaultOriginate  *string                                   `json:"capability-default-originate,omitempty"`
	CapabilityDefaultOriginate6 *string                                   `json:"capability-default-originate6,omitempty"`
	CapabilityDynamic           *string                                   `json:"capability-dynamic,omitempty"`
	CapabilityGracefulRestart   *string                                   `json:"capability-graceful-restart,omitempty"`
	CapabilityGracefulRestart6  *string                                   `json:"capability-graceful-restart6,omitempty"`
	CapabilityOrf               *string                                   `json:"capability-orf,omitempty"`
	CapabilityOrf6              *string                                   `json:"capability-orf6,omitempty"`
	CapabilityRouteRefresh      *string                                   `json:"capability-route-refresh,omitempty"`
	ConditionalAdvertise        *[]RouterbgpNeighborConditionalAdvertise  `json:"conditional-advertise,omitempty"`
	ConditionalAdvertise6       *[]RouterbgpNeighborConditionalAdvertise6 `json:"conditional-advertise6,omitempty"`
	ConnectTimer                *int64                                    `json:"connect-timer,omitempty"`
	DefaultOriginateRoutemap    *string                                   `json:"default-originate-routemap,omitempty"`
	DefaultOriginateRoutemap6   *string                                   `json:"default-originate-routemap6,omitempty"`
	Description                 *string                                   `json:"description,omitempty"`
	DistributeListIn            *string                                   `json:"distribute-list-in,omitempty"`
	DistributeListIn6           *string                                   `json:"distribute-list-in6,omitempty"`
	DistributeListOut           *string                                   `json:"distribute-list-out,omitempty"`
	DistributeListOut6          *string                                   `json:"distribute-list-out6,omitempty"`
	DontCapabilityNegotiate     *string                                   `json:"dont-capability-negotiate,omitempty"`
	EbgpEnforceMultihop         *string                                   `json:"ebgp-enforce-multihop,omitempty"`
	EbgpMultihopTtl             *int64                                    `json:"ebgp-multihop-ttl,omitempty"`
	FilterListIn                *string                                   `json:"filter-list-in,omitempty"`
	FilterListIn6               *string                                   `json:"filter-list-in6,omitempty"`
	FilterListOut               *string                                   `json:"filter-list-out,omitempty"`
	FilterListOut6              *string                                   `json:"filter-list-out6,omitempty"`
	HoldtimeTimer               *int64                                    `json:"holdtime-timer,omitempty"`
	Interface                   *string                                   `json:"interface,omitempty"`
	Ip                          *string                                   `json:"ip,omitempty"`
	KeepAliveTimer              *int64                                    `json:"keep-alive-timer,omitempty"`
	LinkDownFailover            *string                                   `json:"link-down-failover,omitempty"`
	LocalAs                     *int64                                    `json:"local-as,omitempty"`
	LocalAsNoPrepend            *string                                   `json:"local-as-no-prepend,omitempty"`
	LocalAsReplaceAs            *string                                   `json:"local-as-replace-as,omitempty"`
	MaximumPrefix               *int64                                    `json:"maximum-prefix,omitempty"`
	MaximumPrefixThreshold      *int64                                    `json:"maximum-prefix-threshold,omitempty"`
	MaximumPrefixThreshold6     *int64                                    `json:"maximum-prefix-threshold6,omitempty"`
	MaximumPrefixWarningOnly    *string                                   `json:"maximum-prefix-warning-only,omitempty"`
	MaximumPrefixWarningOnly6   *string                                   `json:"maximum-prefix-warning-only6,omitempty"`
	MaximumPrefix6              *int64                                    `json:"maximum-prefix6,omitempty"`
	NextHopSelf                 *string                                   `json:"next-hop-self,omitempty"`
	NextHopSelfRr               *string                                   `json:"next-hop-self-rr,omitempty"`
	NextHopSelfRr6              *string                                   `json:"next-hop-self-rr6,omitempty"`
	NextHopSelf6                *string                                   `json:"next-hop-self6,omitempty"`
	OverrideCapability          *string                                   `json:"override-capability,omitempty"`
	Passive                     *string                                   `json:"passive,omitempty"`
	Password                    *string                                   `json:"password,omitempty"`
	PrefixListIn                *string                                   `json:"prefix-list-in,omitempty"`
	PrefixListIn6               *string                                   `json:"prefix-list-in6,omitempty"`
	PrefixListOut               *string                                   `json:"prefix-list-out,omitempty"`
	PrefixListOut6              *string                                   `json:"prefix-list-out6,omitempty"`
	RemoteAs                    *int64                                    `json:"remote-as,omitempty"`
	RemovePrivateAs             *string                                   `json:"remove-private-as,omitempty"`
	RemovePrivateAs6            *string                                   `json:"remove-private-as6,omitempty"`
	RestartTime                 *int64                                    `json:"restart-time,omitempty"`
	RetainStaleTime             *int64                                    `json:"retain-stale-time,omitempty"`
	RouteMapIn                  *string                                   `json:"route-map-in,omitempty"`
	RouteMapIn6                 *string                                   `json:"route-map-in6,omitempty"`
	RouteMapOut                 *string                                   `json:"route-map-out,omitempty"`
	RouteMapOutPreferable       *string                                   `json:"route-map-out-preferable,omitempty"`
	RouteMapOut6                *string                                   `json:"route-map-out6,omitempty"`
	RouteMapOut6Preferable      *string                                   `json:"route-map-out6-preferable,omitempty"`
	RouteReflectorClient        *string                                   `json:"route-reflector-client,omitempty"`
	RouteReflectorClient6       *string                                   `json:"route-reflector-client6,omitempty"`
	RouteServerClient           *string                                   `json:"route-server-client,omitempty"`
	RouteServerClient6          *string                                   `json:"route-server-client6,omitempty"`
	SendCommunity               *string                                   `json:"send-community,omitempty"`
	SendCommunity6              *string                                   `json:"send-community6,omitempty"`
	Shutdown                    *string                                   `json:"shutdown,omitempty"`
	SoftReconfiguration         *string                                   `json:"soft-reconfiguration,omitempty"`
	SoftReconfiguration6        *string                                   `json:"soft-reconfiguration6,omitempty"`
	StaleRoute                  *string                                   `json:"stale-route,omitempty"`
	StrictCapabilityMatch       *string                                   `json:"strict-capability-match,omitempty"`
	UnsuppressMap               *string                                   `json:"unsuppress-map,omitempty"`
	UnsuppressMap6              *string                                   `json:"unsuppress-map6,omitempty"`
	UpdateSource                *string                                   `json:"update-source,omitempty"`
	Weight                      *int64                                    `json:"weight,omitempty"`
}

type RouterbgpNeighborConditionalAdvertise struct {
	AdvertiseRoutemap *string `json:"advertise-routemap,omitempty"`
	ConditionRoutemap *string `json:"condition-routemap,omitempty"`
	ConditionType     *string `json:"condition-type,omitempty"`
}

type RouterbgpNeighborConditionalAdvertise6 struct {
	AdvertiseRoutemap *string `json:"advertise-routemap,omitempty"`
	ConditionRoutemap *string `json:"condition-routemap,omitempty"`
	ConditionType     *string `json:"condition-type,omitempty"`
}
