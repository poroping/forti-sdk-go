package models

const VpnsslwebRealmPath = "vpn/ssl/web/realm/"

type VpnsslwebRealm struct {
	LoginPage             *string `json:"login-page,omitempty"`
	MaxConcurrentUser     *int64  `json:"max-concurrent-user,omitempty"`
	NasIp                 *string `json:"nas-ip,omitempty"`
	RadiusPort            *int64  `json:"radius-port,omitempty"`
	RadiusServer          *string `json:"radius-server,omitempty"`
	UrlPath               *string `json:"url-path,omitempty"`
	VirtualHost           *string `json:"virtual-host,omitempty"`
	VirtualHostOnly       *string `json:"virtual-host-only,omitempty"`
	VirtualHostServerCert *string `json:"virtual-host-server-cert,omitempty"`
}
