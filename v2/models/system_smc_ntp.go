package models

const SystemSmcNtpPath = "system/smc-ntp/"

type SystemSmcNtp struct {
	Channel      *int64                   `json:"channel,omitempty"`
	Ntpserver    *[]SystemSmcNtpNtpserver `json:"ntpserver,omitempty"`
	Ntpsync      *string                  `json:"ntpsync,omitempty"`
	Syncinterval *int64                   `json:"syncinterval,omitempty"`
}

const SystemSmcNtpNtpserverPath = "system/smc-ntp/ntpserver/"

type SystemSmcNtpNtpserver struct {
	Id     *int64  `json:"id,omitempty"`
	Server *string `json:"server,omitempty"`
}
