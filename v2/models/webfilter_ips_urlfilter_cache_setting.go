package models

const WebfilterIpsUrlfilterCacheSettingPath = "webfilter/ips-urlfilter-cache-setting/"

type WebfilterIpsUrlfilterCacheSetting struct {
	DnsRetryInterval *int64 `json:"dns-retry-interval,omitempty"`
	ExtendedTtl      *int64 `json:"extended-ttl,omitempty"`
}
