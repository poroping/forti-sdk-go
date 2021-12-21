package models

const FirewallconsolidatedPolicyPath = "firewall.consolidated/policy/"

type FirewallconsolidatedPolicy struct {
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
	Dstaddr4                      *[]FirewallconsolidatedPolicyDstaddr4                      `json:"dstaddr4,omitempty"`
	Dstaddr6                      *[]FirewallconsolidatedPolicyDstaddr6                      `json:"dstaddr6,omitempty"`
	Dstintf                       *[]FirewallconsolidatedPolicyDstintf                       `json:"dstintf,omitempty"`
	EmailfilterProfile            *string                                                    `json:"emailfilter-profile,omitempty"`
	FileFilterProfile             *string                                                    `json:"file-filter-profile,omitempty"`
	Fixedport                     *string                                                    `json:"fixedport,omitempty"`
	FssoGroups                    *[]FirewallconsolidatedPolicyFssoGroups                    `json:"fsso-groups,omitempty"`
	Groups                        *[]FirewallconsolidatedPolicyGroups                        `json:"groups,omitempty"`
	HttpPolicyRedirect            *string                                                    `json:"http-policy-redirect,omitempty"`
	IcapProfile                   *string                                                    `json:"icap-profile,omitempty"`
	Inbound                       *string                                                    `json:"inbound,omitempty"`
	InspectionMode                *string                                                    `json:"inspection-mode,omitempty"`
	InternetService               *string                                                    `json:"internet-service,omitempty"`
	InternetServiceCustom         *[]FirewallconsolidatedPolicyInternetServiceCustom         `json:"internet-service-custom,omitempty"`
	InternetServiceCustomGroup    *[]FirewallconsolidatedPolicyInternetServiceCustomGroup    `json:"internet-service-custom-group,omitempty"`
	InternetServiceGroup          *[]FirewallconsolidatedPolicyInternetServiceGroup          `json:"internet-service-group,omitempty"`
	InternetServiceId             *[]FirewallconsolidatedPolicyInternetServiceId             `json:"internet-service-id,omitempty"`
	InternetServiceName           *[]FirewallconsolidatedPolicyInternetServiceName           `json:"internet-service-name,omitempty"`
	InternetServiceNegate         *string                                                    `json:"internet-service-negate,omitempty"`
	InternetServiceSrc            *string                                                    `json:"internet-service-src,omitempty"`
	InternetServiceSrcCustom      *[]FirewallconsolidatedPolicyInternetServiceSrcCustom      `json:"internet-service-src-custom,omitempty"`
	InternetServiceSrcCustomGroup *[]FirewallconsolidatedPolicyInternetServiceSrcCustomGroup `json:"internet-service-src-custom-group,omitempty"`
	InternetServiceSrcGroup       *[]FirewallconsolidatedPolicyInternetServiceSrcGroup       `json:"internet-service-src-group,omitempty"`
	InternetServiceSrcId          *[]FirewallconsolidatedPolicyInternetServiceSrcId          `json:"internet-service-src-id,omitempty"`
	InternetServiceSrcName        *[]FirewallconsolidatedPolicyInternetServiceSrcName        `json:"internet-service-src-name,omitempty"`
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
	Poolname4                     *[]FirewallconsolidatedPolicyPoolname4                     `json:"poolname4,omitempty"`
	Poolname6                     *[]FirewallconsolidatedPolicyPoolname6                     `json:"poolname6,omitempty"`
	ProfileGroup                  *string                                                    `json:"profile-group,omitempty"`
	ProfileProtocolOptions        *string                                                    `json:"profile-protocol-options,omitempty"`
	ProfileType                   *string                                                    `json:"profile-type,omitempty"`
	Schedule                      *string                                                    `json:"schedule,omitempty"`
	Service                       *[]FirewallconsolidatedPolicyService                       `json:"service,omitempty"`
	ServiceNegate                 *string                                                    `json:"service-negate,omitempty"`
	SessionTtl                    *int64                                                     `json:"session-ttl,omitempty"`
	SrcaddrNegate                 *string                                                    `json:"srcaddr-negate,omitempty"`
	Srcaddr4                      *[]FirewallconsolidatedPolicySrcaddr4                      `json:"srcaddr4,omitempty"`
	Srcaddr6                      *[]FirewallconsolidatedPolicySrcaddr6                      `json:"srcaddr6,omitempty"`
	Srcintf                       *[]FirewallconsolidatedPolicySrcintf                       `json:"srcintf,omitempty"`
	SshFilterProfile              *string                                                    `json:"ssh-filter-profile,omitempty"`
	SshPolicyRedirect             *string                                                    `json:"ssh-policy-redirect,omitempty"`
	SslSshProfile                 *string                                                    `json:"ssl-ssh-profile,omitempty"`
	Status                        *string                                                    `json:"status,omitempty"`
	TcpMssReceiver                *int64                                                     `json:"tcp-mss-receiver,omitempty"`
	TcpMssSender                  *int64                                                     `json:"tcp-mss-sender,omitempty"`
	TrafficShaper                 *string                                                    `json:"traffic-shaper,omitempty"`
	TrafficShaperReverse          *string                                                    `json:"traffic-shaper-reverse,omitempty"`
	Users                         *[]FirewallconsolidatedPolicyUsers                         `json:"users,omitempty"`
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

type FirewallconsolidatedPolicyDstaddr4 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyDstaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyDstintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyFssoGroups struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyGroups struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyInternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyInternetServiceCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyInternetServiceGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyInternetServiceId struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallconsolidatedPolicyInternetServiceName struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyInternetServiceSrcCustom struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyInternetServiceSrcCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyInternetServiceSrcGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyInternetServiceSrcId struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallconsolidatedPolicyInternetServiceSrcName struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyPoolname4 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyPoolname6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyService struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicySrcaddr4 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicySrcaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicySrcintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallconsolidatedPolicyUsers struct {
	Name *string `json:"name,omitempty"`
}
