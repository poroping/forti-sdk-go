package models

const System3gModemCustomPath = "system.3g-modem/custom/"

type System3gModemCustom struct {
	ClassId          *string  `json:"class-id,omitempty"`
	Fosid            *float64 `json:"fosid,omitempty"`
	InitString       *string  `json:"init-string,omitempty"`
	Model            *string  `json:"model,omitempty"`
	ModeswitchString *string  `json:"modeswitch-string,omitempty"`
	ProductId        *string  `json:"product-id,omitempty"`
	Vendor           *string  `json:"vendor,omitempty"`
	VendorId         *string  `json:"vendor-id,omitempty"`
}
