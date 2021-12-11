package models

const RouterOspfPath = "router/ospf/"

type RouterOspf struct {
	AbrType                       *string                      `json:"abr-type,omitempty"`
	Area                          []RouterOspfArea             `json:"area,omitempty"`
	AutoCostRefBandwidth          *int64                       `json:"auto-cost-ref-bandwidth,omitempty"`
	Bfd                           *string                      `json:"bfd,omitempty"`
	DatabaseOverflow              *string                      `json:"database-overflow,omitempty"`
	DatabaseOverflowMaxLsas       *int64                       `json:"database-overflow-max-lsas,omitempty"`
	DatabaseOverflowTimeToRecover *int64                       `json:"database-overflow-time-to-recover,omitempty"`
	DefaultInformationMetric      *int64                       `json:"default-information-metric,omitempty"`
	DefaultInformationMetricType  *string                      `json:"default-information-metric-type,omitempty"`
	DefaultInformationOriginate   *string                      `json:"default-information-originate,omitempty"`
	DefaultInformationRouteMap    *string                      `json:"default-information-route-map,omitempty"`
	DefaultMetric                 *int64                       `json:"default-metric,omitempty"`
	Distance                      *int64                       `json:"distance,omitempty"`
	DistanceExternal              *int64                       `json:"distance-external,omitempty"`
	DistanceInterArea             *int64                       `json:"distance-inter-area,omitempty"`
	DistanceIntraArea             *int64                       `json:"distance-intra-area,omitempty"`
	DistributeList                []RouterOspfDistributeList   `json:"distribute-list,omitempty"`
	DistributeListIn              *string                      `json:"distribute-list-in,omitempty"`
	DistributeRouteMapIn          *string                      `json:"distribute-route-map-in,omitempty"`
	LogNeighbourChanges           *string                      `json:"log-neighbour-changes,omitempty"`
	Neighbor                      []RouterOspfNeighbor         `json:"neighbor,omitempty"`
	Network                       []RouterOspfNetwork          `json:"network,omitempty"`
	OspfInterface                 []RouterOspfOspfInterface    `json:"ospf-interface,omitempty"`
	PassiveInterface              []RouterOspfPassiveInterface `json:"passive-interface,omitempty"`
	Redistribute                  []RouterOspfRedistribute     `json:"redistribute,omitempty"`
	RestartMode                   *string                      `json:"restart-mode,omitempty"`
	RestartPeriod                 *int64                       `json:"restart-period,omitempty"`
	Rfc1583Compatible             *string                      `json:"rfc1583-compatible,omitempty"`
	RouterId                      *string                      `json:"router-id,omitempty"`
	SpfTimers                     *string                      `json:"spf-timers,omitempty"`
	SummaryAddress                []RouterOspfSummaryAddress   `json:"summary-address,omitempty"`
}

type RouterOspfArea struct {
	Authentication                            *string                     `json:"authentication,omitempty"`
	Comments                                  *string                     `json:"comments,omitempty"`
	DefaultCost                               *int64                      `json:"default-cost,omitempty"`
	FilterList                                []RouterOspfAreaFilterList  `json:"filter-list,omitempty"`
	Id                                        *string                     `json:"id,omitempty"`
	NssaDefaultInformationOriginate           *string                     `json:"nssa-default-information-originate,omitempty"`
	NssaDefaultInformationOriginateMetric     *int64                      `json:"nssa-default-information-originate-metric,omitempty"`
	NssaDefaultInformationOriginateMetricType *string                     `json:"nssa-default-information-originate-metric-type,omitempty"`
	NssaRedistribution                        *string                     `json:"nssa-redistribution,omitempty"`
	NssaTranslatorRole                        *string                     `json:"nssa-translator-role,omitempty"`
	Range                                     []RouterOspfAreaRange       `json:"range,omitempty"`
	Shortcut                                  *string                     `json:"shortcut,omitempty"`
	StubType                                  *string                     `json:"stub-type,omitempty"`
	Type                                      *string                     `json:"type,omitempty"`
	VirtualLink                               []RouterOspfAreaVirtualLink `json:"virtual-link,omitempty"`
}

