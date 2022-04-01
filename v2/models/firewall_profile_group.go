package models

const FirewallProfileGroupPath = "firewall/profile-group/"

type FirewallProfileGroup struct {
	ApplicationList        *string `json:"application-list,omitempty"`
	AvProfile              *string `json:"av-profile,omitempty"`
	CifsProfile            *string `json:"cifs-profile,omitempty"`
	DlpProfile             *string `json:"dlp-profile,omitempty"`
	DlpSensor              *string `json:"dlp-sensor,omitempty"`
	DnsfilterProfile       *string `json:"dnsfilter-profile,omitempty"`
	EmailfilterProfile     *string `json:"emailfilter-profile,omitempty"`
	FileFilterProfile      *string `json:"file-filter-profile,omitempty"`
	IcapProfile            *string `json:"icap-profile,omitempty"`
	IpsSensor              *string `json:"ips-sensor,omitempty"`
	Name                   *string `json:"name,omitempty"`
	ProfileProtocolOptions *string `json:"profile-protocol-options,omitempty"`
	SctpFilterProfile      *string `json:"sctp-filter-profile,omitempty"`
	SshFilterProfile       *string `json:"ssh-filter-profile,omitempty"`
	SslSshProfile          *string `json:"ssl-ssh-profile,omitempty"`
	VideofilterProfile     *string `json:"videofilter-profile,omitempty"`
	VoipProfile            *string `json:"voip-profile,omitempty"`
	WafProfile             *string `json:"waf-profile,omitempty"`
	WebfilterProfile       *string `json:"webfilter-profile,omitempty"`
}
