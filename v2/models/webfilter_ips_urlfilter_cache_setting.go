package models

const WebfilterIpsUrlfilterCacheSettingPath = "webfilter/ips-urlfilter-cache-setting/"

type WebfilterIpsUrlfilterCacheSetting struct {
	DnsRetryInterval *float64 `json:"dns-retry-interval,omitempty"`
	ExtendedTtl      *float64 `json:"extended-ttl,omitempty"`
}
