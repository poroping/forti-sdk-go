package models

const RouterBgpPath = "router/bgp/"

type RouterBgp struct {
	AdditionalPath                  *string                        `json:"additional-path,omitempty"`
	AdditionalPathSelect            *int64                         `json:"additional-path-select,omitempty"`
	AdditionalPathSelectVpnv4       *int64                         `json:"additional-path-select-vpnv4,omitempty"`
	AdditionalPathSelect6           *int64                         `json:"additional-path-select6,omitempty"`
	AdditionalPathVpnv4             *string                        `json:"additional-path-vpnv4,omitempty"`
	AdditionalPath6                 *string                        `json:"additional-path6,omitempty"`
	AdminDistance                   *[]RouterBgpAdminDistance      `json:"admin-distance,omitempty"`
	AggregateAddress                *[]RouterBgpAggregateAddress   `json:"aggregate-address,omitempty"`
	AggregateAddress6               *[]RouterBgpAggregateAddress6  `json:"aggregate-address6,omitempty"`
	AlwaysCompareMed                *string                        `json:"always-compare-med,omitempty"`
	As                              *int64                         `json:"as,omitempty"`
	BestpathAsPathIgnore            *string                        `json:"bestpath-as-path-ignore,omitempty"`
	BestpathCmpConfedAspath         *string                        `json:"bestpath-cmp-confed-aspath,omitempty"`
	BestpathCmpRouterid             *string                        `json:"bestpath-cmp-routerid,omitempty"`
	BestpathMedConfed               *string                        `json:"bestpath-med-confed,omitempty"`
	BestpathMedMissingAsWorst       *string                        `json:"bestpath-med-missing-as-worst,omitempty"`
	ClientToClientReflection        *string                        `json:"client-to-client-reflection,omitempty"`
	ClusterId                       *string                        `json:"cluster-id,omitempty"`
	ConfederationIdentifier         *int64                         `json:"confederation-identifier,omitempty"`
	ConfederationPeers              *[]RouterBgpConfederationPeers `json:"confederation-peers,omitempty"`
	Dampening                       *string                        `json:"dampening,omitempty"`
	DampeningMaxSuppressTime        *int64                         `json:"dampening-max-suppress-time,omitempty"`
	DampeningReachabilityHalfLife   *int64                         `json:"dampening-reachability-half-life,omitempty"`
	DampeningReuse                  *int64                         `json:"dampening-reuse,omitempty"`
	DampeningRouteMap               *string                        `json:"dampening-route-map,omitempty"`
	DampeningSuppress               *int64                         `json:"dampening-suppress,omitempty"`
	DampeningUnreachabilityHalfLife *int64                         `json:"dampening-unreachability-half-life,omitempty"`
	DefaultLocalPreference          *int64                         `json:"default-local-preference,omitempty"`
	DeterministicMed                *string                        `json:"deterministic-med,omitempty"`
	DistanceExternal                *int64                         `json:"distance-external,omitempty"`
	DistanceInternal                *int64                         `json:"distance-internal,omitempty"`
	DistanceLocal                   *int64                         `json:"distance-local,omitempty"`
	EbgpMultipath                   *string                        `json:"ebgp-multipath,omitempty"`
	EnforceFirstAs                  *string                        `json:"enforce-first-as,omitempty"`
	FastExternalFailover            *string                        `json:"fast-external-failover,omitempty"`
	GracefulEndOnTimer              *string                        `json:"graceful-end-on-timer,omitempty"`
	GracefulRestart                 *string                        `json:"graceful-restart,omitempty"`
	GracefulRestartTime             *int64                         `json:"graceful-restart-time,omitempty"`
	GracefulStalepathTime           *int64                         `json:"graceful-stalepath-time,omitempty"`
	GracefulUpdateDelay             *int64                         `json:"graceful-update-delay,omitempty"`
	HoldtimeTimer                   *int64                         `json:"holdtime-timer,omitempty"`
	IbgpMultipath                   *string                        `json:"ibgp-multipath,omitempty"`
	IgnoreOptionalCapability        *string                        `json:"ignore-optional-capability,omitempty"`
	KeepaliveTimer                  *int64                         `json:"keepalive-timer,omitempty"`
	LogNeighbourChanges             *string                        `json:"log-neighbour-changes,omitempty"`
	MultipathRecursiveDistance      *string                        `json:"multipath-recursive-distance,omitempty"`
	Neighbor                        *[]RouterBgpNeighbor           `json:"neighbor,omitempty"`
	NeighborGroup                   *[]RouterBgpNeighborGroup      `json:"neighbor-group,omitempty"`
	NeighborRange                   *[]RouterBgpNeighborRange      `json:"neighbor-range,omitempty"`
	NeighborRange6                  *[]RouterBgpNeighborRange6     `json:"neighbor-range6,omitempty"`
	Network                         *[]RouterBgpNetwork            `json:"network,omitempty"`
	NetworkImportCheck              *string                        `json:"network-import-check,omitempty"`
	Network6                        *[]RouterBgpNetwork6           `json:"network6,omitempty"`
	RecursiveNextHop                *string                        `json:"recursive-next-hop,omitempty"`
	Redistribute                    *[]RouterBgpRedistribute       `json:"redistribute,omitempty"`
	Redistribute6                   *[]RouterBgpRedistribute6      `json:"redistribute6,omitempty"`
	RouterId                        *string                        `json:"router-id,omitempty"`
	ScanTime                        *int64                         `json:"scan-time,omitempty"`
	Synchronization                 *string                        `json:"synchronization,omitempty"`
	TagResolveMode                  *string                        `json:"tag-resolve-mode,omitempty"`
	Vrf                             *[]RouterBgpVrf                `json:"vrf,omitempty"`
	VrfLeak                         *[]RouterBgpVrfLeak            `json:"vrf-leak,omitempty"`
	VrfLeak6                        *[]RouterBgpVrfLeak6           `json:"vrf-leak6,omitempty"`
	Vrf6                            *[]RouterBgpVrf6               `json:"vrf6,omitempty"`
}

