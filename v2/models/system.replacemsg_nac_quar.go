package models

const SystemReplacemsgNacQuarPath = "system.replacemsg/nac-quar/"

type SystemReplacemsgNacQuar struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

// Set SystemReplacemsgNacQuar values to defaults
func (def *SystemReplacemsgNacQuar) Defaults() {
	def.Buffer = stringPtr("<no value>")
	def.Format = stringPtr("")
	def.Header = stringPtr("")
	def.MsgType = stringPtr("")
}
