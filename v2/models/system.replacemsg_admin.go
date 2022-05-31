package models

const SystemReplacemsgAdminPath = "system.replacemsg/admin/"

type SystemReplacemsgAdmin struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

//defaultfuncs
func (def *SystemReplacemsgAdmin) defaults() {
	def.Buffer = "<no value>"
	def.Format = ""
	def.Header = ""
	def.MsgType = ""
}
