package models

const FirewallAddress6TemplatePath = "firewall/address6-template/"

type FirewallAddress6Template struct {
	FabricObject       *string                                 `json:"fabric-object,omitempty"`
	Ip6                *string                                 `json:"ip6,omitempty"`
	Name               *string                                 `json:"name,omitempty"`
	SubnetSegment      []FirewallAddress6TemplateSubnetSegment `json:"subnet-segment,omitempty"`
	SubnetSegmentCount *int64                                  `json:"subnet-segment-count,omitempty"`
}

type FirewallAddress6TemplateSubnetSegment struct {
	Bits      *int64                                        `json:"bits,omitempty"`
	Exclusive *string                                       `json:"exclusive,omitempty"`
	Id        *int64                                        `json:"id,omitempty"`
	Name      *string                                       `json:"name,omitempty"`
	Values    []FirewallAddress6TemplateSubnetSegmentValues `json:"values,omitempty"`
}

type FirewallAddress6TemplateSubnetSegmentValues struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}
