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

//defaultfuncs
func (def *FirewallInternetService) defaults() {
	def.Database = "isdb"
	def.Direction = "both"
	def.ExtraIpRangeNumber = "0"
	def.IconId = "0"
	def.Id = "0"
	def.IpNumber = "0"
	def.IpRangeNumber = "0"
	def.Name = ""
	def.Obsolete = "0"
	def.Reputation = "<no value>"
	def.Singularity = "0"
	def.SldId = "<no value>"
}
