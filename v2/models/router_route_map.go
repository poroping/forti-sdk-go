package models

const RouterRouteMapPath = "router/route-map/"

type RouterRouteMap struct {
	Comments *string               `json:"comments,omitempty"`
	Name     *string               `json:"name,omitempty"`
	Rule     *[]RouterRouteMapRule `json:"rule,omitempty"`
}

type RouterRouteMapRule struct {
	Action                             *string                                 `json:"action,omitempty"`
	Id                                 *int64                                  `json:"id,omitempty"`
	MatchAsPath                        *string                                 `json:"match-as-path,omitempty"`
	MatchCommunity                     *string                                 `json:"match-community,omitempty"`
	MatchCommunityExact                *string                                 `json:"match-community-exact,omitempty"`
	MatchFlags                         *int64                                  `json:"match-flags,omitempty"`
	MatchInterface                     *string                                 `json:"match-interface,omitempty"`
	MatchIpAddress                     *string                                 `json:"match-ip-address,omitempty"`
	MatchIpNexthop                     *string                                 `json:"match-ip-nexthop,omitempty"`
	MatchIp6Address                    *string                                 `json:"match-ip6-address,omitempty"`
	MatchIp6Nexthop                    *string                                 `json:"match-ip6-nexthop,omitempty"`
	MatchMetric                        *int64                                  `json:"match-metric,omitempty"`
	MatchOrigin                        *string                                 `json:"match-origin,omitempty"`
	MatchRouteType                     *string                                 `json:"match-route-type,omitempty"`
	MatchTag                           *int64                                  `json:"match-tag,omitempty"`
	MatchVrf                           *int64                                  `json:"match-vrf,omitempty"`
	SetAggregatorAs                    *int64                                  `json:"set-aggregator-as,omitempty"`
	SetAggregatorIp                    *string                                 `json:"set-aggregator-ip,omitempty"`
	SetAspath                          *[]RouterRouteMapRuleSetAspath          `json:"set-aspath,omitempty"`
	SetAspathAction                    *string                                 `json:"set-aspath-action,omitempty"`
	SetAtomicAggregate                 *string                                 `json:"set-atomic-aggregate,omitempty"`
	SetCommunity                       *[]RouterRouteMapRuleSetCommunity       `json:"set-community,omitempty"`
	SetCommunityAdditive               *string                                 `json:"set-community-additive,omitempty"`
	SetCommunityDelete                 *string                                 `json:"set-community-delete,omitempty"`
	SetDampeningMaxSuppress            *int64                                  `json:"set-dampening-max-suppress,omitempty"`
	SetDampeningReachabilityHalfLife   *int64                                  `json:"set-dampening-reachability-half-life,omitempty"`
	SetDampeningReuse                  *int64                                  `json:"set-dampening-reuse,omitempty"`
	SetDampeningSuppress               *int64                                  `json:"set-dampening-suppress,omitempty"`
	SetDampeningUnreachabilityHalfLife *int64                                  `json:"set-dampening-unreachability-half-life,omitempty"`
	SetExtcommunityRt                  *[]RouterRouteMapRuleSetExtcommunityRt  `json:"set-extcommunity-rt,omitempty"`
	SetExtcommunitySoo                 *[]RouterRouteMapRuleSetExtcommunitySoo `json:"set-extcommunity-soo,omitempty"`
	SetFlags                           *int64                                  `json:"set-flags,omitempty"`
	SetIpNexthop                       *string                                 `json:"set-ip-nexthop,omitempty"`
	SetIp6Nexthop                      *string                                 `json:"set-ip6-nexthop,omitempty"`
	SetIp6NexthopLocal                 *string                                 `json:"set-ip6-nexthop-local,omitempty"`
	SetLocalPreference                 *int64                                  `json:"set-local-preference,omitempty"`
	SetMetric                          *int64                                  `json:"set-metric,omitempty"`
	SetMetricType                      *string                                 `json:"set-metric-type,omitempty"`
	SetOrigin                          *string                                 `json:"set-origin,omitempty"`
	SetOriginatorId                    *string                                 `json:"set-originator-id,omitempty"`
	SetRouteTag                        *int64                                  `json:"set-route-tag,omitempty"`
	SetTag                             *int64                                  `json:"set-tag,omitempty"`
	SetWeight                          *int64                                  `json:"set-weight,omitempty"`
}

type RouterRouteMapRuleSetAspath struct {
	As *string `json:"as,omitempty"`
}

type RouterRouteMapRuleSetCommunity struct {
	Community *string `json:"community,omitempty"`
}

type RouterRouteMapRuleSetExtcommunityRt struct {
	Community *string `json:"community,omitempty"`
}

type RouterRouteMapRuleSetExtcommunitySoo struct {
	Community *string `json:"community,omitempty"`
}
