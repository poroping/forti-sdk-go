package models

const SystemReplacemsgIcapPath = "system.replacemsg/icap/"

type SystemReplacemsgIcap struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

//defaultfuncs
func (def *SystemReplacemsgIcap) defaults() {
	def.Buffer = "<no value>"
	def.Format = ""
	def.Header = ""
	def.MsgType = ""
}
