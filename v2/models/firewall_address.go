package models

const FirewallAddressPath = "firewall/address/"

type FirewallAddress struct {
	AllowRouting        *string                     `json:"allow-routing,omitempty"`
	AssociatedInterface *string                     `json:"associated-interface,omitempty"`
	CacheTtl            *int64                      `json:"cache-ttl,omitempty"`
	ClearpassSpt        *string                     `json:"clearpass-spt,omitempty"`
	Color               *int64                      `json:"color,omitempty"`
	Comment             *string                     `json:"comment,omitempty"`
	Country             *string                     `json:"country,omitempty"`
	EndIp               *string                     `json:"end-ip,omitempty"`
	EndMac              *string                     `json:"end-mac,omitempty"`
	EpgName             *string                     `json:"epg-name,omitempty"`
	FabricObject        *string                     `json:"fabric-object,omitempty"`
	Filter              *string                     `json:"filter,omitempty"`
	Fqdn                *string                     `json:"fqdn,omitempty"`
	FssoGroup           *[]FirewallAddressFssoGroup `json:"fsso-group,omitempty"`
	Interface           *string                     `json:"interface,omitempty"`
	List                *[]FirewallAddressList      `json:"list,omitempty"`
	Macaddr             *[]FirewallAddressMacaddr   `json:"macaddr,omitempty"`
	Name                *string                     `json:"name,omitempty"`
	NodeIpOnly          *string                     `json:"node-ip-only,omitempty"`
	ObjId               *string                     `json:"obj-id,omitempty"`
	ObjTag              *string                     `json:"obj-tag,omitempty"`
	ObjType             *string                     `json:"obj-type,omitempty"`
	Organization        *string                     `json:"organization,omitempty"`
	PolicyGroup         *string                     `json:"policy-group,omitempty"`
	Sdn                 *string                     `json:"sdn,omitempty"`
	SdnAddrType         *string                     `json:"sdn-addr-type,omitempty"`
	SdnTag              *string                     `json:"sdn-tag,omitempty"`
	StartIp             *string                     `json:"start-ip,omitempty"`
	StartMac            *string                     `json:"start-mac,omitempty"`
	SubType             *string                     `json:"sub-type,omitempty"`
	Subnet              *string                     `json:"subnet,omitempty"`
	SubnetName          *string                     `json:"subnet-name,omitempty"`
	TagDetectionLevel   *string                     `json:"tag-detection-level,omitempty"`
	TagType             *string                     `json:"tag-type,omitempty"`
	Tagging             *[]FirewallAddressTagging   `json:"tagging,omitempty"`
	Tenant              *string                     `json:"tenant,omitempty"`
	Type                *string                     `json:"type,omitempty"`
	Uuid                *string                     `json:"uuid,omitempty"`
	Wildcard            *string                     `json:"wildcard,omitempty"`
	WildcardFqdn        *string                     `json:"wildcard-fqdn,omitempty"`
}

const FirewallAddressFssoGroupPath = "firewall/address/fsso-group/"

type FirewallAddressFssoGroup struct {
	Name *string `json:"name,omitempty"`
}

const FirewallAddressListPath = "firewall/address/list/"

type FirewallAddressList struct {
	Ip *string `json:"ip,omitempty"`
}

const FirewallAddressMacaddrPath = "firewall/address/macaddr/"

type FirewallAddressMacaddr struct {
	Macaddr *string `json:"macaddr,omitempty"`
}

const FirewallAddressTaggingPath = "firewall/address/tagging/"

type FirewallAddressTagging struct {
	Category *string                       `json:"category,omitempty"`
	Name     *string                       `json:"name,omitempty"`
	Tags     *[]FirewallAddressTaggingTags `json:"tags,omitempty"`
}

const FirewallAddressTaggingTagsPath = "firewall/address/tagging/tags/"

type FirewallAddressTaggingTags struct {
	Name *string `json:"name,omitempty"`
}
