package models

const Logsyslogd4SettingPath = "log.syslogd4/setting/"

type Logsyslogd4Setting struct {
	Certificate           *string                              `json:"certificate,omitempty"`
	CustomFieldName       *[]Logsyslogd4SettingCustomFieldName `json:"custom-field-name,omitempty"`
	EncAlgorithm          *string                              `json:"enc-algorithm,omitempty"`
	Facility              *string                              `json:"facility,omitempty"`
	Format                *string                              `json:"format,omitempty"`
	Interface             *string                              `json:"interface,omitempty"`
	InterfaceSelectMethod *string                              `json:"interface-select-method,omitempty"`
	MaxLogRate            *float64                             `json:"max-log-rate,omitempty"`
	Mode                  *string                              `json:"mode,omitempty"`
	Port                  *float64                             `json:"port,omitempty"`
	Priority              *string                              `json:"priority,omitempty"`
	Server                *string                              `json:"server,omitempty"`
	SourceIp              *string                              `json:"source-ip,omitempty"`
	SslMinProtoVersion    *string                              `json:"ssl-min-proto-version,omitempty"`
	Status                *string                              `json:"status,omitempty"`
}

type Logsyslogd4SettingCustomFieldName struct {
	Custom *string  `json:"custom,omitempty"`
	Id     *float64 `json:"id,omitempty"`
	Name   *string  `json:"name,omitempty"`
}
