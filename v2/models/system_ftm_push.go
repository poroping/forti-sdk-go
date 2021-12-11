package models

const SystemFtmPushPath = "system/ftm-push/"

type SystemFtmPush struct {
	Server     *string `json:"server,omitempty"`
	ServerCert *string `json:"server-cert,omitempty"`
	ServerIp   *string `json:"server-ip,omitempty"`
	ServerPort *int64  `json:"server-port,omitempty"`
	Status     *string `json:"status,omitempty"`
}
