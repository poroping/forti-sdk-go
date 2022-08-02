package models

const SystemSmsServerPath = "system/sms-server/"

type SystemSmsServer struct {
	MailServer *string `json:"mail-server,omitempty"`
	Name       *string `json:"name,omitempty"`
}
