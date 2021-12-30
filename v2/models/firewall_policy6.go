package models

const FirewallPolicy6Path = "firewall/policy6/"

type FirewallPolicy6 struct {
	Action                  *string                           `json:"action,omitempty"`
	AntiReplay              *string                           `json:"anti-replay,omitempty"`
	AppCategory             *[]FirewallPolicy6AppCategory     `json:"app-category,omitempty"`
	AppGroup                *[]FirewallPolicy6AppGroup        `json:"app-group,omitempty"`
	Application             *[]FirewallPolicy6Application     `json:"application,omitempty"`
	ApplicationList         *string                           `json:"application-list,omitempty"`
	AutoAsicOffload         *string                           `json:"auto-asic-offload,omitempty"`
	AvProfile               *string                           `json:"av-profile,omitempty"`
	CifsProfile             *string                           `json:"cifs-profile,omitempty"`
	Comments                *string                           `json:"comments,omitempty"`
	CustomLogFields         *[]FirewallPolicy6CustomLogFields `json:"custom-log-fields,omitempty"`
	DecryptedTrafficMirror  *string                           `json:"decrypted-traffic-mirror,omitempty"`
	DiffservForward         *string                           `json:"diffserv-forward,omitempty"`
	DiffservReverse         *string                           `json:"diffserv-reverse,omitempty"`
	DiffservcodeForward     *string                           `json:"diffservcode-forward,omitempty"`
	DiffservcodeRev         *string                           `json:"diffservcode-rev,omitempty"`
	DlpSensor               *string                           `json:"dlp-sensor,omitempty"`
	DnsfilterProfile        *string                           `json:"dnsfilter-profile,omitempty"`
	Dsri                    *string                           `json:"dsri,omitempty"`
	Dstaddr                 *[]FirewallPolicy6Dstaddr         `json:"dstaddr,omitempty"`
	DstaddrNegate           *string                           `json:"dstaddr-negate,omitempty"`
	Dstintf                 *[]FirewallPolicy6Dstintf         `json:"dstintf,omitempty"`
	EmailfilterProfile      *string                           `json:"emailfilter-profile,omitempty"`
	FileFilterProfile       *string                           `json:"file-filter-profile,omitempty"`
	FirewallSessionDirty    *string                           `json:"firewall-session-dirty,omitempty"`
	Fixedport               *string                           `json:"fixedport,omitempty"`
	FssoGroups              *[]FirewallPolicy6FssoGroups      `json:"fsso-groups,omitempty"`
	Groups                  *[]FirewallPolicy6Groups          `json:"groups,omitempty"`
	HttpPolicyRedirect      *string                           `json:"http-policy-redirect,omitempty"`
	IcapProfile             *string                           `json:"icap-profile,omitempty"`
	Inbound                 *string                           `json:"inbound,omitempty"`
	InspectionMode          *string                           `json:"inspection-mode,omitempty"`
	Ippool                  *string                           `json:"ippool,omitempty"`
	IpsSensor               *string                           `json:"ips-sensor,omitempty"`
	Logtraffic              *string                           `json:"logtraffic,omitempty"`
	LogtrafficStart         *string                           `json:"logtraffic-start,omitempty"`
	Name                    *string                           `json:"name,omitempty"`
	Nat                     *string                           `json:"nat,omitempty"`
	Natinbound              *string                           `json:"natinbound,omitempty"`
	Natoutbound             *string                           `json:"natoutbound,omitempty"`
	NpAcceleration          *string                           `json:"np-acceleration,omitempty"`
	Outbound                *string                           `json:"outbound,omitempty"`
	PerIpShaper             *string                           `json:"per-ip-shaper,omitempty"`
	Policyid                *int64                            `json:"policyid,omitempty"`
	Poolname                *[]FirewallPolicy6Poolname        `json:"poolname,omitempty"`
	ProfileGroup            *string                           `json:"profile-group,omitempty"`
	ProfileProtocolOptions  *string                           `json:"profile-protocol-options,omitempty"`
	ProfileType             *string                           `json:"profile-type,omitempty"`
	ReplacemsgOverrideGroup *string                           `json:"replacemsg-override-group,omitempty"`
	Rsso                    *string                           `json:"rsso,omitempty"`
	Schedule                *string                           `json:"schedule,omitempty"`
	SendDenyPacket          *string                           `json:"send-deny-packet,omitempty"`
	Service                 *[]FirewallPolicy6Service         `json:"service,omitempty"`
	ServiceNegate           *string                           `json:"service-negate,omitempty"`
	SessionTtl              *string                           `json:"session-ttl,omitempty"`
	Srcaddr                 *[]FirewallPolicy6Srcaddr         `json:"srcaddr,omitempty"`
	SrcaddrNegate           *string                           `json:"srcaddr-negate,omitempty"`
	Srcintf                 *[]FirewallPolicy6Srcintf         `json:"srcintf,omitempty"`
	SshFilterProfile        *string                           `json:"ssh-filter-profile,omitempty"`
	SshPolicyRedirect       *string                           `json:"ssh-policy-redirect,omitempty"`
	SslMirror               *string                           `json:"ssl-mirror,omitempty"`
	SslMirrorIntf           *[]FirewallPolicy6SslMirrorIntf   `json:"ssl-mirror-intf,omitempty"`
	SslSshProfile           *string                           `json:"ssl-ssh-profile,omitempty"`
	Status                  *string                           `json:"status,omitempty"`
	TcpMssReceiver          *int64                            `json:"tcp-mss-receiver,omitempty"`
	TcpMssSender            *int64                            `json:"tcp-mss-sender,omitempty"`
	TcpSessionWithoutSyn    *string                           `json:"tcp-session-without-syn,omitempty"`
	TimeoutSendRst          *string                           `json:"timeout-send-rst,omitempty"`
	Tos                     *string                           `json:"tos,omitempty"`
	TosMask                 *string                           `json:"tos-mask,omitempty"`
	TosNegate               *string                           `json:"tos-negate,omitempty"`
	TrafficShaper           *string                           `json:"traffic-shaper,omitempty"`
	TrafficShaperReverse    *string                           `json:"traffic-shaper-reverse,omitempty"`
	UrlCategory             *[]FirewallPolicy6UrlCategory     `json:"url-category,omitempty"`
	Users                   *[]FirewallPolicy6Users           `json:"users,omitempty"`
	UtmStatus               *string                           `json:"utm-status,omitempty"`
	Uuid                    *string                           `json:"uuid,omitempty"`
	VlanCosFwd              *int64                            `json:"vlan-cos-fwd,omitempty"`
	VlanCosRev              *int64                            `json:"vlan-cos-rev,omitempty"`
	VlanFilter              *string                           `json:"vlan-filter,omitempty"`
	VoipProfile             *string                           `json:"voip-profile,omitempty"`
	Vpntunnel               *string                           `json:"vpntunnel,omitempty"`
	WafProfile              *string                           `json:"waf-profile,omitempty"`
	Webcache                *string                           `json:"webcache,omitempty"`
	WebcacheHttps           *string                           `json:"webcache-https,omitempty"`
	WebfilterProfile        *string                           `json:"webfilter-profile,omitempty"`
	WebproxyForwardServer   *string                           `json:"webproxy-forward-server,omitempty"`
	WebproxyProfile         *string                           `json:"webproxy-profile,omitempty"`
}

type FirewallPolicy6AppCategory struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallPolicy6AppGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy6Application struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallPolicy6CustomLogFields struct {
	FieldId *string `json:"field-id,omitempty"`
}

type FirewallPolicy6Dstaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy6Dstintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy6FssoGroups struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy6Groups struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy6Poolname struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy6Service struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy6Srcaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy6Srcintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy6SslMirrorIntf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy6UrlCategory struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallPolicy6Users struct {
	Name *string `json:"name,omitempty"`
}
