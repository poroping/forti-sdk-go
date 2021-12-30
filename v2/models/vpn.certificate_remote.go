package models

const VpncertificateRemotePath = "vpn/certificate/remote/"

type VpncertificateRemote struct {
	Name   *string `json:"name,omitempty"`
	Range  *string `json:"range,omitempty"`
	Remote *string `json:"remote,omitempty"`
	Source *string `json:"source,omitempty"`
}
