package models

const FirewallInternetServicePath = "firewall/internet-service/"

type FirewallInternetService struct {
	Database           *string `json:"database,omitempty"`
	Direction          *string `json:"direction,omitempty"`
	ExtraIpRangeNumber *int64  `json:"extra-ip-range-number,omitempty"`
	IconId             *int64  `json:"icon-id,omitempty"`
	Fosid              *int64  `json:"fosid,omitempty"`
	IpNumber           *int64  `json:"ip-number,omitempty"`
	IpRangeNumber      *int64  `json:"ip-range-number,omitempty"`
	Name               *string `json:"name,omitempty"`
	Obsolete           *int64  `json:"obsolete,omitempty"`
	Singularity        *int64  `json:"singularity,omitempty"`
}