type RouterBgpAdminDistance struct {
	Distance        *int64  `json:"distance,omitempty"`
	Id              *int64  `json:"id,omitempty"`
	NeighbourPrefix *string `json:"neighbour-prefix,omitempty"`
	RouteList       *string `json:"route-list,omitempty"`
}

type RouterBgpAggregateAddress struct {
	AsSet       *string `json:"as-set,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	Prefix      *string `json:"prefix,omitempty"`
	SummaryOnly *string `json:"summary-only,omitempty"`
}

type RouterBgpAggregateAddress6 struct {
	AsSet       *string `json:"as-set,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	Prefix6     *string `json:"prefix6,omitempty"`
	SummaryOnly *string `json:"summary-only,omitempty"`
}

type RouterBgpConfederationPeers struct {
	Peer *string `json:"peer,omitempty"`
}

type RouterBgpNeighbor struct {
	Activate                       *string                                   `json:"activate,omitempty"`
	ActivateVpnv4                  *string                                   `json:"activate-vpnv4,omitempty"`
	Activate6                      *string                                   `json:"activate6,omitempty"`
	AdditionalPath                 *string                                   `json:"additional-path,omitempty"`
	AdditionalPathVpnv4            *string                                   `json:"additional-path-vpnv4,omitempty"`
	AdditionalPath6                *string                                   `json:"additional-path6,omitempty"`
	AdvAdditionalPath              *int64                                    `json:"adv-additional-path,omitempty"`
	AdvAdditionalPathVpnv4         *int64                                    `json:"adv-additional-path-vpnv4,omitempty"`
	AdvAdditionalPath6             *int64                                    `json:"adv-additional-path6,omitempty"`
	AdvertisementInterval          *int64                                    `json:"advertisement-interval,omitempty"`
	AllowasIn                      *int64                                    `json:"allowas-in,omitempty"`
	AllowasInEnable                *string                                   `json:"allowas-in-enable,omitempty"`
	AllowasInEnable6               *string                                   `json:"allowas-in-enable6,omitempty"`
	AllowasInVpnv4                 *int64                                    `json:"allowas-in-vpnv4,omitempty"`
	AllowasIn6                     *int64                                    `json:"allowas-in6,omitempty"`
	AsOverride                     *string                                   `json:"as-override,omitempty"`
	AsOverride6                    *string                                   `json:"as-override6,omitempty"`
	AttributeUnchanged             *string                                   `json:"attribute-unchanged,omitempty"`
	AttributeUnchangedVpnv4        *string                                   `json:"attribute-unchanged-vpnv4,omitempty"`
	AttributeUnchanged6            *string                                   `json:"attribute-unchanged6,omitempty"`
	Bfd                            *string                                   `json:"bfd,omitempty"`
	CapabilityDefaultOriginate     *string                                   `json:"capability-default-originate,omitempty"`
	CapabilityDefaultOriginate6    *string                                   `json:"capability-default-originate6,omitempty"`
	CapabilityDynamic              *string                                   `json:"capability-dynamic,omitempty"`
	CapabilityGracefulRestart      *string                                   `json:"capability-graceful-restart,omitempty"`
	CapabilityGracefulRestartVpnv4 *string                                   `json:"capability-graceful-restart-vpnv4,omitempty"`
	CapabilityGracefulRestart6     *string                                   `json:"capability-graceful-restart6,omitempty"`
	CapabilityOrf                  *string                                   `json:"capability-orf,omitempty"`
	CapabilityOrf6                 *string                                   `json:"capability-orf6,omitempty"`
	CapabilityRouteRefresh         *string                                   `json:"capability-route-refresh,omitempty"`
	ConditionalAdvertise           *[]RouterBgpNeighborConditionalAdvertise  `json:"conditional-advertise,omitempty"`
	ConditionalAdvertise6          *[]RouterBgpNeighborConditionalAdvertise6 `json:"conditional-advertise6,omitempty"`
	ConnectTimer                   *int64                                    `json:"connect-timer,omitempty"`
	DefaultOriginateRoutemap       *string                                   `json:"default-originate-routemap,omitempty"`
	DefaultOriginateRoutemap6      *string                                   `json:"default-originate-routemap6,omitempty"`
	Description                    *string                                   `json:"description,omitempty"`
	DistributeListIn               *string                                   `json:"distribute-list-in,omitempty"`
	DistributeListInVpnv4          *string                                   `json:"distribute-list-in-vpnv4,omitempty"`
	DistributeListIn6              *string                                   `json:"distribute-list-in6,omitempty"`
	DistributeListOut              *string                                   `json:"distribute-list-out,omitempty"`
	DistributeListOutVpnv4         *string                                   `json:"distribute-list-out-vpnv4,omitempty"`
	DistributeListOut6             *string                                   `json:"distribute-list-out6,omitempty"`
	DontCapabilityNegotiate        *string                                   `json:"dont-capability-negotiate,omitempty"`
	EbgpEnforceMultihop            *string                                   `json:"ebgp-enforce-multihop,omitempty"`
	EbgpMultihopTtl                *int64                                    `json:"ebgp-multihop-ttl,omitempty"`
	FilterListIn                   *string                                   `json:"filter-list-in,omitempty"`
	FilterListIn6                  *string                                   `json:"filter-list-in6,omitempty"`
	FilterListOut                  *string                                   `json:"filter-list-out,omitempty"`
	FilterListOut6                 *string                                   `json:"filter-list-out6,omitempty"`
	HoldtimeTimer                  *int64                                    `json:"holdtime-timer,omitempty"`
	Interface                      *string                                   `json:"interface,omitempty"`
	Ip                             *string                                   `json:"ip,omitempty"`
	KeepAliveTimer                 *int64                                    `json:"keep-alive-timer,omitempty"`
	LinkDownFailover               *string                                   `json:"link-down-failover,omitempty"`
	LocalAs                        *int64                                    `json:"local-as,omitempty"`
	LocalAsNoPrepend               *string                                   `json:"local-as-no-prepend,omitempty"`
	LocalAsReplaceAs               *string                                   `json:"local-as-replace-as,omitempty"`
	MaximumPrefix                  *int64                                    `json:"maximum-prefix,omitempty"`
	MaximumPrefixThreshold         *int64                                    `json:"maximum-prefix-threshold,omitempty"`
	MaximumPrefixThresholdVpnv4    *int64                                    `json:"maximum-prefix-threshold-vpnv4,omitempty"`
	MaximumPrefixThreshold6        *int64                                    `json:"maximum-prefix-threshold6,omitempty"`
	MaximumPrefixVpnv4             *int64                                    `json:"maximum-prefix-vpnv4,omitempty"`
	MaximumPrefixWarningOnly       *string                                   `json:"maximum-prefix-warning-only,omitempty"`
	MaximumPrefixWarningOnlyVpnv4  *string                                   `json:"maximum-prefix-warning-only-vpnv4,omitempty"`
	MaximumPrefixWarningOnly6      *string                                   `json:"maximum-prefix-warning-only6,omitempty"`
	MaximumPrefix6                 *int64                                    `json:"maximum-prefix6,omitempty"`
	NextHopSelf                    *string                                   `json:"next-hop-self,omitempty"`
	NextHopSelfRr                  *string                                   `json:"next-hop-self-rr,omitempty"`
	NextHopSelfRr6                 *string                                   `json:"next-hop-self-rr6,omitempty"`
	NextHopSelfVpnv4               *string                                   `json:"next-hop-self-vpnv4,omitempty"`
	NextHopSelf6                   *string                                   `json:"next-hop-self6,omitempty"`
	OverrideCapability             *string                                   `json:"override-capability,omitempty"`
	Passive                        *string                                   `json:"passive,omitempty"`
	Password                       *string                                   `json:"password,omitempty"`
	PrefixListIn                   *string                                   `json:"prefix-list-in,omitempty"`
	PrefixListInVpnv4              *string                                   `json:"prefix-list-in-vpnv4,omitempty"`
	PrefixListIn6                  *string                                   `json:"prefix-list-in6,omitempty"`
	PrefixListOut                  *string                                   `json:"prefix-list-out,omitempty"`
	PrefixListOutVpnv4             *string                                   `json:"prefix-list-out-vpnv4,omitempty"`
	PrefixListOut6                 *string                                   `json:"prefix-list-out6,omitempty"`
	RemoteAs                       *int64                                    `json:"remote-as,omitempty"`
	RemovePrivateAs                *string                                   `json:"remove-private-as,omitempty"`
	RemovePrivateAsVpnv4           *string                                   `json:"remove-private-as-vpnv4,omitempty"`
	RemovePrivateAs6               *string                                   `json:"remove-private-as6,omitempty"`
	RestartTime                    *int64                                    `json:"restart-time,omitempty"`
	RetainStaleTime                *int64                                    `json:"retain-stale-time,omitempty"`
	RouteMapIn                     *string                                   `json:"route-map-in,omitempty"`
	RouteMapInVpnv4                *string                                   `json:"route-map-in-vpnv4,omitempty"`
	RouteMapIn6                    *string                                   `json:"route-map-in6,omitempty"`
	RouteMapOut                    *string                                   `json:"route-map-out,omitempty"`
	RouteMapOutPreferable          *string                                   `json:"route-map-out-preferable,omitempty"`
	RouteMapOutVpnv4               *string                                   `json:"route-map-out-vpnv4,omitempty"`
	RouteMapOutVpnv4Preferable     *string                                   `json:"route-map-out-vpnv4-preferable,omitempty"`
	RouteMapOut6                   *string                                   `json:"route-map-out6,omitempty"`
	RouteMapOut6Preferable         *string                                   `json:"route-map-out6-preferable,omitempty"`
	RouteReflectorClient           *string                                   `json:"route-reflector-client,omitempty"`
	RouteReflectorClientVpnv4      *string                                   `json:"route-reflector-client-vpnv4,omitempty"`
	RouteReflectorClient6          *string                                   `json:"route-reflector-client6,omitempty"`
	RouteServerClient              *string                                   `json:"route-server-client,omitempty"`
	RouteServerClientVpnv4         *string                                   `json:"route-server-client-vpnv4,omitempty"`
	RouteServerClient6             *string                                   `json:"route-server-client6,omitempty"`
	SendCommunity                  *string                                   `json:"send-community,omitempty"`
	SendCommunityVpnv4             *string                                   `json:"send-community-vpnv4,omitempty"`
	SendCommunity6                 *string                                   `json:"send-community6,omitempty"`
	Shutdown                       *string                                   `json:"shutdown,omitempty"`
	SoftReconfiguration            *string                                   `json:"soft-reconfiguration,omitempty"`
	SoftReconfigurationVpnv4       *string                                   `json:"soft-reconfiguration-vpnv4,omitempty"`
	SoftReconfiguration6           *string                                   `json:"soft-reconfiguration6,omitempty"`
	StaleRoute                     *string                                   `json:"stale-route,omitempty"`
	StrictCapabilityMatch          *string                                   `json:"strict-capability-match,omitempty"`
	UnsuppressMap                  *string                                   `json:"unsuppress-map,omitempty"`
	UnsuppressMap6                 *string                                   `json:"unsuppress-map6,omitempty"`
	UpdateSource                   *string                                   `json:"update-source,omitempty"`
	Weight                         *int64                                    `json:"weight,omitempty"`
}

