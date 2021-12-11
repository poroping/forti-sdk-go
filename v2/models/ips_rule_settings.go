package models

const IpsRuleSettingsPath = "ips/rule-settings/"

type IpsRuleSettings struct {
	Fosid *int64 `json:"fosid,omitempty"`
}
