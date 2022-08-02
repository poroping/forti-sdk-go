package models

const System3GModemCustomPath = "system.3g-modem/custom/"

type System3GModemCustom struct {
	ClassId          *string `json:"class-id,omitempty"`
	Id               *int64  `json:"id,omitempty"`
	InitString       *string `json:"init-string,omitempty"`
	Model            *string `json:"model,omitempty"`
	ModeswitchString *string `json:"modeswitch-string,omitempty"`
	ProductId        *string `json:"product-id,omitempty"`
	Vendor           *string `json:"vendor,omitempty"`
	VendorId         *string `json:"vendor-id,omitempty"`
}