type RouterBgpNeighborConditionalAdvertise struct {
	AdvertiseRoutemap *string                                                   `json:"advertise-routemap,omitempty"`
	ConditionRoutemap *[]RouterBgpNeighborConditionalAdvertiseConditionRoutemap `json:"condition-routemap,omitempty"`
	ConditionType     *string                                                   `json:"condition-type,omitempty"`
}

type RouterBgpNeighborConditionalAdvertiseConditionRoutemap struct {
	Name *string `json:"name,omitempty"`
}

type RouterBgpNeighborConditionalAdvertise6 struct {
	AdvertiseRoutemap *string                                                    `json:"advertise-routemap,omitempty"`
	ConditionRoutemap *[]RouterBgpNeighborConditionalAdvertise6ConditionRoutemap `json:"condition-routemap,omitempty"`
	ConditionType     *string                                                    `json:"condition-type,omitempty"`
}

type RouterBgpNeighborConditionalAdvertise6ConditionRoutemap struct {
	Name *string `json:"name,omitempty"`
}

type RouterBgpNeighborGroup struct {
	Activate                       *string `json:"activate,omitempty"`
	ActivateVpnv4                  *string `json:"activate-vpnv4,omitempty"`
	Activate6                      *string `json:"activate6,omitempty"`
	AdditionalPath                 *string `json:"additional-path,omitempty"`
	AdditionalPathVpnv4            *string `json:"additional-path-vpnv4,omitempty"`
	AdditionalPath6                *string `json:"additional-path6,omitempty"`
	AdvAdditionalPath              *int64  `json:"adv-additional-path,omitempty"`
	AdvAdditionalPathVpnv4         *int64  `json:"adv-additional-path-vpnv4,omitempty"`
	AdvAdditionalPath6             *int64  `json:"adv-additional-path6,omitempty"`
	AdvertisementInterval          *int64  `json:"advertisement-interval,omitempty"`
	AllowasIn                      *int64  `json:"allowas-in,omitempty"`
	AllowasInEnable                *string `json:"allowas-in-enable,omitempty"`
	AllowasInEnable6               *string `json:"allowas-in-enable6,omitempty"`
	AllowasInVpnv4                 *int64  `json:"allowas-in-vpnv4,omitempty"`
	AllowasIn6                     *int64  `json:"allowas-in6,omitempty"`
	AsOverride                     *string `json:"as-override,omitempty"`
	AsOverride6                    *string `json:"as-override6,omitempty"`
	AttributeUnchanged             *string `json:"attribute-unchanged,omitempty"`
	AttributeUnchangedVpnv4        *string `json:"attribute-unchanged-vpnv4,omitempty"`
	AttributeUnchanged6            *string `json:"attribute-unchanged6,omitempty"`
	Bfd                            *string `json:"bfd,omitempty"`
	CapabilityDefaultOriginate     *string `json:"capability-default-originate,omitempty"`
	CapabilityDefaultOriginate6    *string `json:"capability-default-originate6,omitempty"`
	CapabilityDynamic              *string `json:"capability-dynamic,omitempty"`
	CapabilityGracefulRestart      *string `json:"capability-graceful-restart,omitempty"`
	CapabilityGracefulRestartVpnv4 *string `json:"capability-graceful-restart-vpnv4,omitempty"`
	CapabilityGracefulRestart6     *string `json:"capability-graceful-restart6,omitempty"`
	CapabilityOrf                  *string `json:"capability-orf,omitempty"`
	CapabilityOrf6                 *string `json:"capability-orf6,omitempty"`
	CapabilityRouteRefresh         *string `json:"capability-route-refresh,omitempty"`
	ConnectTimer                   *int64  `json:"connect-timer,omitempty"`
	DefaultOriginateRoutemap       *string `json:"default-originate-routemap,omitempty"`
	DefaultOriginateRoutemap6      *string `json:"default-originate-routemap6,omitempty"`
	Description                    *string `json:"description,omitempty"`
	DistributeListIn               *string `json:"distribute-list-in,omitempty"`
	DistributeListInVpnv4          *string `json:"distribute-list-in-vpnv4,omitempty"`
	DistributeListIn6              *string `json:"distribute-list-in6,omitempty"`
	DistributeListOut              *string `json:"distribute-list-out,omitempty"`
	DistributeListOutVpnv4         *string `json:"distribute-list-out-vpnv4,omitempty"`
	DistributeListOut6             *string `json:"distribute-list-out6,omitempty"`
	DontCapabilityNegotiate        *string `json:"dont-capability-negotiate,omitempty"`
	EbgpEnforceMultihop            *string `json:"ebgp-enforce-multihop,omitempty"`
	EbgpMultihopTtl                *int64  `json:"ebgp-multihop-ttl,omitempty"`
	FilterListIn                   *string `json:"filter-list-in,omitempty"`
	FilterListIn6                  *string `json:"filter-list-in6,omitempty"`
	FilterListOut                  *string `json:"filter-list-out,omitempty"`
	FilterListOut6                 *string `json:"filter-list-out6,omitempty"`
	HoldtimeTimer                  *int64  `json:"holdtime-timer,omitempty"`
	Interface                      *string `json:"interface,omitempty"`
	KeepAliveTimer                 *int64  `json:"keep-alive-timer,omitempty"`
	LinkDownFailover               *string `json:"link-down-failover,omitempty"`
	LocalAs                        *int64  `json:"local-as,omitempty"`
	LocalAsNoPrepend               *string `json:"local-as-no-prepend,omitempty"`
	LocalAsReplaceAs               *string `json:"local-as-replace-as,omitempty"`
	MaximumPrefix                  *int64  `json:"maximum-prefix,omitempty"`
	MaximumPrefixThreshold         *int64  `json:"maximum-prefix-threshold,omitempty"`
	MaximumPrefixThresholdVpnv4    *int64  `json:"maximum-prefix-threshold-vpnv4,omitempty"`
	MaximumPrefixThreshold6        *int64  `json:"maximum-prefix-threshold6,omitempty"`
	MaximumPrefixVpnv4             *int64  `json:"maximum-prefix-vpnv4,omitempty"`
	MaximumPrefixWarningOnly       *string `json:"maximum-prefix-warning-only,omitempty"`
	MaximumPrefixWarningOnlyVpnv4  *string `json:"maximum-prefix-warning-only-vpnv4,omitempty"`
	MaximumPrefixWarningOnly6      *string `json:"maximum-prefix-warning-only6,omitempty"`
	MaximumPrefix6                 *int64  `json:"maximum-prefix6,omitempty"`
	Name                           *string `json:"name,omitempty"`
	NextHopSelf                    *string `json:"next-hop-self,omitempty"`
	NextHopSelfRr                  *string `json:"next-hop-self-rr,omitempty"`
	NextHopSelfRr6                 *string `json:"next-hop-self-rr6,omitempty"`
	NextHopSelfVpnv4               *string `json:"next-hop-self-vpnv4,omitempty"`
	NextHopSelf6                   *string `json:"next-hop-self6,omitempty"`
	OverrideCapability             *string `json:"override-capability,omitempty"`
	Passive                        *string `json:"passive,omitempty"`
	PrefixListIn                   *string `json:"prefix-list-in,omitempty"`
	PrefixListInVpnv4              *string `json:"prefix-list-in-vpnv4,omitempty"`
	PrefixListIn6                  *string `json:"prefix-list-in6,omitempty"`
	PrefixListOut                  *string `json:"prefix-list-out,omitempty"`
	PrefixListOutVpnv4             *string `json:"prefix-list-out-vpnv4,omitempty"`
	PrefixListOut6                 *string `json:"prefix-list-out6,omitempty"`
	RemoteAs                       *int64  `json:"remote-as,omitempty"`
	RemovePrivateAs                *string `json:"remove-private-as,omitempty"`
	RemovePrivateAsVpnv4           *string `json:"remove-private-as-vpnv4,omitempty"`
	RemovePrivateAs6               *string `json:"remove-private-as6,omitempty"`
	RestartTime                    *int64  `json:"restart-time,omitempty"`
	RetainStaleTime                *int64  `json:"retain-stale-time,omitempty"`
	RouteMapIn                     *string `json:"route-map-in,omitempty"`
	RouteMapInVpnv4                *string `json:"route-map-in-vpnv4,omitempty"`
	RouteMapIn6                    *string `json:"route-map-in6,omitempty"`
	RouteMapOut                    *string `json:"route-map-out,omitempty"`
	RouteMapOutPreferable          *string `json:"route-map-out-preferable,omitempty"`
	RouteMapOutVpnv4               *string `json:"route-map-out-vpnv4,omitempty"`
	RouteMapOutVpnv4Preferable     *string `json:"route-map-out-vpnv4-preferable,omitempty"`
	RouteMapOut6                   *string `json:"route-map-out6,omitempty"`
	RouteMapOut6Preferable         *string `json:"route-map-out6-preferable,omitempty"`
	RouteReflectorClient           *string `json:"route-reflector-client,omitempty"`
	RouteReflectorClientVpnv4      *string `json:"route-reflector-client-vpnv4,omitempty"`
	RouteReflectorClient6          *string `json:"route-reflector-client6,omitempty"`
	RouteServerClient              *string `json:"route-server-client,omitempty"`
	RouteServerClientVpnv4         *string `json:"route-server-client-vpnv4,omitempty"`
	RouteServerClient6             *string `json:"route-server-client6,omitempty"`
	SendCommunity                  *string `json:"send-community,omitempty"`
	SendCommunityVpnv4             *string `json:"send-community-vpnv4,omitempty"`
	SendCommunity6                 *string `json:"send-community6,omitempty"`
	Shutdown                       *string `json:"shutdown,omitempty"`
	SoftReconfiguration            *string `json:"soft-reconfiguration,omitempty"`
	SoftReconfigurationVpnv4       *string `json:"soft-reconfiguration-vpnv4,omitempty"`
	SoftReconfiguration6           *string `json:"soft-reconfiguration6,omitempty"`
	StaleRoute                     *string `json:"stale-route,omitempty"`
	StrictCapabilityMatch          *string `json:"strict-capability-match,omitempty"`
	UnsuppressMap                  *string `json:"unsuppress-map,omitempty"`
	UnsuppressMap6                 *string `json:"unsuppress-map6,omitempty"`
	UpdateSource                   *string `json:"update-source,omitempty"`
	Weight                         *int64  `json:"weight,omitempty"`
}

