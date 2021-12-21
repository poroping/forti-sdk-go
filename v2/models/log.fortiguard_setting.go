package models

const LogfortiguardSettingPath = "log.fortiguard/setting/"

type LogfortiguardSetting struct {
	AccessConfig          *string  `json:"access-config,omitempty"`
	ConnTimeout           *float64 `json:"conn-timeout,omitempty"`
	EncAlgorithm          *string  `json:"enc-algorithm,omitempty"`
	Interface             *string  `json:"interface,omitempty"`
	InterfaceSelectMethod *string  `json:"interface-select-method,omitempty"`
	MaxLogRate            *float64 `json:"max-log-rate,omitempty"`
	Priority              *string  `json:"priority,omitempty"`
	SourceIp              *string  `json:"source-ip,omitempty"`
	SslMinProtoVersion    *string  `json:"ssl-min-proto-version,omitempty"`
	Status                *string  `json:"status,omitempty"`
	UploadDay             *string  `json:"upload-day,omitempty"`
	UploadInterval        *string  `json:"upload-interval,omitempty"`
	UploadOption          *string  `json:"upload-option,omitempty"`
	UploadTime            *string  `json:"upload-time,omitempty"`
}
