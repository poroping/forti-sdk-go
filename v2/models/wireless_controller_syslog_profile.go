package models

const WirelessControllerSyslogProfilePath = "wireless-controller/syslog-profile/"

type WirelessControllerSyslogProfile struct {
	Comment        *string  `json:"comment,omitempty"`
	LogLevel       *string  `json:"log-level,omitempty"`
	Name           *string  `json:"name,omitempty"`
	ServerAddrType *string  `json:"server-addr-type,omitempty"`
	ServerFqdn     *string  `json:"server-fqdn,omitempty"`
	ServerIp       *string  `json:"server-ip,omitempty"`
	ServerPort     *float64 `json:"server-port,omitempty"`
	ServerStatus   *string  `json:"server-status,omitempty"`
}
