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
	Singularity        *int64  `json:"singularity,omitempty"`
}

// Set FirewallInternetService values to defaults
func (def *FirewallInternetService) Defaults() {
	def.Database = stringPtr("isdb")
	def.Direction = stringPtr("both")
	def.ExtraIpRangeNumber = intPtr(0)
	def.IconId = intPtr(0)
	def.Id = intPtr(0)
	def.IpNumber = intPtr(0)
	def.IpRangeNumber = intPtr(0)
	def.Name = stringPtr("")
	def.Obsolete = intPtr(0)
	def.Singularity = intPtr(0)
}
