package models

const ApplicationRuleSettingsPath = "application/rule-settings/"

type ApplicationRuleSettings struct {
	Fosid *int64 `json:"fosid,omitempty"`
}
