package models

const FirewallAddress6Path = "firewall/address6/"

type FirewallAddress6 struct {
	CacheTtl      *int64                           `json:"cache-ttl,omitempty"`
	Color         *int64                           `json:"color,omitempty"`
	Comment       *string                          `json:"comment,omitempty"`
	Country       *string                          `json:"country,omitempty"`
	EndIp         *string                          `json:"end-ip,omitempty"`
	EndMac        *string                          `json:"end-mac,omitempty"`
	FabricObject  *string                          `json:"fabric-object,omitempty"`
	Fqdn          *string                          `json:"fqdn,omitempty"`
	Host          *string                          `json:"host,omitempty"`
	HostType      *string                          `json:"host-type,omitempty"`
	Ip6           *string                          `json:"ip6,omitempty"`
	List          *[]FirewallAddress6List          `json:"list,omitempty"`
	Macaddr       *[]FirewallAddress6Macaddr       `json:"macaddr,omitempty"`
	Name          *string                          `json:"name,omitempty"`
	ObjId         *string                          `json:"obj-id,omitempty"`
	Sdn           *string                          `json:"sdn,omitempty"`
	StartIp       *string                          `json:"start-ip,omitempty"`
	StartMac      *string                          `json:"start-mac,omitempty"`
	SubnetSegment *[]FirewallAddress6SubnetSegment `json:"subnet-segment,omitempty"`
	Tagging       *[]FirewallAddress6Tagging       `json:"tagging,omitempty"`
	Template      *string                          `json:"template,omitempty"`
	Type          *string                          `json:"type,omitempty"`
	Uuid          *string                          `json:"uuid,omitempty"`
}

const FirewallAddress6ListPath = "firewall/address6/list/"

type FirewallAddress6List struct {
	Ip *string `json:"ip,omitempty"`
}

const FirewallAddress6MacaddrPath = "firewall/address6/macaddr/"

type FirewallAddress6Macaddr struct {
	Macaddr *string `json:"macaddr,omitempty"`
}

const FirewallAddress6SubnetSegmentPath = "firewall/address6/subnet-segment/"

type FirewallAddress6SubnetSegment struct {
	Name  *string `json:"name,omitempty"`
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

const FirewallAddress6TaggingPath = "firewall/address6/tagging/"

type FirewallAddress6Tagging struct {
	Category *string                        `json:"category,omitempty"`
	Name     *string                        `json:"name,omitempty"`
	Tags     *[]FirewallAddress6TaggingTags `json:"tags,omitempty"`
}

const FirewallAddress6TaggingTagsPath = "firewall/address6/tagging/tags/"

type FirewallAddress6TaggingTags struct {
	Name *string `json:"name,omitempty"`
}
