package models

const VpnSslClientPath = "vpn.ssl/client/"

type VpnSslClient struct {
	Certificate *string `json:"certificate,omitempty"`
	ClassId     *int64  `json:"class-id,omitempty"`
	Comment     *string `json:"comment,omitempty"`
	Distance    *int64  `json:"distance,omitempty"`
	Interface   *string `json:"interface,omitempty"`
	Name        *string `json:"name,omitempty"`
	Peer        *string `json:"peer,omitempty"`
	Port        *int64  `json:"port,omitempty"`
	Priority    *int64  `json:"priority,omitempty"`
	Psk         *string `json:"psk,omitempty"`
	Realm       *string `json:"realm,omitempty"`
	Server      *string `json:"server,omitempty"`
	SourceIp    *string `json:"source-ip,omitempty"`
	Status      *string `json:"status,omitempty"`
	User        *string `json:"user,omitempty"`
}
