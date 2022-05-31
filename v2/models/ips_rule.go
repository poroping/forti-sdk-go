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

type IpsRuleMetadata struct {
	Id      *int64 `json:"id,omitempty"`
	Metaid  *int64 `json:"metaid,omitempty"`
	Valueid *int64 `json:"valueid,omitempty"`
}

//defaultfuncs
func (def *IpsRuleMetadata) defaults() {
	def.Id = "0"
	def.Metaid = "0"
	def.Valueid = "0"
}

//defaultfuncs
func (def *IpsRule) defaults() {
	def.Action = "pass"
	def.Application = ""
	def.Date = "0"
	def.Group = ""
	def.Location = ""
	def.Log = "enable"
	def.LogPacket = "disable"
	def.Name = ""
	def.Os = ""
	def.Rev = "0"
	def.RuleId = "0"
	def.Service = ""
	def.Severity = ""
	def.Status = "enable"
}
