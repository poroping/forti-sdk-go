package models

const RouterOspf6Path = "router/ospf6/"

type RouterOspf6 struct {
	AbrType                      *string                        `json:"abr-type,omitempty"`
	Area                         *[]RouterOspf6Area             `json:"area,omitempty"`
	AutoCostRefBandwidth         *float64                       `json:"auto-cost-ref-bandwidth,omitempty"`
	Bfd                          *string                        `json:"bfd,omitempty"`
	DefaultInformationMetric     *float64                       `json:"default-information-metric,omitempty"`
	DefaultInformationMetricType *string                        `json:"default-information-metric-type,omitempty"`
	DefaultInformationOriginate  *string                        `json:"default-information-originate,omitempty"`
	DefaultInformationRouteMap   *string                        `json:"default-information-route-map,omitempty"`
	DefaultMetric                *float64                       `json:"default-metric,omitempty"`
	LogNeighbourChanges          *string                        `json:"log-neighbour-changes,omitempty"`
	Ospf6Interface               *[]RouterOspf6Ospf6Interface   `json:"ospf6-interface,omitempty"`
	PassiveInterface             *[]RouterOspf6PassiveInterface `json:"passive-interface,omitempty"`
	Redistribute                 *[]RouterOspf6Redistribute     `json:"redistribute,omitempty"`
	RouterId                     *string                        `json:"router-id,omitempty"`
	SpfTimers                    *string                        `json:"spf-timers,omitempty"`
	SummaryAddress               *[]RouterOspf6SummaryAddress   `json:"summary-address,omitempty"`
}

type RouterOspf6Area struct {
	Authentication                            *string                       `json:"authentication,omitempty"`
	DefaultCost                               *float64                      `json:"default-cost,omitempty"`
	Id                                        *string                       `json:"id,omitempty"`
	IpsecAuthAlg                              *string                       `json:"ipsec-auth-alg,omitempty"`
	IpsecEncAlg                               *string                       `json:"ipsec-enc-alg,omitempty"`
	IpsecKeys                                 *[]RouterOspf6AreaIpsecKeys   `json:"ipsec-keys,omitempty"`
	KeyRolloverInterval                       *float64                      `json:"key-rollover-interval,omitempty"`
	NssaDefaultInformationOriginate           *string                       `json:"nssa-default-information-originate,omitempty"`
	NssaDefaultInformationOriginateMetric     *float64                      `json:"nssa-default-information-originate-metric,omitempty"`
	NssaDefaultInformationOriginateMetricType *string                       `json:"nssa-default-information-originate-metric-type,omitempty"`
	NssaRedistribution                        *string                       `json:"nssa-redistribution,omitempty"`
	NssaTranslatorRole                        *string                       `json:"nssa-translator-role,omitempty"`
	Range                                     *[]RouterOspf6AreaRange       `json:"range,omitempty"`
	StubType                                  *string                       `json:"stub-type,omitempty"`
	Type                                      *string                       `json:"type,omitempty"`
	VirtualLink                               *[]RouterOspf6AreaVirtualLink `json:"virtual-link,omitempty"`
}

type RouterOspf6AreaIpsecKeys struct {
	AuthKey *string  `json:"auth-key,omitempty"`
	EncKey  *string  `json:"enc-key,omitempty"`
	Spi     *float64 `json:"spi,omitempty"`
}

type RouterOspf6AreaRange struct {
	Advertise *string  `json:"advertise,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	Prefix6   *string  `json:"prefix6,omitempty"`
}

type RouterOspf6AreaVirtualLink struct {
	Authentication      *string                                `json:"authentication,omitempty"`
	DeadInterval        *float64                               `json:"dead-interval,omitempty"`
	HelloInterval       *float64                               `json:"hello-interval,omitempty"`
	IpsecAuthAlg        *string                                `json:"ipsec-auth-alg,omitempty"`
	IpsecEncAlg         *string                                `json:"ipsec-enc-alg,omitempty"`
	IpsecKeys           *[]RouterOspf6AreaVirtualLinkIpsecKeys `json:"ipsec-keys,omitempty"`
	KeyRolloverInterval *float64                               `json:"key-rollover-interval,omitempty"`
	Name                *string                                `json:"name,omitempty"`
	Peer                *string                                `json:"peer,omitempty"`
	RetransmitInterval  *float64                               `json:"retransmit-interval,omitempty"`
	TransmitDelay       *float64                               `json:"transmit-delay,omitempty"`
}

type RouterOspf6AreaVirtualLinkIpsecKeys struct {
	AuthKey *string  `json:"auth-key,omitempty"`
	EncKey  *string  `json:"enc-key,omitempty"`
	Spi     *float64 `json:"spi,omitempty"`
}

type RouterOspf6Ospf6Interface struct {
	AreaId              *string                               `json:"area-id,omitempty"`
	Authentication      *string                               `json:"authentication,omitempty"`
	Bfd                 *string                               `json:"bfd,omitempty"`
	Cost                *float64                              `json:"cost,omitempty"`
	DeadInterval        *float64                              `json:"dead-interval,omitempty"`
	HelloInterval       *float64                              `json:"hello-interval,omitempty"`
	Interface           *string                               `json:"interface,omitempty"`
	IpsecAuthAlg        *string                               `json:"ipsec-auth-alg,omitempty"`
	IpsecEncAlg         *string                               `json:"ipsec-enc-alg,omitempty"`
	IpsecKeys           *[]RouterOspf6Ospf6InterfaceIpsecKeys `json:"ipsec-keys,omitempty"`
	KeyRolloverInterval *float64                              `json:"key-rollover-interval,omitempty"`
	Mtu                 *float64                              `json:"mtu,omitempty"`
	MtuIgnore           *string                               `json:"mtu-ignore,omitempty"`
	Name                *string                               `json:"name,omitempty"`
	Neighbor            *[]RouterOspf6Ospf6InterfaceNeighbor  `json:"neighbor,omitempty"`
	NetworkType         *string                               `json:"network-type,omitempty"`
	Priority            *float64                              `json:"priority,omitempty"`
	RetransmitInterval  *float64                              `json:"retransmit-interval,omitempty"`
	Status              *string                               `json:"status,omitempty"`
	TransmitDelay       *float64                              `json:"transmit-delay,omitempty"`
}

type RouterOspf6Ospf6InterfaceIpsecKeys struct {
	AuthKey *string  `json:"auth-key,omitempty"`
	EncKey  *string  `json:"enc-key,omitempty"`
	Spi     *float64 `json:"spi,omitempty"`
}

type RouterOspf6Ospf6InterfaceNeighbor struct {
	Cost         *float64 `json:"cost,omitempty"`
	Ip6          *string  `json:"ip6,omitempty"`
	PollInterval *float64 `json:"poll-interval,omitempty"`
	Priority     *float64 `json:"priority,omitempty"`
}

type RouterOspf6PassiveInterface struct {
	Name *string `json:"name,omitempty"`
}

type RouterOspf6Redistribute struct {
	Metric     *float64 `json:"metric,omitempty"`
	MetricType *string  `json:"metric-type,omitempty"`
	Name       *string  `json:"name,omitempty"`
	Routemap   *string  `json:"routemap,omitempty"`
	Status     *string  `json:"status,omitempty"`
}

type RouterOspf6SummaryAddress struct {
	Advertise *string  `json:"advertise,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	Prefix6   *string  `json:"prefix6,omitempty"`
	Tag       *float64 `json:"tag,omitempty"`
}
