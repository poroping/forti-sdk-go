package models

const SystemReplacemsgNntpPath = "system.replacemsg/nntp/"

type SystemReplacemsgNntp struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

//defaultfuncs
func (def *SystemReplacemsgNntp) defaults() {
	def.Buffer = "<no value>"
	def.Format = "<no value>"
	def.Header = "<no value>"
	def.MsgType = "<no value>"
}
