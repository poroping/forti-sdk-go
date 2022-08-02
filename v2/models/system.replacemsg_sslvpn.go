package models

const SystemReplacemsgSslvpnPath = "system.replacemsg/sslvpn/"

type SystemReplacemsgSslvpn struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

// Set SystemReplacemsgSslvpn values to defaults
func (def *SystemReplacemsgSslvpn) Defaults() {
	def.Buffer = stringPtr("<no value>")
	def.Format = stringPtr("")
	def.Header = stringPtr("")
	def.MsgType = stringPtr("")
}
