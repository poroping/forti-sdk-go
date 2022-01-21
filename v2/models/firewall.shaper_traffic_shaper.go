package models

const FirewallShaperTrafficShaperPath = "firewall.shaper/traffic-shaper/"

type FirewallShaperTrafficShaper struct {
	BandwidthUnit       *string `json:"bandwidth-unit,omitempty"`
	Diffserv            *string `json:"diffserv,omitempty"`
	Diffservcode        *string `json:"diffservcode,omitempty"`
	DscpMarkingMethod   *string `json:"dscp-marking-method,omitempty"`
	ExceedBandwidth     *int64  `json:"exceed-bandwidth,omitempty"`
	ExceedClassId       *int64  `json:"exceed-class-id,omitempty"`
	ExceedDscp          *string `json:"exceed-dscp,omitempty"`
	GuaranteedBandwidth *int64  `json:"guaranteed-bandwidth,omitempty"`
	MaximumBandwidth    *int64  `json:"maximum-bandwidth,omitempty"`
	MaximumDscp         *string `json:"maximum-dscp,omitempty"`
	Name                *string `json:"name,omitempty"`
	Overhead            *int64  `json:"overhead,omitempty"`
	PerPolicy           *string `json:"per-policy,omitempty"`
	Priority            *string `json:"priority,omitempty"`
}
