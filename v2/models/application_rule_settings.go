package models

const ApplicationRuleSettingsPath = "application/rule-settings/"

type ApplicationRuleSettings struct {
	Id *int64 `json:"id,omitempty"`
}

//defaultfuncs
func (def *ApplicationRuleSettings) defaults() {
	def.Id = "0"
}
