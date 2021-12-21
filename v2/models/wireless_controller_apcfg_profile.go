package models

const WirelessControllerApcfgProfilePath = "wireless-controller/apcfg-profile/"

type WirelessControllerApcfgProfile struct {
	AcIp        *string                                      `json:"ac-ip,omitempty"`
	AcPort      *float64                                     `json:"ac-port,omitempty"`
	AcTimer     *float64                                     `json:"ac-timer,omitempty"`
	AcType      *string                                      `json:"ac-type,omitempty"`
	ApFamily    *string                                      `json:"ap-family,omitempty"`
	CommandList *[]WirelessControllerApcfgProfileCommandList `json:"command-list,omitempty"`
	Comment     *string                                      `json:"comment,omitempty"`
	Name        *string                                      `json:"name,omitempty"`
}

type WirelessControllerApcfgProfileCommandList struct {
	Id          *float64 `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	PasswdValue *string  `json:"passwd-value,omitempty"`
	Type        *string  `json:"type,omitempty"`
	Value       *string  `json:"value,omitempty"`
}
