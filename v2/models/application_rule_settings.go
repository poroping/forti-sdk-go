package models

const ApplicationRuleSettingsPath = "application/rule-settings/"

type ApplicationRuleSettings struct {
	Id *int64 `json:"id,omitempty"`
}

// Set ApplicationRuleSettings values to defaults
func (def *ApplicationRuleSettings) Defaults() {
	def.Id = intPtr(0)
}
