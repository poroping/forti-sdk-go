package models

const SystemReplacemsgMailPath = "system.replacemsg/mail/"

type SystemReplacemsgMail struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

// Set SystemReplacemsgMail values to defaults
func (def *SystemReplacemsgMail) Defaults() {
	def.Buffer = stringPtr("<no value>")
	def.Format = stringPtr("")
	def.Header = stringPtr("")
	def.MsgType = stringPtr("")
}
