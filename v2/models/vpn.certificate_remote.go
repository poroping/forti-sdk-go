package models

const VpnCertificateRemotePath = "vpn.certificate/remote/"

type VpnCertificateRemote struct {
	Name   *string `json:"name,omitempty"`
	Range  *string `json:"range,omitempty"`
	Remote *string `json:"remote,omitempty"`
	Source *string `json:"source,omitempty"`
}
