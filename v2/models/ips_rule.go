package models

const IpsRulePath = "ips/rule/"

type IpsRule struct {
	Action      *string            `json:"action,omitempty"`
	Application *string            `json:"application,omitempty"`
	Date        *float64           `json:"date,omitempty"`
	Group       *string            `json:"group,omitempty"`
	Location    *string            `json:"location,omitempty"`
	Log         *string            `json:"log,omitempty"`
	LogPacket   *string            `json:"log-packet,omitempty"`
	Metadata    *[]IpsRuleMetadata `json:"metadata,omitempty"`
	Name        *string            `json:"name,omitempty"`
	Os          *string            `json:"os,omitempty"`
	Rev         *float64           `json:"rev,omitempty"`
	RuleId      *float64           `json:"rule-id,omitempty"`
	Service     *string            `json:"service,omitempty"`
	Severity    *string            `json:"severity,omitempty"`
	Status      *string            `json:"status,omitempty"`
}

type IpsRuleMetadata struct {
	Id      *float64 `json:"id,omitempty"`
	Metaid  *float64 `json:"metaid,omitempty"`
	Valueid *float64 `json:"valueid,omitempty"`
}
