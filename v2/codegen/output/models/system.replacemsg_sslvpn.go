package models

const SystemreplacemsgSslvpnPath = "system.replacemsg/sslvpn/"

type SystemreplacemsgSslvpn struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
