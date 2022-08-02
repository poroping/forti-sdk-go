package models

const SystemReplacemsgAlertmailPath = "system.replacemsg/alertmail/"

type SystemReplacemsgAlertmail struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

// Set SystemReplacemsgAlertmail values to defaults
func (def *SystemReplacemsgAlertmail) Defaults() {
	def.Buffer = stringPtr("<no value>")
	def.Format = stringPtr("")
	def.Header = stringPtr("")
	def.MsgType = stringPtr("")
}
