package models

const SystemReplacemsgFtpPath = "system.replacemsg/ftp/"

type SystemReplacemsgFtp struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
