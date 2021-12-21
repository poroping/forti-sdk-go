package models

const ApplicationRuleSettingsPath = "application/rule-settings/"

type ApplicationRuleSettings struct {
	Fosid *float64 `json:"fosid,omitempty"`
}
