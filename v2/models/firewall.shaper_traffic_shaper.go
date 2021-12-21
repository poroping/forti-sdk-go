package models

const FirewallshaperTrafficShaperPath = "firewall.shaper/traffic-shaper/"

type FirewallshaperTrafficShaper struct {
	BandwidthUnit       *string  `json:"bandwidth-unit,omitempty"`
	Diffserv            *string  `json:"diffserv,omitempty"`
	Diffservcode        *string  `json:"diffservcode,omitempty"`
	DscpMarkingMethod   *string  `json:"dscp-marking-method,omitempty"`
	ExceedBandwidth     *float64 `json:"exceed-bandwidth,omitempty"`
	ExceedClassId       *float64 `json:"exceed-class-id,omitempty"`
	ExceedDscp          *string  `json:"exceed-dscp,omitempty"`
	GuaranteedBandwidth *float64 `json:"guaranteed-bandwidth,omitempty"`
	MaximumBandwidth    *float64 `json:"maximum-bandwidth,omitempty"`
	MaximumDscp         *string  `json:"maximum-dscp,omitempty"`
	Name                *string  `json:"name,omitempty"`
	Overhead            *float64 `json:"overhead,omitempty"`
	PerPolicy           *string  `json:"per-policy,omitempty"`
	Priority            *string  `json:"priority,omitempty"`
}
