package models

const IpsRuleSettingsPath = "ips/rule-settings/"

type IpsRuleSettings struct {
	Id *int64 `json:"id,omitempty"`
}

//defaultfuncs
func (def *IpsRuleSettings) defaults() {
	def.Id = "0"
}
