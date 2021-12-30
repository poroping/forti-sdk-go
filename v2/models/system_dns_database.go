package models

const SystemDnsDatabasePath = "system/dns-database/"

type SystemDnsDatabase struct {
	AllowTransfer *string                      `json:"allow-transfer,omitempty"`
	Authoritative *string                      `json:"authoritative,omitempty"`
	Contact       *string                      `json:"contact,omitempty"`
	DnsEntry      *[]SystemDnsDatabaseDnsEntry `json:"dns-entry,omitempty"`
	Domain        *string                      `json:"domain,omitempty"`
	Forwarder     *string                      `json:"forwarder,omitempty"`
	IpMaster      *string                      `json:"ip-master,omitempty"`
	IpPrimary     *string                      `json:"ip-primary,omitempty"`
	Name          *string                      `json:"name,omitempty"`
	PrimaryName   *string                      `json:"primary-name,omitempty"`
	RrMax         *int64                       `json:"rr-max,omitempty"`
	SourceIp      *string                      `json:"source-ip,omitempty"`
	Status        *string                      `json:"status,omitempty"`
	Ttl           *int64                       `json:"ttl,omitempty"`
	Type          *string                      `json:"type,omitempty"`
	View          *string                      `json:"view,omitempty"`
}

type SystemDnsDatabaseDnsEntry struct {
	CanonicalName *string `json:"canonical-name,omitempty"`
	Hostname      *string `json:"hostname,omitempty"`
	Id            *int64  `json:"id,omitempty"`
	Ip            *string `json:"ip,omitempty"`
	Ipv6          *string `json:"ipv6,omitempty"`
	Preference    *int64  `json:"preference,omitempty"`
	Status        *string `json:"status,omitempty"`
	Ttl           *int64  `json:"ttl,omitempty"`
	Type          *string `json:"type,omitempty"`
}
