package models

const VpnsslClientPath = "vpn.ssl/client/"

type VpnsslClient struct {
	Certificate *string  `json:"certificate,omitempty"`
	Comment     *string  `json:"comment,omitempty"`
	Distance    *float64 `json:"distance,omitempty"`
	Interface   *string  `json:"interface,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Peer        *string  `json:"peer,omitempty"`
	Port        *float64 `json:"port,omitempty"`
	Priority    *float64 `json:"priority,omitempty"`
	Psk         *string  `json:"psk,omitempty"`
	Realm       *string  `json:"realm,omitempty"`
	Server      *string  `json:"server,omitempty"`
	SourceIp    *string  `json:"source-ip,omitempty"`
	Status      *string  `json:"status,omitempty"`
	User        *string  `json:"user,omitempty"`
}
