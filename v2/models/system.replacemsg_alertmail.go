package models

const SystemreplacemsgAlertmailPath = "system/replacemsg/alertmail/"

type SystemreplacemsgAlertmail struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
