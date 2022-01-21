package models

const FtpProxyExplicitPath = "ftp-proxy/explicit/"

type FtpProxyExplicit struct {
	IncomingIp       *string `json:"incoming-ip,omitempty"`
	IncomingPort     *string `json:"incoming-port,omitempty"`
	OutgoingIp       *string `json:"outgoing-ip,omitempty"`
	SecDefaultAction *string `json:"sec-default-action,omitempty"`
	Ssl              *string `json:"ssl,omitempty"`
	SslAlgorithm     *string `json:"ssl-algorithm,omitempty"`
	SslCert          *string `json:"ssl-cert,omitempty"`
	SslDhBits        *string `json:"ssl-dh-bits,omitempty"`
	Status           *string `json:"status,omitempty"`
}
