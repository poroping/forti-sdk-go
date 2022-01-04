package models

const SystemReplacemsgAlertmailPath = "system.replacemsg/alertmail/"

type SystemReplacemsgAlertmail struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
