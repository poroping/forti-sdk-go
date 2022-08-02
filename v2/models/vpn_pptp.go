package models

const VpnPptpPath = "vpn/pptp/"

type VpnPptp struct {
	Eip     *string `json:"eip,omitempty"`
	IpMode  *string `json:"ip-mode,omitempty"`
	LocalIp *string `json:"local-ip,omitempty"`
	Sip     *string `json:"sip,omitempty"`
	Status  *string `json:"status,omitempty"`
	Usrgrp  *string `json:"usrgrp,omitempty"`
}