type RouterBgpNeighborRange struct {
	Id             *int64  `json:"id,omitempty"`
	MaxNeighborNum *int64  `json:"max-neighbor-num,omitempty"`
	NeighborGroup  *string `json:"neighbor-group,omitempty"`
	Prefix         *string `json:"prefix,omitempty"`
}

type RouterBgpNeighborRange6 struct {
	Id             *int64  `json:"id,omitempty"`
	MaxNeighborNum *int64  `json:"max-neighbor-num,omitempty"`
	NeighborGroup  *string `json:"neighbor-group,omitempty"`
	Prefix6        *string `json:"prefix6,omitempty"`
}

type RouterBgpNetwork struct {
	Backdoor           *string `json:"backdoor,omitempty"`
	Id                 *int64  `json:"id,omitempty"`
	NetworkImportCheck *string `json:"network-import-check,omitempty"`
	Prefix             *string `json:"prefix,omitempty"`
	RouteMap           *string `json:"route-map,omitempty"`
}

type RouterBgpNetwork6 struct {
	Backdoor           *string `json:"backdoor,omitempty"`
	Id                 *int64  `json:"id,omitempty"`
	NetworkImportCheck *string `json:"network-import-check,omitempty"`
	Prefix6            *string `json:"prefix6,omitempty"`
	RouteMap           *string `json:"route-map,omitempty"`
}

