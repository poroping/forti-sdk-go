package models

const WanoptProfilePath = "wanopt/profile/"

type WanoptProfile struct {
	AuthGroup   *string              `json:"auth-group,omitempty"`
	Cifs        *[]WanoptProfileCifs `json:"cifs,omitempty"`
	Comments    *string              `json:"comments,omitempty"`
	Ftp         *[]WanoptProfileFtp  `json:"ftp,omitempty"`
	Http        *[]WanoptProfileHttp `json:"http,omitempty"`
	Mapi        *[]WanoptProfileMapi `json:"mapi,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Tcp         *[]WanoptProfileTcp  `json:"tcp,omitempty"`
	Transparent *string              `json:"transparent,omitempty"`
}

type WanoptProfileCifs struct {
	ByteCaching    *string  `json:"byte-caching,omitempty"`
	LogTraffic     *string  `json:"log-traffic,omitempty"`
	Port           *float64 `json:"port,omitempty"`
	PreferChunking *string  `json:"prefer-chunking,omitempty"`
	ProtocolOpt    *string  `json:"protocol-opt,omitempty"`
	SecureTunnel   *string  `json:"secure-tunnel,omitempty"`
	Status         *string  `json:"status,omitempty"`
	TunnelSharing  *string  `json:"tunnel-sharing,omitempty"`
}

type WanoptProfileFtp struct {
	ByteCaching    *string  `json:"byte-caching,omitempty"`
	LogTraffic     *string  `json:"log-traffic,omitempty"`
	Port           *float64 `json:"port,omitempty"`
	PreferChunking *string  `json:"prefer-chunking,omitempty"`
	ProtocolOpt    *string  `json:"protocol-opt,omitempty"`
	SecureTunnel   *string  `json:"secure-tunnel,omitempty"`
	Ssl            *string  `json:"ssl,omitempty"`
	Status         *string  `json:"status,omitempty"`
	TunnelSharing  *string  `json:"tunnel-sharing,omitempty"`
}

type WanoptProfileHttp struct {
	ByteCaching        *string  `json:"byte-caching,omitempty"`
	LogTraffic         *string  `json:"log-traffic,omitempty"`
	Port               *float64 `json:"port,omitempty"`
	PreferChunking     *string  `json:"prefer-chunking,omitempty"`
	ProtocolOpt        *string  `json:"protocol-opt,omitempty"`
	SecureTunnel       *string  `json:"secure-tunnel,omitempty"`
	Ssl                *string  `json:"ssl,omitempty"`
	SslPort            *float64 `json:"ssl-port,omitempty"`
	Status             *string  `json:"status,omitempty"`
	TunnelNonHttp      *string  `json:"tunnel-non-http,omitempty"`
	TunnelSharing      *string  `json:"tunnel-sharing,omitempty"`
	UnknownHttpVersion *string  `json:"unknown-http-version,omitempty"`
}

type WanoptProfileMapi struct {
	ByteCaching   *string  `json:"byte-caching,omitempty"`
	LogTraffic    *string  `json:"log-traffic,omitempty"`
	Port          *float64 `json:"port,omitempty"`
	SecureTunnel  *string  `json:"secure-tunnel,omitempty"`
	Status        *string  `json:"status,omitempty"`
	TunnelSharing *string  `json:"tunnel-sharing,omitempty"`
}

type WanoptProfileTcp struct {
	ByteCaching    *string `json:"byte-caching,omitempty"`
	ByteCachingOpt *string `json:"byte-caching-opt,omitempty"`
	LogTraffic     *string `json:"log-traffic,omitempty"`
	Port           *string `json:"port,omitempty"`
	SecureTunnel   *string `json:"secure-tunnel,omitempty"`
	Ssl            *string `json:"ssl,omitempty"`
	SslPort        *string `json:"ssl-port,omitempty"`
	Status         *string `json:"status,omitempty"`
	TunnelSharing  *string `json:"tunnel-sharing,omitempty"`
}
