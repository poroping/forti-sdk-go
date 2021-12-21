package models

const FirewallProxyPolicyPath = "firewall/proxy-policy/"

type FirewallProxyPolicy struct {
	AccessProxy                *[]FirewallProxyPolicyAccessProxy                `json:"access-proxy,omitempty"`
	AccessProxy6               *[]FirewallProxyPolicyAccessProxy6               `json:"access-proxy6,omitempty"`
	Action                     *string                                          `json:"action,omitempty"`
	ApplicationList            *string                                          `json:"application-list,omitempty"`
	AvProfile                  *string                                          `json:"av-profile,omitempty"`
	CifsProfile                *string                                          `json:"cifs-profile,omitempty"`
	Comments                   *string                                          `json:"comments,omitempty"`
	DecryptedTrafficMirror     *string                                          `json:"decrypted-traffic-mirror,omitempty"`
	DeviceOwnership            *string                                          `json:"device-ownership,omitempty"`
	Disclaimer                 *string                                          `json:"disclaimer,omitempty"`
	DlpSensor                  *string                                          `json:"dlp-sensor,omitempty"`
	Dstaddr                    *[]FirewallProxyPolicyDstaddr                    `json:"dstaddr,omitempty"`
	DstaddrNegate              *string                                          `json:"dstaddr-negate,omitempty"`
	Dstaddr6                   *[]FirewallProxyPolicyDstaddr6                   `json:"dstaddr6,omitempty"`
	Dstintf                    *[]FirewallProxyPolicyDstintf                    `json:"dstintf,omitempty"`
	EmailfilterProfile         *string                                          `json:"emailfilter-profile,omitempty"`
	FileFilterProfile          *string                                          `json:"file-filter-profile,omitempty"`
	Groups                     *[]FirewallProxyPolicyGroups                     `json:"groups,omitempty"`
	HttpTunnelAuth             *string                                          `json:"http-tunnel-auth,omitempty"`
	IcapProfile                *string                                          `json:"icap-profile,omitempty"`
	InternetService            *string                                          `json:"internet-service,omitempty"`
	InternetServiceCustom      *[]FirewallProxyPolicyInternetServiceCustom      `json:"internet-service-custom,omitempty"`
	InternetServiceCustomGroup *[]FirewallProxyPolicyInternetServiceCustomGroup `json:"internet-service-custom-group,omitempty"`
	InternetServiceGroup       *[]FirewallProxyPolicyInternetServiceGroup       `json:"internet-service-group,omitempty"`
	InternetServiceId          *[]FirewallProxyPolicyInternetServiceId          `json:"internet-service-id,omitempty"`
	InternetServiceName        *[]FirewallProxyPolicyInternetServiceName        `json:"internet-service-name,omitempty"`
	InternetServiceNegate      *string                                          `json:"internet-service-negate,omitempty"`
	IpsSensor                  *string                                          `json:"ips-sensor,omitempty"`
	Logtraffic                 *string                                          `json:"logtraffic,omitempty"`
	LogtrafficStart            *string                                          `json:"logtraffic-start,omitempty"`
	Name                       *string                                          `json:"name,omitempty"`
	Policyid                   *float64                                         `json:"policyid,omitempty"`
	Poolname                   *[]FirewallProxyPolicyPoolname                   `json:"poolname,omitempty"`
	ProfileGroup               *string                                          `json:"profile-group,omitempty"`
	ProfileProtocolOptions     *string                                          `json:"profile-protocol-options,omitempty"`
	ProfileType                *string                                          `json:"profile-type,omitempty"`
	Proxy                      *string                                          `json:"proxy,omitempty"`
	RedirectUrl                *string                                          `json:"redirect-url,omitempty"`
	ReplacemsgOverrideGroup    *string                                          `json:"replacemsg-override-group,omitempty"`
	Schedule                   *string                                          `json:"schedule,omitempty"`
	SctpFilterProfile          *string                                          `json:"sctp-filter-profile,omitempty"`
	Service                    *[]FirewallProxyPolicyService                    `json:"service,omitempty"`
	ServiceNegate              *string                                          `json:"service-negate,omitempty"`
	SessionTtl                 *float64                                         `json:"session-ttl,omitempty"`
	Srcaddr                    *[]FirewallProxyPolicySrcaddr                    `json:"srcaddr,omitempty"`
	SrcaddrNegate              *string                                          `json:"srcaddr-negate,omitempty"`
	Srcaddr6                   *[]FirewallProxyPolicySrcaddr6                   `json:"srcaddr6,omitempty"`
	Srcintf                    *[]FirewallProxyPolicySrcintf                    `json:"srcintf,omitempty"`
	SshFilterProfile           *string                                          `json:"ssh-filter-profile,omitempty"`
	SshPolicyRedirect          *string                                          `json:"ssh-policy-redirect,omitempty"`
	SslSshProfile              *string                                          `json:"ssl-ssh-profile,omitempty"`
	Status                     *string                                          `json:"status,omitempty"`
	Transparent                *string                                          `json:"transparent,omitempty"`
	Users                      *[]FirewallProxyPolicyUsers                      `json:"users,omitempty"`
	UtmStatus                  *string                                          `json:"utm-status,omitempty"`
	Uuid                       *string                                          `json:"uuid,omitempty"`
	VideofilterProfile         *string                                          `json:"videofilter-profile,omitempty"`
	VoipProfile                *string                                          `json:"voip-profile,omitempty"`
	WafProfile                 *string                                          `json:"waf-profile,omitempty"`
	Webcache                   *string                                          `json:"webcache,omitempty"`
	WebcacheHttps              *string                                          `json:"webcache-https,omitempty"`
	WebfilterProfile           *string                                          `json:"webfilter-profile,omitempty"`
	WebproxyForwardServer      *string                                          `json:"webproxy-forward-server,omitempty"`
	WebproxyProfile            *string                                          `json:"webproxy-profile,omitempty"`
	ZtnaEmsTag                 *[]FirewallProxyPolicyZtnaEmsTag                 `json:"ztna-ems-tag,omitempty"`
	ZtnaTagsMatchLogic         *string                                          `json:"ztna-tags-match-logic,omitempty"`
}

type FirewallProxyPolicyAccessProxy struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicyAccessProxy6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicyDstaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicyDstintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicyGroups struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicyInternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicyInternetServiceCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicyInternetServiceGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicyInternetServiceId struct {
	Id *float64 `json:"id,omitempty"`
}

type FirewallProxyPolicyInternetServiceName struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicyPoolname struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicyService struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicySrcaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicySrcintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicyUsers struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyPolicyZtnaEmsTag struct {
	Name *string `json:"name,omitempty"`
}
