package models

const FirewallSslServerPath = "firewall/ssl-server/"

type FirewallSslServer struct {
	AddHeaderXForwardedProto *string  `json:"add-header-x-forwarded-proto,omitempty"`
	Ip                       *string  `json:"ip,omitempty"`
	MappedPort               *float64 `json:"mapped-port,omitempty"`
	Name                     *string  `json:"name,omitempty"`
	Port                     *float64 `json:"port,omitempty"`
	SslAlgorithm             *string  `json:"ssl-algorithm,omitempty"`
	SslCert                  *string  `json:"ssl-cert,omitempty"`
	SslClientRenegotiation   *string  `json:"ssl-client-renegotiation,omitempty"`
	SslDhBits                *string  `json:"ssl-dh-bits,omitempty"`
	SslMaxVersion            *string  `json:"ssl-max-version,omitempty"`
	SslMinVersion            *string  `json:"ssl-min-version,omitempty"`
	SslMode                  *string  `json:"ssl-mode,omitempty"`
	SslSendEmptyFrags        *string  `json:"ssl-send-empty-frags,omitempty"`
	UrlRewrite               *string  `json:"url-rewrite,omitempty"`
}
