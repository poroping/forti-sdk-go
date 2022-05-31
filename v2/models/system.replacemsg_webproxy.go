package models

const SystemReplacemsgWebproxyPath = "system.replacemsg/webproxy/"

type SystemReplacemsgWebproxy struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

//defaultfuncs
func (def *SystemReplacemsgWebproxy) defaults() {
	def.Buffer = "<no value>"
	def.Format = ""
	def.Header = ""
	def.MsgType = ""
}
