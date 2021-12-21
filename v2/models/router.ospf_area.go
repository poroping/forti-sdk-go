package models

const RouterospfAreaPath = "router.ospf/area/"

type RouterospfArea struct {
	Authentication                            *string                      `json:"authentication,omitempty"`
	Comments                                  *string                      `json:"comments,omitempty"`
	DefaultCost                               *int64                       `json:"default-cost,omitempty"`
	FilterList                                *[]RouterospfAreaFilterList  `json:"filter-list,omitempty"`
	Fosid                                     *string                      `json:"fosid,omitempty"`
	NssaDefaultInformationOriginate           *string                      `json:"nssa-default-information-originate,omitempty"`
	NssaDefaultInformationOriginateMetric     *int64                       `json:"nssa-default-information-originate-metric,omitempty"`
	NssaDefaultInformationOriginateMetricType *string                      `json:"nssa-default-information-originate-metric-type,omitempty"`
	NssaRedistribution                        *string                      `json:"nssa-redistribution,omitempty"`
	NssaTranslatorRole                        *string                      `json:"nssa-translator-role,omitempty"`
	Range                                     *[]RouterospfAreaRange       `json:"range,omitempty"`
	Shortcut                                  *string                      `json:"shortcut,omitempty"`
	StubType                                  *string                      `json:"stub-type,omitempty"`
	Type                                      *string                      `json:"type,omitempty"`
	VirtualLink                               *[]RouterospfAreaVirtualLink `json:"virtual-link,omitempty"`
}

type RouterospfAreaFilterList struct {
	Direction *string `json:"direction,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	List      *string `json:"list,omitempty"`
}

type RouterospfAreaRange struct {
	Advertise        *string `json:"advertise,omitempty"`
	Id               *int64  `json:"id,omitempty"`
	Prefix           *string `json:"prefix,omitempty"`
	Substitute       *string `json:"substitute,omitempty"`
	SubstituteStatus *string `json:"substitute-status,omitempty"`
}

type RouterospfAreaVirtualLink struct {
	Authentication     *string                             `json:"authentication,omitempty"`
	AuthenticationKey  *string                             `json:"authentication-key,omitempty"`
	DeadInterval       *int64                              `json:"dead-interval,omitempty"`
	HelloInterval      *int64                              `json:"hello-interval,omitempty"`
	Keychain           *string                             `json:"keychain,omitempty"`
	Md5Keychain        *string                             `json:"md5-keychain,omitempty"`
	Md5Keys            *[]RouterospfAreaVirtualLinkMd5Keys `json:"md5-keys,omitempty"`
	Name               *string                             `json:"name,omitempty"`
	Peer               *string                             `json:"peer,omitempty"`
	RetransmitInterval *int64                              `json:"retransmit-interval,omitempty"`
	TransmitDelay      *int64                              `json:"transmit-delay,omitempty"`
}

type RouterospfAreaVirtualLinkMd5Keys struct {
	Id        *int64  `json:"id,omitempty"`
	KeyString *string `json:"key-string,omitempty"`
}
