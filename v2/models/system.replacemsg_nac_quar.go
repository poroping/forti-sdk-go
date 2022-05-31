package models

const SystemReplacemsgNacQuarPath = "system.replacemsg/nac-quar/"

type SystemReplacemsgNacQuar struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

//defaultfuncs
func (def *SystemReplacemsgNacQuar) defaults() {
	def.Buffer = "<no value>"
	def.Format = ""
	def.Header = ""
	def.MsgType = ""
}
