package models

const SystemReplacemsgFortiguardWfPath = "system.replacemsg/fortiguard-wf/"

type SystemReplacemsgFortiguardWf struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

// Set SystemReplacemsgFortiguardWf values to defaults
func (def *SystemReplacemsgFortiguardWf) Defaults() {
	def.Buffer = stringPtr("<no value>")
	def.Format = stringPtr("")
	def.Header = stringPtr("")
	def.MsgType = stringPtr("")
}
