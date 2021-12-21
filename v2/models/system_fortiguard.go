package models

const SystemFortiguardPath = "system/fortiguard/"

type SystemFortiguard struct {
	AntispamCache                   *string  `json:"antispam-cache,omitempty"`
	AntispamCacheMpercent           *float64 `json:"antispam-cache-mpercent,omitempty"`
	AntispamCacheTtl                *float64 `json:"antispam-cache-ttl,omitempty"`
	AntispamExpiration              *float64 `json:"antispam-expiration,omitempty"`
	AntispamForceOff                *string  `json:"antispam-force-off,omitempty"`
	AntispamLicense                 *float64 `json:"antispam-license,omitempty"`
	AntispamTimeout                 *float64 `json:"antispam-timeout,omitempty"`
	AnycastSdnsServerIp             *string  `json:"anycast-sdns-server-ip,omitempty"`
	AnycastSdnsServerPort           *float64 `json:"anycast-sdns-server-port,omitempty"`
	AutoJoinForticloud              *string  `json:"auto-join-forticloud,omitempty"`
	DdnsServerIp                    *string  `json:"ddns-server-ip,omitempty"`
	DdnsServerIp6                   *string  `json:"ddns-server-ip6,omitempty"`
	DdnsServerPort                  *float64 `json:"ddns-server-port,omitempty"`
	FortiguardAnycast               *string  `json:"fortiguard-anycast,omitempty"`
	FortiguardAnycastSource         *string  `json:"fortiguard-anycast-source,omitempty"`
	Interface                       *string  `json:"interface,omitempty"`
	InterfaceSelectMethod           *string  `json:"interface-select-method,omitempty"`
	LoadBalanceServers              *float64 `json:"load-balance-servers,omitempty"`
	OutbreakPreventionCache         *string  `json:"outbreak-prevention-cache,omitempty"`
	OutbreakPreventionCacheMpercent *float64 `json:"outbreak-prevention-cache-mpercent,omitempty"`
	OutbreakPreventionCacheTtl      *float64 `json:"outbreak-prevention-cache-ttl,omitempty"`
	OutbreakPreventionExpiration    *float64 `json:"outbreak-prevention-expiration,omitempty"`
	OutbreakPreventionForceOff      *string  `json:"outbreak-prevention-force-off,omitempty"`
	OutbreakPreventionLicense       *float64 `json:"outbreak-prevention-license,omitempty"`
	OutbreakPreventionTimeout       *float64 `json:"outbreak-prevention-timeout,omitempty"`
	PersistentConnection            *string  `json:"persistent-connection,omitempty"`
	Port                            *string  `json:"port,omitempty"`
	Protocol                        *string  `json:"protocol,omitempty"`
	ProxyPassword                   *string  `json:"proxy-password,omitempty"`
	ProxyServerIp                   *string  `json:"proxy-server-ip,omitempty"`
	ProxyServerPort                 *float64 `json:"proxy-server-port,omitempty"`
	ProxyUsername                   *string  `json:"proxy-username,omitempty"`
	SandboxRegion                   *string  `json:"sandbox-region,omitempty"`
	SdnsOptions                     *string  `json:"sdns-options,omitempty"`
	SdnsServerIp                    *string  `json:"sdns-server-ip,omitempty"`
	SdnsServerPort                  *float64 `json:"sdns-server-port,omitempty"`
	SourceIp                        *string  `json:"source-ip,omitempty"`
	SourceIp6                       *string  `json:"source-ip6,omitempty"`
	UpdateBuildProxy                *string  `json:"update-build-proxy,omitempty"`
	UpdateExtdb                     *string  `json:"update-extdb,omitempty"`
	UpdateFfdb                      *string  `json:"update-ffdb,omitempty"`
	UpdateServerLocation            *string  `json:"update-server-location,omitempty"`
	UpdateUwdb                      *string  `json:"update-uwdb,omitempty"`
	VideofilterExpiration           *float64 `json:"videofilter-expiration,omitempty"`
	VideofilterLicense              *float64 `json:"videofilter-license,omitempty"`
	WebfilterCache                  *string  `json:"webfilter-cache,omitempty"`
	WebfilterCacheTtl               *float64 `json:"webfilter-cache-ttl,omitempty"`
	WebfilterExpiration             *float64 `json:"webfilter-expiration,omitempty"`
	WebfilterForceOff               *string  `json:"webfilter-force-off,omitempty"`
	WebfilterLicense                *float64 `json:"webfilter-license,omitempty"`
	WebfilterTimeout                *float64 `json:"webfilter-timeout,omitempty"`
}
