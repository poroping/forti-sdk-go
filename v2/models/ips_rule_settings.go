package models

const IpsRuleSettingsPath = "ips/rule-settings/"

type IpsRuleSettings struct {
	Id *int64 `json:"id,omitempty"`
}

// Set IpsRuleSettings values to defaults
func (def *IpsRuleSettings) Defaults() {
	def.Id = intPtr(0)
}
