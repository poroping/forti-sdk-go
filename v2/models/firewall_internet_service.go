package models

const FirewallInternetServicePath = "firewall/internet-service/"

type FirewallInternetService struct {
	Database           *string  `json:"database,omitempty"`
	Direction          *string  `json:"direction,omitempty"`
	ExtraIpRangeNumber *float64 `json:"extra-ip-range-number,omitempty"`
	IconId             *float64 `json:"icon-id,omitempty"`
	Fosid              *float64 `json:"fosid,omitempty"`
	IpNumber           *float64 `json:"ip-number,omitempty"`
	IpRangeNumber      *float64 `json:"ip-range-number,omitempty"`
	Name               *string  `json:"name,omitempty"`
	Obsolete           *float64 `json:"obsolete,omitempty"`
	Reputation         *float64 `json:"reputation,omitempty"`
	Singularity        *float64 `json:"singularity,omitempty"`
	SldId              *float64 `json:"sld-id,omitempty"`
}
