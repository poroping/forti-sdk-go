package models

const SystemReplacemsgWebproxyPath = "system.replacemsg/webproxy/"

type SystemReplacemsgWebproxy struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

// Set SystemReplacemsgWebproxy values to defaults
func (def *SystemReplacemsgWebproxy) Defaults() {
	def.Buffer = stringPtr("<no value>")
	def.Format = stringPtr("")
	def.Header = stringPtr("")
	def.MsgType = stringPtr("")
}
