package models

const SystemPtpPath = "system/ptp/"

type SystemPtp struct {
	DelayMechanism  *string                     `json:"delay-mechanism,omitempty"`
	Interface       *string                     `json:"interface,omitempty"`
	Mode            *string                     `json:"mode,omitempty"`
	RequestInterval *int64                      `json:"request-interval,omitempty"`
	ServerInterface *[]SystemPtpServerInterface `json:"server-interface,omitempty"`
	ServerMode      *string                     `json:"server-mode,omitempty"`
	Status          *string                     `json:"status,omitempty"`
}

type SystemPtpServerInterface struct {
	DelayMechanism      *string `json:"delay-mechanism,omitempty"`
	Id                  *int64  `json:"id,omitempty"`
	ServerInterfaceName *string `json:"server-interface-name,omitempty"`
}
