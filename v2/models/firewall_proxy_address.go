package models

const FirewallProxyAddressPath = "firewall/proxy-address/"

type FirewallProxyAddress struct {
	CaseSensitivity *string                           `json:"case-sensitivity,omitempty"`
	Category        []FirewallProxyAddressCategory    `json:"category,omitempty"`
	Color           *int64                            `json:"color,omitempty"`
	Comment         *string                           `json:"comment,omitempty"`
	Header          *string                           `json:"header,omitempty"`
	HeaderGroup     []FirewallProxyAddressHeaderGroup `json:"header-group,omitempty"`
	HeaderName      *string                           `json:"header-name,omitempty"`
	Host            *string                           `json:"host,omitempty"`
	HostRegex       *string                           `json:"host-regex,omitempty"`
	Method          *string                           `json:"method,omitempty"`
	Name            *string                           `json:"name,omitempty"`
	Path            *string                           `json:"path,omitempty"`
	Query           *string                           `json:"query,omitempty"`
	Referrer        *string                           `json:"referrer,omitempty"`
	Tagging         []FirewallProxyAddressTagging     `json:"tagging,omitempty"`
	Type            *string                           `json:"type,omitempty"`
	Ua              *string                           `json:"ua,omitempty"`
	Uuid            *string                           `json:"uuid,omitempty"`
}

type FirewallProxyAddressCategory struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallProxyAddressHeaderGroup struct {
	CaseSensitivity *string `json:"case-sensitivity,omitempty"`
	Header          *string `json:"header,omitempty"`
	HeaderName      *string `json:"header-name,omitempty"`
	Id              *int64  `json:"id,omitempty"`
}

type FirewallProxyAddressTagging struct {
	Category *string                           `json:"category,omitempty"`
	Name     *string                           `json:"name,omitempty"`
	Tags     []FirewallProxyAddressTaggingTags `json:"tags,omitempty"`
}

type FirewallProxyAddressTaggingTags struct {
	Name *string `json:"name,omitempty"`
}
