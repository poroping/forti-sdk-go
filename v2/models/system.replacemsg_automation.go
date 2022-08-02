package models

const SystemReplacemsgAutomationPath = "system.replacemsg/automation/"

type SystemReplacemsgAutomation struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

// Set SystemReplacemsgAutomation values to defaults
func (def *SystemReplacemsgAutomation) Defaults() {
	def.Buffer = stringPtr("<no value>")
	def.Format = stringPtr("")
	def.Header = stringPtr("")
	def.MsgType = stringPtr("")
}
