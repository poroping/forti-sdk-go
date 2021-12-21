package models

const WanoptWebcachePath = "wanopt/webcache/"

type WanoptWebcache struct {
	AlwaysRevalidate  *string  `json:"always-revalidate,omitempty"`
	CacheByDefault    *string  `json:"cache-by-default,omitempty"`
	CacheCookie       *string  `json:"cache-cookie,omitempty"`
	CacheExpired      *string  `json:"cache-expired,omitempty"`
	DefaultTtl        *float64 `json:"default-ttl,omitempty"`
	External          *string  `json:"external,omitempty"`
	FreshFactor       *float64 `json:"fresh-factor,omitempty"`
	HostValidate      *string  `json:"host-validate,omitempty"`
	IgnoreConditional *string  `json:"ignore-conditional,omitempty"`
	IgnoreIeReload    *string  `json:"ignore-ie-reload,omitempty"`
	IgnoreIms         *string  `json:"ignore-ims,omitempty"`
	IgnorePnc         *string  `json:"ignore-pnc,omitempty"`
	MaxObjectSize     *float64 `json:"max-object-size,omitempty"`
	MaxTtl            *float64 `json:"max-ttl,omitempty"`
	MinTtl            *float64 `json:"min-ttl,omitempty"`
	NegRespTime       *float64 `json:"neg-resp-time,omitempty"`
	RevalPnc          *string  `json:"reval-pnc,omitempty"`
}
