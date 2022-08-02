package models

const DnsfilterProfilePath = "dnsfilter/profile/"

type DnsfilterProfile struct {
	BlockAction         *string                                `json:"block-action,omitempty"`
	BlockBotnet         *string                                `json:"block-botnet,omitempty"`
	Comment             *string                                `json:"comment,omitempty"`
	DnsTranslation      *[]DnsfilterProfileDnsTranslation      `json:"dns-translation,omitempty"`
	DomainFilter        *DnsfilterProfileDomainFilter          `json:"domain-filter,omitempty"`
	ExternalIpBlocklist *[]DnsfilterProfileExternalIpBlocklist `json:"external-ip-blocklist,omitempty"`
	FtgdDns             *DnsfilterProfileFtgdDns               `json:"ftgd-dns,omitempty"`
	LogAllDomain        *string                                `json:"log-all-domain,omitempty"`
	Name                *string                                `json:"name,omitempty"`
	RedirectPortal      *string                                `json:"redirect-portal,omitempty"`
	RedirectPortal6     *string                                `json:"redirect-portal6,omitempty"`
	SafeSearch          *string                                `json:"safe-search,omitempty"`
	SdnsDomainLog       *string                                `json:"sdns-domain-log,omitempty"`
	SdnsFtgdErrLog      *string                                `json:"sdns-ftgd-err-log,omitempty"`
	YoutubeRestrict     *string                                `json:"youtube-restrict,omitempty"`
}

const DnsfilterProfileDnsTranslationPath = "dnsfilter/profile/dns-translation/"

type DnsfilterProfileDnsTranslation struct {
	AddrType *string `json:"addr-type,omitempty"`
	Dst      *string `json:"dst,omitempty"`
	Dst6     *string `json:"dst6,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Netmask  *string `json:"netmask,omitempty"`
	Prefix   *int64  `json:"prefix,omitempty"`
	Src      *string `json:"src,omitempty"`
	Src6     *string `json:"src6,omitempty"`
	Status   *string `json:"status,omitempty"`
}

const DnsfilterProfileDomainFilterPath = "dnsfilter/profile/domain-filter/"

type DnsfilterProfileDomainFilter struct {
	DomainFilterTable *int64 `json:"domain-filter-table,omitempty"`
}

const DnsfilterProfileExternalIpBlocklistPath = "dnsfilter/profile/external-ip-blocklist/"

type DnsfilterProfileExternalIpBlocklist struct {
	Name *string `json:"name,omitempty"`
}

const DnsfilterProfileFtgdDnsPath = "dnsfilter/profile/ftgd-dns/"

type DnsfilterProfileFtgdDns struct {
	Filters *[]DnsfilterProfileFtgdDnsFilters `json:"filters,omitempty"`
	Options *string                           `json:"options,omitempty"`
}

const DnsfilterProfileFtgdDnsFiltersPath = "dnsfilter/profile/ftgd-dns/filters/"

type DnsfilterProfileFtgdDnsFilters struct {
	Action   *string `json:"action,omitempty"`
	Category *int64  `json:"category,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Log      *string `json:"log,omitempty"`
}
