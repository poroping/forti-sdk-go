package models

const UserCertificatePath = "user/certificate/"

type UserCertificate struct {
	CommonName *string `json:"common-name,omitempty"`
	Fosid      *int64  `json:"fosid,omitempty"`
	Issuer     *string `json:"issuer,omitempty"`
	Name       *string `json:"name,omitempty"`
	Status     *string `json:"status,omitempty"`
	Type       *string `json:"type,omitempty"`
}
