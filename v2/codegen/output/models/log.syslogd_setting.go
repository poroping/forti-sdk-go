package models

const LogsyslogdSettingPath = "log.syslogd/setting/"

type LogsyslogdSetting struct {
	Certificate           *string                            `json:"certificate,omitempty"`
	CustomFieldName       []LogsyslogdSettingCustomFieldName `json:"custom-field-name,omitempty"`
	EncAlgorithm          *string                            `json:"enc-algorithm,omitempty"`
	Facility              *string                            `json:"facility,omitempty"`
	Format                *string                            `json:"format,omitempty"`
	Interface             *string                            `json:"interface,omitempty"`
	InterfaceSelectMethod *string                            `json:"interface-select-method,omitempty"`
	MaxLogRate            *int64                             `json:"max-log-rate,omitempty"`
	Mode                  *string                            `json:"mode,omitempty"`
	Port                  *int64                             `json:"port,omitempty"`
	Priority              *string                            `json:"priority,omitempty"`
	Server                *string                            `json:"server,omitempty"`
	SourceIp              *string                            `json:"source-ip,omitempty"`
	SslMinProtoVersion    *string                            `json:"ssl-min-proto-version,omitempty"`
	Status                *string                            `json:"status,omitempty"`
}

type LogsyslogdSettingCustomFieldName struct {
	Custom *string `json:"custom,omitempty"`
	Id     *int64  `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
}
