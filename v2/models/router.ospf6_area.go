package models

const Routerospf6AreaPath = "router.ospf6/area/"

type Routerospf6Area struct {
	Authentication                            *string                       `json:"authentication,omitempty"`
	DefaultCost                               *float64                      `json:"default-cost,omitempty"`
	Fosid                                     *string                       `json:"fosid,omitempty"`
	IpsecAuthAlg                              *string                       `json:"ipsec-auth-alg,omitempty"`
	IpsecEncAlg                               *string                       `json:"ipsec-enc-alg,omitempty"`
	IpsecKeys                                 *[]Routerospf6AreaIpsecKeys   `json:"ipsec-keys,omitempty"`
	KeyRolloverInterval                       *float64                      `json:"key-rollover-interval,omitempty"`
	NssaDefaultInformationOriginate           *string                       `json:"nssa-default-information-originate,omitempty"`
	NssaDefaultInformationOriginateMetric     *float64                      `json:"nssa-default-information-originate-metric,omitempty"`
	NssaDefaultInformationOriginateMetricType *string                       `json:"nssa-default-information-originate-metric-type,omitempty"`
	NssaRedistribution                        *string                       `json:"nssa-redistribution,omitempty"`
	NssaTranslatorRole                        *string                       `json:"nssa-translator-role,omitempty"`
	Range                                     *[]Routerospf6AreaRange       `json:"range,omitempty"`
	StubType                                  *string                       `json:"stub-type,omitempty"`
	Type                                      *string                       `json:"type,omitempty"`
	VirtualLink                               *[]Routerospf6AreaVirtualLink `json:"virtual-link,omitempty"`
}

type Routerospf6AreaIpsecKeys struct {
	AuthKey *string  `json:"auth-key,omitempty"`
	EncKey  *string  `json:"enc-key,omitempty"`
	Spi     *float64 `json:"spi,omitempty"`
}

type Routerospf6AreaRange struct {
	Advertise *string  `json:"advertise,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	Prefix6   *string  `json:"prefix6,omitempty"`
}

type Routerospf6AreaVirtualLink struct {
	Authentication      *string                                `json:"authentication,omitempty"`
	DeadInterval        *float64                               `json:"dead-interval,omitempty"`
	HelloInterval       *float64                               `json:"hello-interval,omitempty"`
	IpsecAuthAlg        *string                                `json:"ipsec-auth-alg,omitempty"`
	IpsecEncAlg         *string                                `json:"ipsec-enc-alg,omitempty"`
	IpsecKeys           *[]Routerospf6AreaVirtualLinkIpsecKeys `json:"ipsec-keys,omitempty"`
	KeyRolloverInterval *float64                               `json:"key-rollover-interval,omitempty"`
	Name                *string                                `json:"name,omitempty"`
	Peer                *string                                `json:"peer,omitempty"`
	RetransmitInterval  *float64                               `json:"retransmit-interval,omitempty"`
	TransmitDelay       *float64                               `json:"transmit-delay,omitempty"`
}

type Routerospf6AreaVirtualLinkIpsecKeys struct {
	AuthKey *string  `json:"auth-key,omitempty"`
	EncKey  *string  `json:"enc-key,omitempty"`
	Spi     *float64 `json:"spi,omitempty"`
}
