package models

const SystemFortiaiPath = "system/fortiai/"

type SystemFortiai struct {
	Interface             *string `json:"interface,omitempty"`
	InterfaceSelectMethod *string `json:"interface-select-method,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
	Status                *string `json:"status,omitempty"`
}
