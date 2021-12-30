package models

const WirelessControllerAccessControlListPath = "wireless-controller/access-control-list/"

type WirelessControllerAccessControlList struct {
	Comment         *string                                               `json:"comment,omitempty"`
	Layer3Ipv4Rules *[]WirelessControllerAccessControlListLayer3Ipv4Rules `json:"layer3-ipv4-rules,omitempty"`
	Layer3Ipv6Rules *[]WirelessControllerAccessControlListLayer3Ipv6Rules `json:"layer3-ipv6-rules,omitempty"`
	Name            *string                                               `json:"name,omitempty"`
}

type WirelessControllerAccessControlListLayer3Ipv4Rules struct {
	Action   *string `json:"action,omitempty"`
	Comment  *string `json:"comment,omitempty"`
	Dstaddr  *string `json:"dstaddr,omitempty"`
	Dstport  *int64  `json:"dstport,omitempty"`
	Protocol *int64  `json:"protocol,omitempty"`
	RuleId   *int64  `json:"rule-id,omitempty"`
	Srcaddr  *string `json:"srcaddr,omitempty"`
	Srcport  *int64  `json:"srcport,omitempty"`
}

type WirelessControllerAccessControlListLayer3Ipv6Rules struct {
	Action   *string `json:"action,omitempty"`
	Comment  *string `json:"comment,omitempty"`
	Dstaddr  *string `json:"dstaddr,omitempty"`
	Dstport  *int64  `json:"dstport,omitempty"`
	Protocol *int64  `json:"protocol,omitempty"`
	RuleId   *int64  `json:"rule-id,omitempty"`
	Srcaddr  *string `json:"srcaddr,omitempty"`
	Srcport  *int64  `json:"srcport,omitempty"`
}
