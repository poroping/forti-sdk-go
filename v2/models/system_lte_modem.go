package models

const SystemLteModemPath = "system/lte-modem/"

type SystemLteModem struct {
	Apn           *string `json:"apn,omitempty"`
	Authtype      *string `json:"authtype,omitempty"`
	ExtraInit     *string `json:"extra-init,omitempty"`
	HolddownTimer *int64  `json:"holddown-timer,omitempty"`
	Interface     *string `json:"interface,omitempty"`
	Mode          *string `json:"mode,omitempty"`
	ModemPort     *int64  `json:"modem-port,omitempty"`
	Passwd        *string `json:"passwd,omitempty"`
	Status        *string `json:"status,omitempty"`
	Username      *string `json:"username,omitempty"`
}