type RouterOspfAreaFilterList struct {
	Direction *string `json:"direction,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	List      *string `json:"list,omitempty"`
}

type RouterOspfAreaRange struct {
	Advertise        *string `json:"advertise,omitempty"`
	Id               *int64  `json:"id,omitempty"`
	Prefix           *string `json:"prefix,omitempty"`
	Substitute       *string `json:"substitute,omitempty"`
	SubstituteStatus *string `json:"substitute-status,omitempty"`
}

type RouterOspfAreaVirtualLink struct {
	Authentication     *string                            `json:"authentication,omitempty"`
	AuthenticationKey  *string                            `json:"authentication-key,omitempty"`
	DeadInterval       *int64                             `json:"dead-interval,omitempty"`
	HelloInterval      *int64                             `json:"hello-interval,omitempty"`
	Keychain           *string                            `json:"keychain,omitempty"`
	Md5Keys            []RouterOspfAreaVirtualLinkMd5Keys `json:"md5-keys,omitempty"`
	Name               *string                            `json:"name,omitempty"`
	Peer               *string                            `json:"peer,omitempty"`
	RetransmitInterval *int64                             `json:"retransmit-interval,omitempty"`
	TransmitDelay      *int64                             `json:"transmit-delay,omitempty"`
}

type RouterOspfAreaVirtualLinkMd5Keys struct {
	Id        *int64  `json:"id,omitempty"`
	KeyString *string `json:"key-string,omitempty"`
}

type RouterOspfDistributeList struct {
	AccessList *string `json:"access-list,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	Protocol   *string `json:"protocol,omitempty"`
}

type RouterOspfNeighbor struct {
	Cost         *int64  `json:"cost,omitempty"`
	Id           *int64  `json:"id,omitempty"`
	Ip           *string `json:"ip,omitempty"`
	PollInterval *int64  `json:"poll-interval,omitempty"`
	Priority     *int64  `json:"priority,omitempty"`
}

type RouterOspfNetwork struct {
	Area     *string `json:"area,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Prefix   *string `json:"prefix,omitempty"`
}

type RouterOspfOspfInterface struct {
	Authentication     *string                          `json:"authentication,omitempty"`
	AuthenticationKey  *string                          `json:"authentication-key,omitempty"`
	Bfd                *string                          `json:"bfd,omitempty"`
	Comments           *string                          `json:"comments,omitempty"`
	Cost               *int64                           `json:"cost,omitempty"`
	DatabaseFilterOut  *string                          `json:"database-filter-out,omitempty"`
	DeadInterval       *int64                           `json:"dead-interval,omitempty"`
	HelloInterval      *int64                           `json:"hello-interval,omitempty"`
	HelloMultiplier    *int64                           `json:"hello-multiplier,omitempty"`
	Interface          *string                          `json:"interface,omitempty"`
	Ip                 *string                          `json:"ip,omitempty"`
	Keychain           *string                          `json:"keychain,omitempty"`
	Md5Keys            []RouterOspfOspfInterfaceMd5Keys `json:"md5-keys,omitempty"`
	Mtu                *int64                           `json:"mtu,omitempty"`
	MtuIgnore          *string                          `json:"mtu-ignore,omitempty"`
	Name               *string                          `json:"name,omitempty"`
	NetworkType        *string                          `json:"network-type,omitempty"`
	PrefixLength       *int64                           `json:"prefix-length,omitempty"`
	Priority           *int64                           `json:"priority,omitempty"`
	ResyncTimeout      *int64                           `json:"resync-timeout,omitempty"`
	RetransmitInterval *int64                           `json:"retransmit-interval,omitempty"`
	Status             *string                          `json:"status,omitempty"`
	TransmitDelay      *int64                           `json:"transmit-delay,omitempty"`
}

type RouterOspfOspfInterfaceMd5Keys struct {
	Id        *int64  `json:"id,omitempty"`
	KeyString *string `json:"key-string,omitempty"`
}

type RouterOspfPassiveInterface struct {
	Name *string `json:"name,omitempty"`
}

type RouterOspfRedistribute struct {
	Metric     *int64  `json:"metric,omitempty"`
	MetricType *string `json:"metric-type,omitempty"`
	Name       *string `json:"name,omitempty"`
	Routemap   *string `json:"routemap,omitempty"`
	Status     *string `json:"status,omitempty"`
	Tag        *int64  `json:"tag,omitempty"`
}

type RouterOspfSummaryAddress struct {
	Advertise *string `json:"advertise,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	Prefix    *string `json:"prefix,omitempty"`
	Tag       *int64  `json:"tag,omitempty"`
}
