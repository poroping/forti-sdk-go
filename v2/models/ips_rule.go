package models

const IpsRulePath = "ips/rule/"

type IpsRule struct {
	Action      *string            `json:"action,omitempty"`
	Application *string            `json:"application,omitempty"`
	Date        *int64             `json:"date,omitempty"`
	Group       *string            `json:"group,omitempty"`
	Location    *string            `json:"location,omitempty"`
	Log         *string            `json:"log,omitempty"`
	LogPacket   *string            `json:"log-packet,omitempty"`
	Metadata    *[]IpsRuleMetadata `json:"metadata,omitempty"`
	Name        *string            `json:"name,omitempty"`
	Os          *string            `json:"os,omitempty"`
	Rev         *int64             `json:"rev,omitempty"`
	RuleId      *int64             `json:"rule-id,omitempty"`
	Service     *string            `json:"service,omitempty"`
	Severity    *string            `json:"severity,omitempty"`
	Status      *string            `json:"status,omitempty"`
}

const IpsRuleMetadataPath = "ips/rule/metadata/"

type IpsRuleMetadata struct {
	Id      *int64 `json:"id,omitempty"`
	Metaid  *int64 `json:"metaid,omitempty"`
	Valueid *int64 `json:"valueid,omitempty"`
}

// Set IpsRuleMetadata values to defaults
func (def *IpsRuleMetadata) Defaults() {
	def.Id = intPtr(0)
	def.Metaid = intPtr(0)
	def.Valueid = intPtr(0)
}

// Set IpsRule values to defaults
func (def *IpsRule) Defaults() {
	def.Action = stringPtr("pass")
	def.Application = stringPtr("")
	def.Date = intPtr(0)
	def.Group = stringPtr("")
	def.Location = stringPtr("")
	def.Log = stringPtr("enable")
	def.LogPacket = stringPtr("disable")
	def.Name = stringPtr("")
	def.Os = stringPtr("")
	def.Rev = intPtr(0)
	def.RuleId = intPtr(0)
	def.Service = stringPtr("")
	def.Severity = stringPtr("")
	def.Status = stringPtr("enable")
}
