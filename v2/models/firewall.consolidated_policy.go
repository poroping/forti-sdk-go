package models

const FirewallConsolidatedPolicyPath = "firewall.consolidated/policy/"

type FirewallConsolidatedPolicy struct {
	Action                        *string                                                    `json:"action,omitempty"`
	ApplicationList               *string                                                    `json:"application-list,omitempty"`
	AutoAsicOffload               *string                                                    `json:"auto-asic-offload,omitempty"`
	AvProfile                     *string                                                    `json:"av-profile,omitempty"`
	CaptivePortalExempt           *string                                                    `json:"captive-portal-exempt,omitempty"`
	CifsProfile                   *string                                                    `json:"cifs-profile,omitempty"`
	Comments                      *string                                                    `json:"comments,omitempty"`
	DiffservForward               *string                                                    `json:"diffserv-forward,omitempty"`
	DiffservReverse               *string                                                    `json:"diffserv-reverse,omitempty"`
	DiffservcodeForward           *string                                                    `json:"diffservcode-forward,omitempty"`
	DiffservcodeRev               *string                                                    `json:"diffservcode-rev,omitempty"`
	DlpSensor                     *string                                                    `json:"dlp-sensor,omitempty"`
	DnsfilterProfile              *string                                                    `json:"dnsfilter-profile,omitempty"`
	DstaddrNegate                 *string                                                    `json:"dstaddr-negate,omitempty"`
	Dstaddr4                      *[]FirewallConsolidatedPolicyDstaddr4                      `json:"dstaddr4,omitempty"`
	Dstaddr6                      *[]FirewallConsolidatedPolicyDstaddr6                      `json:"dstaddr6,omitempty"`
	Dstintf                       *[]FirewallConsolidatedPolicyDstintf                       `json:"dstintf,omitempty"`
	EmailfilterProfile            *string                                                    `json:"emailfilter-profile,omitempty"`
	FileFilterProfile             *string                                                    `json:"file-filter-profile,omitempty"`
	Fixedport                     *string                                                    `json:"fixedport,omitempty"`
	FssoGroups                    *[]FirewallConsolidatedPolicyFssoGroups                    `json:"fsso-groups,omitempty"`
	Groups                        *[]FirewallConsolidatedPolicyGroups                        `json:"groups,omitempty"`
	HttpPolicyRedirect            *string                                                    `json:"http-policy-redirect,omitempty"`
	IcapProfile                   *string                                                    `json:"icap-profile,omitempty"`
	Inbound                       *string                                                    `json:"inbound,omitempty"`
	InspectionMode                *string                                                    `json:"inspection-mode,omitempty"`
	InternetService               *string                                                    `json:"internet-service,omitempty"`
	InternetServiceCustom         *[]FirewallConsolidatedPolicyInternetServiceCustom         `json:"internet-service-custom,omitempty"`
	InternetServiceCustomGroup    *[]FirewallConsolidatedPolicyInternetServiceCustomGroup    `json:"internet-service-custom-group,omitempty"`
	InternetServiceGroup          *[]FirewallConsolidatedPolicyInternetServiceGroup          `json:"internet-service-group,omitempty"`
	InternetServiceId             *[]FirewallConsolidatedPolicyInternetServiceId             `json:"internet-service-id,omitempty"`
	InternetServiceName           *[]FirewallConsolidatedPolicyInternetServiceName           `json:"internet-service-name,omitempty"`
	InternetServiceNegate         *string                                                    `json:"internet-service-negate,omitempty"`
	InternetServiceSrc            *string                                                    `json:"internet-service-src,omitempty"`
	InternetServiceSrcCustom      *[]FirewallConsolidatedPolicyInternetServiceSrcCustom      `json:"internet-service-src-custom,omitempty"`
	InternetServiceSrcCustomGroup *[]FirewallConsolidatedPolicyInternetServiceSrcCustomGroup `json:"internet-service-src-custom-group,omitempty"`
	InternetServiceSrcGroup       *[]FirewallConsolidatedPolicyInternetServiceSrcGroup       `json:"internet-service-src-group,omitempty"`
	InternetServiceSrcId          *[]FirewallConsolidatedPolicyInternetServiceSrcId          `json:"internet-service-src-id,omitempty"`
	InternetServiceSrcName        *[]FirewallConsolidatedPolicyInternetServiceSrcName        `json:"internet-service-src-name,omitempty"`
	InternetServiceSrcNegate      *string                                                    `json:"internet-service-src-negate,omitempty"`
	Ippool                        *string                                                    `json:"ippool,omitempty"`
	IpsSensor                     *string                                                    `json:"ips-sensor,omitempty"`
	Logtraffic                    *string                                                    `json:"logtraffic,omitempty"`
	LogtrafficStart               *string                                                    `json:"logtraffic-start,omitempty"`
	Name                          *string                                                    `json:"name,omitempty"`
	Nat                           *string                                                    `json:"nat,omitempty"`
	Outbound                      *string                                                    `json:"outbound,omitempty"`
	PerIpShaper                   *string                                                    `json:"per-ip-shaper,omitempty"`
	Policyid                      *int64                                                     `json:"policyid,omitempty"`
	Poolname4                     *[]FirewallConsolidatedPolicyPoolname4                     `json:"poolname4,omitempty"`
	Poolname6                     *[]FirewallConsolidatedPolicyPoolname6                     `json:"poolname6,omitempty"`
	ProfileGroup                  *string                                                    `json:"profile-group,omitempty"`
	ProfileProtocolOptions        *string                                                    `json:"profile-protocol-options,omitempty"`
	ProfileType                   *string                                                    `json:"profile-type,omitempty"`
	Schedule                      *string                                                    `json:"schedule,omitempty"`
	Service                       *[]FirewallConsolidatedPolicyService                       `json:"service,omitempty"`
	ServiceNegate                 *string                                                    `json:"service-negate,omitempty"`
	SessionTtl                    *int64                                                     `json:"session-ttl,omitempty"`
	SrcaddrNegate                 *string                                                    `json:"srcaddr-negate,omitempty"`
	Srcaddr4                      *[]FirewallConsolidatedPolicySrcaddr4                      `json:"srcaddr4,omitempty"`
	Srcaddr6                      *[]FirewallConsolidatedPolicySrcaddr6                      `json:"srcaddr6,omitempty"`
	Srcintf                       *[]FirewallConsolidatedPolicySrcintf                       `json:"srcintf,omitempty"`
	SshFilterProfile              *string                                                    `json:"ssh-filter-profile,omitempty"`
	SshPolicyRedirect             *string                                                    `json:"ssh-policy-redirect,omitempty"`
	SslSshProfile                 *string                                                    `json:"ssl-ssh-profile,omitempty"`
	Status                        *string                                                    `json:"status,omitempty"`
	TcpMssReceiver                *int64                                                     `json:"tcp-mss-receiver,omitempty"`
	TcpMssSender                  *int64                                                     `json:"tcp-mss-sender,omitempty"`
	TrafficShaper                 *string                                                    `json:"traffic-shaper,omitempty"`
	TrafficShaperReverse          *string                                                    `json:"traffic-shaper-reverse,omitempty"`
	Users                         *[]FirewallConsolidatedPolicyUsers                         `json:"users,omitempty"`
	UtmStatus                     *string                                                    `json:"utm-status,omitempty"`
	Uuid                          *string                                                    `json:"uuid,omitempty"`
	VoipProfile                   *string                                                    `json:"voip-profile,omitempty"`
	Vpntunnel                     *string                                                    `json:"vpntunnel,omitempty"`
	WafProfile                    *string                                                    `json:"waf-profile,omitempty"`
	Wanopt                        *string                                                    `json:"wanopt,omitempty"`
	WanoptDetection               *string                                                    `json:"wanopt-detection,omitempty"`
	WanoptPassiveOpt              *string                                                    `json:"wanopt-passive-opt,omitempty"`
	WanoptPeer                    *string                                                    `json:"wanopt-peer,omitempty"`
	WanoptProfile                 *string                                                    `json:"wanopt-profile,omitempty"`
	Webcache                      *string                                                    `json:"webcache,omitempty"`
	WebcacheHttps                 *string                                                    `json:"webcache-https,omitempty"`
	WebfilterProfile              *string                                                    `json:"webfilter-profile,omitempty"`
	WebproxyForwardServer         *string                                                    `json:"webproxy-forward-server,omitempty"`
	WebproxyProfile               *string                                                    `json:"webproxy-profile,omitempty"`
}

type FirewallConsolidatedPolicyDstaddr4 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyDstaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyDstintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyFssoGroups struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyGroups struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyInternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyInternetServiceCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyInternetServiceGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyInternetServiceId struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallConsolidatedPolicyInternetServiceName struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyInternetServiceSrcCustom struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyInternetServiceSrcCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyInternetServiceSrcGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyInternetServiceSrcId struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallConsolidatedPolicyInternetServiceSrcName struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyPoolname4 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyPoolname6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyService struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicySrcaddr4 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicySrcaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicySrcintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallConsolidatedPolicyUsers struct {
	Name *string `json:"name,omitempty"`
}
