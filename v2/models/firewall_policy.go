package models

const FirewallPolicyPath = "firewall/policy/"

type FirewallPolicy struct {
	Action                        *string                                        `json:"action,omitempty"`
	AntiReplay                    *string                                        `json:"anti-replay,omitempty"`
	AppCategory                   *[]FirewallPolicyAppCategory                   `json:"app-category,omitempty"`
	AppGroup                      *[]FirewallPolicyAppGroup                      `json:"app-group,omitempty"`
	Application                   *[]FirewallPolicyApplication                   `json:"application,omitempty"`
	ApplicationList               *string                                        `json:"application-list,omitempty"`
	AuthCert                      *string                                        `json:"auth-cert,omitempty"`
	AuthPath                      *string                                        `json:"auth-path,omitempty"`
	AuthRedirectAddr              *string                                        `json:"auth-redirect-addr,omitempty"`
	AutoAsicOffload               *string                                        `json:"auto-asic-offload,omitempty"`
	AvProfile                     *string                                        `json:"av-profile,omitempty"`
	BlockNotification             *string                                        `json:"block-notification,omitempty"`
	CaptivePortalExempt           *string                                        `json:"captive-portal-exempt,omitempty"`
	CapturePacket                 *string                                        `json:"capture-packet,omitempty"`
	CifsProfile                   *string                                        `json:"cifs-profile,omitempty"`
	Comments                      *string                                        `json:"comments,omitempty"`
	CustomLogFields               *[]FirewallPolicyCustomLogFields               `json:"custom-log-fields,omitempty"`
	DecryptedTrafficMirror        *string                                        `json:"decrypted-traffic-mirror,omitempty"`
	DelayTcpNpuSession            *string                                        `json:"delay-tcp-npu-session,omitempty"`
	DiffservForward               *string                                        `json:"diffserv-forward,omitempty"`
	DiffservReverse               *string                                        `json:"diffserv-reverse,omitempty"`
	DiffservcodeForward           *string                                        `json:"diffservcode-forward,omitempty"`
	DiffservcodeRev               *string                                        `json:"diffservcode-rev,omitempty"`
	Disclaimer                    *string                                        `json:"disclaimer,omitempty"`
	DlpSensor                     *string                                        `json:"dlp-sensor,omitempty"`
	DnsfilterProfile              *string                                        `json:"dnsfilter-profile,omitempty"`
	Dsri                          *string                                        `json:"dsri,omitempty"`
	Dstaddr                       *[]FirewallPolicyDstaddr                       `json:"dstaddr,omitempty"`
	DstaddrNegate                 *string                                        `json:"dstaddr-negate,omitempty"`
	Dstaddr6                      *[]FirewallPolicyDstaddr6                      `json:"dstaddr6,omitempty"`
	Dstintf                       *[]FirewallPolicyDstintf                       `json:"dstintf,omitempty"`
	DynamicShaping                *string                                        `json:"dynamic-shaping,omitempty"`
	EmailCollect                  *string                                        `json:"email-collect,omitempty"`
	EmailfilterProfile            *string                                        `json:"emailfilter-profile,omitempty"`
	Fec                           *string                                        `json:"fec,omitempty"`
	FileFilterProfile             *string                                        `json:"file-filter-profile,omitempty"`
	FirewallSessionDirty          *string                                        `json:"firewall-session-dirty,omitempty"`
	Fixedport                     *string                                        `json:"fixedport,omitempty"`
	Fsso                          *string                                        `json:"fsso,omitempty"`
	FssoAgentForNtlm              *string                                        `json:"fsso-agent-for-ntlm,omitempty"`
	FssoGroups                    *[]FirewallPolicyFssoGroups                    `json:"fsso-groups,omitempty"`
	GeoipAnycast                  *string                                        `json:"geoip-anycast,omitempty"`
	GeoipMatch                    *string                                        `json:"geoip-match,omitempty"`
	Groups                        *[]FirewallPolicyGroups                        `json:"groups,omitempty"`
	HttpPolicyRedirect            *string                                        `json:"http-policy-redirect,omitempty"`
	IcapProfile                   *string                                        `json:"icap-profile,omitempty"`
	IdentityBasedRoute            *string                                        `json:"identity-based-route,omitempty"`
	Inbound                       *string                                        `json:"inbound,omitempty"`
	InspectionMode                *string                                        `json:"inspection-mode,omitempty"`
	InternetService               *string                                        `json:"internet-service,omitempty"`
	InternetServiceCustom         *[]FirewallPolicyInternetServiceCustom         `json:"internet-service-custom,omitempty"`
	InternetServiceCustomGroup    *[]FirewallPolicyInternetServiceCustomGroup    `json:"internet-service-custom-group,omitempty"`
	InternetServiceGroup          *[]FirewallPolicyInternetServiceGroup          `json:"internet-service-group,omitempty"`
	InternetServiceId             *[]FirewallPolicyInternetServiceId             `json:"internet-service-id,omitempty"`
	InternetServiceName           *[]FirewallPolicyInternetServiceName           `json:"internet-service-name,omitempty"`
	InternetServiceNegate         *string                                        `json:"internet-service-negate,omitempty"`
	InternetServiceSrc            *string                                        `json:"internet-service-src,omitempty"`
	InternetServiceSrcCustom      *[]FirewallPolicyInternetServiceSrcCustom      `json:"internet-service-src-custom,omitempty"`
	InternetServiceSrcCustomGroup *[]FirewallPolicyInternetServiceSrcCustomGroup `json:"internet-service-src-custom-group,omitempty"`
	InternetServiceSrcGroup       *[]FirewallPolicyInternetServiceSrcGroup       `json:"internet-service-src-group,omitempty"`
	InternetServiceSrcId          *[]FirewallPolicyInternetServiceSrcId          `json:"internet-service-src-id,omitempty"`
	InternetServiceSrcName        *[]FirewallPolicyInternetServiceSrcName        `json:"internet-service-src-name,omitempty"`
	InternetServiceSrcNegate      *string                                        `json:"internet-service-src-negate,omitempty"`
	Ippool                        *string                                        `json:"ippool,omitempty"`
	IpsSensor                     *string                                        `json:"ips-sensor,omitempty"`
	Logtraffic                    *string                                        `json:"logtraffic,omitempty"`
	LogtrafficStart               *string                                        `json:"logtraffic-start,omitempty"`
	MatchVip                      *string                                        `json:"match-vip,omitempty"`
	MatchVipOnly                  *string                                        `json:"match-vip-only,omitempty"`
	Name                          *string                                        `json:"name,omitempty"`
	Nat                           *string                                        `json:"nat,omitempty"`
	Nat46                         *string                                        `json:"nat46,omitempty"`
	Nat64                         *string                                        `json:"nat64,omitempty"`
	Natinbound                    *string                                        `json:"natinbound,omitempty"`
	Natip                         *string                                        `json:"natip,omitempty"`
	Natoutbound                   *string                                        `json:"natoutbound,omitempty"`
	NpAcceleration                *string                                        `json:"np-acceleration,omitempty"`
	Ntlm                          *string                                        `json:"ntlm,omitempty"`
	NtlmEnabledBrowsers           *[]FirewallPolicyNtlmEnabledBrowsers           `json:"ntlm-enabled-browsers,omitempty"`
	NtlmGuest                     *string                                        `json:"ntlm-guest,omitempty"`
	Outbound                      *string                                        `json:"outbound,omitempty"`
	PassiveWanHealthMeasurement   *string                                        `json:"passive-wan-health-measurement,omitempty"`
	PerIpShaper                   *string                                        `json:"per-ip-shaper,omitempty"`
	PermitAnyHost                 *string                                        `json:"permit-any-host,omitempty"`
	PermitStunHost                *string                                        `json:"permit-stun-host,omitempty"`
	Policyid                      *int64                                         `json:"policyid,omitempty"`
	Poolname                      *[]FirewallPolicyPoolname                      `json:"poolname,omitempty"`
	Poolname6                     *[]FirewallPolicyPoolname6                     `json:"poolname6,omitempty"`
	ProfileGroup                  *string                                        `json:"profile-group,omitempty"`
	ProfileProtocolOptions        *string                                        `json:"profile-protocol-options,omitempty"`
	ProfileType                   *string                                        `json:"profile-type,omitempty"`
	RadiusMacAuthBypass           *string                                        `json:"radius-mac-auth-bypass,omitempty"`
	RedirectUrl                   *string                                        `json:"redirect-url,omitempty"`
	ReplacemsgOverrideGroup       *string                                        `json:"replacemsg-override-group,omitempty"`
	ReputationDirection           *string                                        `json:"reputation-direction,omitempty"`
	ReputationMinimum             *int64                                         `json:"reputation-minimum,omitempty"`
	Rsso                          *string                                        `json:"rsso,omitempty"`
	RtpAddr                       *[]FirewallPolicyRtpAddr                       `json:"rtp-addr,omitempty"`
	RtpNat                        *string                                        `json:"rtp-nat,omitempty"`
	Schedule                      *string                                        `json:"schedule,omitempty"`
	ScheduleTimeout               *string                                        `json:"schedule-timeout,omitempty"`
	SctpFilterProfile             *string                                        `json:"sctp-filter-profile,omitempty"`
	SendDenyPacket                *string                                        `json:"send-deny-packet,omitempty"`
	Service                       *[]FirewallPolicyService                       `json:"service,omitempty"`
	ServiceNegate                 *string                                        `json:"service-negate,omitempty"`
	SessionTtl                    *string                                        `json:"session-ttl,omitempty"`
	Sgt                           *[]FirewallPolicySgt                           `json:"sgt,omitempty"`
	SgtCheck                      *string                                        `json:"sgt-check,omitempty"`
	SrcVendorMac                  *[]FirewallPolicySrcVendorMac                  `json:"src-vendor-mac,omitempty"`
	Srcaddr                       *[]FirewallPolicySrcaddr                       `json:"srcaddr,omitempty"`
	SrcaddrNegate                 *string                                        `json:"srcaddr-negate,omitempty"`
	Srcaddr6                      *[]FirewallPolicySrcaddr6                      `json:"srcaddr6,omitempty"`
	Srcintf                       *[]FirewallPolicySrcintf                       `json:"srcintf,omitempty"`
	SshFilterProfile              *string                                        `json:"ssh-filter-profile,omitempty"`
	SshPolicyRedirect             *string                                        `json:"ssh-policy-redirect,omitempty"`
	SslMirror                     *string                                        `json:"ssl-mirror,omitempty"`
	SslMirrorIntf                 *[]FirewallPolicySslMirrorIntf                 `json:"ssl-mirror-intf,omitempty"`
	SslSshProfile                 *string                                        `json:"ssl-ssh-profile,omitempty"`
	Status                        *string                                        `json:"status,omitempty"`
	TcpMssReceiver                *int64                                         `json:"tcp-mss-receiver,omitempty"`
	TcpMssSender                  *int64                                         `json:"tcp-mss-sender,omitempty"`
	TcpSessionWithoutSyn          *string                                        `json:"tcp-session-without-syn,omitempty"`
	TimeoutSendRst                *string                                        `json:"timeout-send-rst,omitempty"`
	Tos                           *string                                        `json:"tos,omitempty"`
	TosMask                       *string                                        `json:"tos-mask,omitempty"`
	TosNegate                     *string                                        `json:"tos-negate,omitempty"`
	TrafficShaper                 *string                                        `json:"traffic-shaper,omitempty"`
	TrafficShaperReverse          *string                                        `json:"traffic-shaper-reverse,omitempty"`
	UrlCategory                   *[]FirewallPolicyUrlCategory                   `json:"url-category,omitempty"`
	Users                         *[]FirewallPolicyUsers                         `json:"users,omitempty"`
	UtmStatus                     *string                                        `json:"utm-status,omitempty"`
	Uuid                          *string                                        `json:"uuid,omitempty"`
	VideofilterProfile            *string                                        `json:"videofilter-profile,omitempty"`
	VlanCosFwd                    *int64                                         `json:"vlan-cos-fwd,omitempty"`
	VlanCosRev                    *int64                                         `json:"vlan-cos-rev,omitempty"`
	VlanFilter                    *string                                        `json:"vlan-filter,omitempty"`
	VoipProfile                   *string                                        `json:"voip-profile,omitempty"`
	Vpntunnel                     *string                                        `json:"vpntunnel,omitempty"`
	WafProfile                    *string                                        `json:"waf-profile,omitempty"`
	Wanopt                        *string                                        `json:"wanopt,omitempty"`
	WanoptDetection               *string                                        `json:"wanopt-detection,omitempty"`
	WanoptPassiveOpt              *string                                        `json:"wanopt-passive-opt,omitempty"`
	WanoptPeer                    *string                                        `json:"wanopt-peer,omitempty"`
	WanoptProfile                 *string                                        `json:"wanopt-profile,omitempty"`
	Wccp                          *string                                        `json:"wccp,omitempty"`
	Webcache                      *string                                        `json:"webcache,omitempty"`
	WebcacheHttps                 *string                                        `json:"webcache-https,omitempty"`
	WebfilterProfile              *string                                        `json:"webfilter-profile,omitempty"`
	WebproxyForwardServer         *string                                        `json:"webproxy-forward-server,omitempty"`
	WebproxyProfile               *string                                        `json:"webproxy-profile,omitempty"`
	Wsso                          *string                                        `json:"wsso,omitempty"`
	ZtnaEmsTag                    *[]FirewallPolicyZtnaEmsTag                    `json:"ztna-ems-tag,omitempty"`
	ZtnaGeoTag                    *[]FirewallPolicyZtnaGeoTag                    `json:"ztna-geo-tag,omitempty"`
	ZtnaStatus                    *string                                        `json:"ztna-status,omitempty"`
}

type FirewallPolicyAppCategory struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallPolicyAppGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyApplication struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallPolicyCustomLogFields struct {
	FieldId *string `json:"field-id,omitempty"`
}

type FirewallPolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyDstaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyDstintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyFssoGroups struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyGroups struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyInternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyInternetServiceCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyInternetServiceGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyInternetServiceId struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallPolicyInternetServiceName struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyInternetServiceSrcCustom struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyInternetServiceSrcCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyInternetServiceSrcGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyInternetServiceSrcId struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallPolicyInternetServiceSrcName struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyNtlmEnabledBrowsers struct {
	UserAgentString *string `json:"user-agent-string,omitempty"`
}

type FirewallPolicyPoolname struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyPoolname6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyRtpAddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyService struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicySgt struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallPolicySrcVendorMac struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicySrcaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicySrcintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicySslMirrorIntf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyUrlCategory struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallPolicyUsers struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyZtnaEmsTag struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicyZtnaGeoTag struct {
	Name *string `json:"name,omitempty"`
}
