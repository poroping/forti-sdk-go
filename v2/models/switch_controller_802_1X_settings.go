package models

const SwitchController8021XSettingsPath = "switch-controller/802-1x-settings/"

type SwitchController8021XSettings struct {
	LinkDownAuth     *string `json:"link-down-auth,omitempty"`
	MaxReauthAttempt *int64  `json:"max-reauth-attempt,omitempty"`
	ReauthPeriod     *int64  `json:"reauth-period,omitempty"`
	TxPeriod         *int64  `json:"tx-period,omitempty"`
}