type RouterBgpRedistribute struct {
	Name     *string `json:"name,omitempty"`
	RouteMap *string `json:"route-map,omitempty"`
	Status   *string `json:"status,omitempty"`
}

//defaultfuncs
func (def *RouterBgpRedistribute) Defaults() {
	def.Name = ""
	def.RouteMap = ""
	def.Status = "disable"
}

type RouterBgpRedistribute6 struct {
	Name     *string `json:"name,omitempty"`
	RouteMap *string `json:"route-map,omitempty"`
	Status   *string `json:"status,omitempty"`
}

//defaultfuncs
func (def *RouterBgpRedistribute6) defaults() {
	def.Name = ""
	def.RouteMap = ""
	def.Status = "disable"
}

type RouterBgpVrf struct {
	ExportRt       *[]RouterBgpVrfExportRt   `json:"export-rt,omitempty"`
	ImportRouteMap *string                   `json:"import-route-map,omitempty"`
	ImportRt       *[]RouterBgpVrfImportRt   `json:"import-rt,omitempty"`
	LeakTarget     *[]RouterBgpVrfLeakTarget `json:"leak-target,omitempty"`
	Rd             *string                   `json:"rd,omitempty"`
	Role           *string                   `json:"role,omitempty"`
	Vrf            *string                   `json:"vrf,omitempty"`
}

