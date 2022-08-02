package models

const CertificateRemotePath = "certificate/remote/"

type CertificateRemote struct {
	Name   *string `json:"name,omitempty"`
	Range  *string `json:"range,omitempty"`
	Remote *string `json:"remote,omitempty"`
	Source *string `json:"source,omitempty"`
}
