package models

const IpsRuleSettingsPath = "ips/rule-settings/"

type IpsRuleSettings struct {
	Fosid *float64 `json:"fosid,omitempty"`
}
