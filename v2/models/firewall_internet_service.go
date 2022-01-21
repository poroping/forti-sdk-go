package models

const FirewallInternetServicePath = "firewall/internet-service/"

type FirewallInternetService struct {
	Database           *string `json:"database,omitempty"`
	Direction          *string `json:"direction,omitempty"`
	ExtraIpRangeNumber *int64  `json:"extra-ip-range-number,omitempty"`
	IconId             *int64  `json:"icon-id,omitempty"`
	Id                 *int64  `json:"id,omitempty"`
	IpNumber           *int64  `json:"ip-number,omitempty"`
	IpRangeNumber      *int64  `json:"ip-range-number,omitempty"`
	Name               *string `json:"name,omitempty"`
	Obsolete           *int64  `json:"obsolete,omitempty"`
	Reputation         *int64  `json:"reputation,omitempty"`
	Singularity        *int64  `json:"singularity,omitempty"`
	SldId              *int64  `json:"sld-id,omitempty"`
}
