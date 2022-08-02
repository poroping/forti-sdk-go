package models

const ExtenderModemStatusPath = "extender/modem-status/"

type ExtenderModemStatus struct {
	Sn *string `json:"<sn>,omitempty"`
}