type RouterBgpVrfExportRt struct {
	RouteTarget *string `json:"route-target,omitempty"`
}

type RouterBgpVrfImportRt struct {
	RouteTarget *string `json:"route-target,omitempty"`
}

type RouterBgpVrfLeakTarget struct {
	Interface *string `json:"interface,omitempty"`
	RouteMap  *string `json:"route-map,omitempty"`
	Vrf       *string `json:"vrf,omitempty"`
}

type RouterBgpVrfLeak struct {
	Target *[]RouterBgpVrfLeakTarget `json:"target,omitempty"`
	Vrf    *string                   `json:"vrf,omitempty"`
}

type RouterBgpVrfLeakTarget struct {
	Interface *string `json:"interface,omitempty"`
	RouteMap  *string `json:"route-map,omitempty"`
	Vrf       *string `json:"vrf,omitempty"`
}

type RouterBgpVrfLeak6 struct {
	Target *[]RouterBgpVrfLeak6Target `json:"target,omitempty"`
	Vrf    *string                    `json:"vrf,omitempty"`
}

type RouterBgpVrfLeak6Target struct {
	Interface *string `json:"interface,omitempty"`
	RouteMap  *string `json:"route-map,omitempty"`
	Vrf       *string `json:"vrf,omitempty"`
}

type RouterBgpVrf6 struct {
	LeakTarget *[]RouterBgpVrf6LeakTarget `json:"leak-target,omitempty"`
	Vrf        *string                    `json:"vrf,omitempty"`
}

type RouterBgpVrf6LeakTarget struct {
	Interface *string `json:"interface,omitempty"`
	RouteMap  *string `json:"route-map,omitempty"`
	Vrf       *string `json:"vrf,omitempty"`
}
