package models

const FirewallSecurityPolicyPath = "firewall/security-policy/"

type FirewallSecurityPolicy struct {
	Action                         *string                                                 `json:"action,omitempty"`
	AppCategory                    *[]FirewallSecurityPolicyAppCategory                    `json:"app-category,omitempty"`
	AppGroup                       *[]FirewallSecurityPolicyAppGroup                       `json:"app-group,omitempty"`
	Application                    *[]FirewallSecurityPolicyApplication                    `json:"application,omitempty"`
	ApplicationList                *string                                                 `json:"application-list,omitempty"`
	AvProfile                      *string                                                 `json:"av-profile,omitempty"`
	CifsProfile                    *string                                                 `json:"cifs-profile,omitempty"`
	Comments                       *string                                                 `json:"comments,omitempty"`
	DlpProfile                     *string                                                 `json:"dlp-profile,omitempty"`
	DlpSensor                      *string                                                 `json:"dlp-sensor,omitempty"`
	DnsfilterProfile               *string                                                 `json:"dnsfilter-profile,omitempty"`
	Dstaddr                        *[]FirewallSecurityPolicyDstaddr                        `json:"dstaddr,omitempty"`
	DstaddrNegate                  *string                                                 `json:"dstaddr-negate,omitempty"`
	Dstaddr4                       *[]FirewallSecurityPolicyDstaddr4                       `json:"dstaddr4,omitempty"`
	Dstaddr6                       *[]FirewallSecurityPolicyDstaddr6                       `json:"dstaddr6,omitempty"`
	Dstaddr6Negate                 *string                                                 `json:"dstaddr6-negate,omitempty"`
	Dstintf                        *[]FirewallSecurityPolicyDstintf                        `json:"dstintf,omitempty"`
	EmailfilterProfile             *string                                                 `json:"emailfilter-profile,omitempty"`
	EnforceDefaultAppPort          *string                                                 `json:"enforce-default-app-port,omitempty"`
	FileFilterProfile              *string                                                 `json:"file-filter-profile,omitempty"`
	FssoGroups                     *[]FirewallSecurityPolicyFssoGroups                     `json:"fsso-groups,omitempty"`
	Groups                         *[]FirewallSecurityPolicyGroups                         `json:"groups,omitempty"`
	IcapProfile                    *string                                                 `json:"icap-profile,omitempty"`
	InternetService                *string                                                 `json:"internet-service,omitempty"`
	InternetServiceCustom          *[]FirewallSecurityPolicyInternetServiceCustom          `json:"internet-service-custom,omitempty"`
	InternetServiceCustomGroup     *[]FirewallSecurityPolicyInternetServiceCustomGroup     `json:"internet-service-custom-group,omitempty"`
	InternetServiceGroup           *[]FirewallSecurityPolicyInternetServiceGroup           `json:"internet-service-group,omitempty"`
	InternetServiceId              *[]FirewallSecurityPolicyInternetServiceId              `json:"internet-service-id,omitempty"`
	InternetServiceName            *[]FirewallSecurityPolicyInternetServiceName            `json:"internet-service-name,omitempty"`
	InternetServiceNegate          *string                                                 `json:"internet-service-negate,omitempty"`
	InternetServiceSrc             *string                                                 `json:"internet-service-src,omitempty"`
	InternetServiceSrcCustom       *[]FirewallSecurityPolicyInternetServiceSrcCustom       `json:"internet-service-src-custom,omitempty"`
	InternetServiceSrcCustomGroup  *[]FirewallSecurityPolicyInternetServiceSrcCustomGroup  `json:"internet-service-src-custom-group,omitempty"`
	InternetServiceSrcGroup        *[]FirewallSecurityPolicyInternetServiceSrcGroup        `json:"internet-service-src-group,omitempty"`
	InternetServiceSrcId           *[]FirewallSecurityPolicyInternetServiceSrcId           `json:"internet-service-src-id,omitempty"`
	InternetServiceSrcName         *[]FirewallSecurityPolicyInternetServiceSrcName         `json:"internet-service-src-name,omitempty"`
	InternetServiceSrcNegate       *string                                                 `json:"internet-service-src-negate,omitempty"`
	InternetService6               *string                                                 `json:"internet-service6,omitempty"`
	InternetService6Custom         *[]FirewallSecurityPolicyInternetService6Custom         `json:"internet-service6-custom,omitempty"`
	InternetService6CustomGroup    *[]FirewallSecurityPolicyInternetService6CustomGroup    `json:"internet-service6-custom-group,omitempty"`
	InternetService6Group          *[]FirewallSecurityPolicyInternetService6Group          `json:"internet-service6-group,omitempty"`
	InternetService6Name           *[]FirewallSecurityPolicyInternetService6Name           `json:"internet-service6-name,omitempty"`
	InternetService6Negate         *string                                                 `json:"internet-service6-negate,omitempty"`
	InternetService6Src            *string                                                 `json:"internet-service6-src,omitempty"`
	InternetService6SrcCustom      *[]FirewallSecurityPolicyInternetService6SrcCustom      `json:"internet-service6-src-custom,omitempty"`
	InternetService6SrcCustomGroup *[]FirewallSecurityPolicyInternetService6SrcCustomGroup `json:"internet-service6-src-custom-group,omitempty"`
	InternetService6SrcGroup       *[]FirewallSecurityPolicyInternetService6SrcGroup       `json:"internet-service6-src-group,omitempty"`
	InternetService6SrcName        *[]FirewallSecurityPolicyInternetService6SrcName        `json:"internet-service6-src-name,omitempty"`
	InternetService6SrcNegate      *string                                                 `json:"internet-service6-src-negate,omitempty"`
	IpsSensor                      *string                                                 `json:"ips-sensor,omitempty"`
	IpsVoipFilter                  *string                                                 `json:"ips-voip-filter,omitempty"`
	LearningMode                   *string                                                 `json:"learning-mode,omitempty"`
	Logtraffic                     *string                                                 `json:"logtraffic,omitempty"`
	LogtrafficStart                *string                                                 `json:"logtraffic-start,omitempty"`
	Name                           *string                                                 `json:"name,omitempty"`
	Nat46                          *string                                                 `json:"nat46,omitempty"`
	Nat64                          *string                                                 `json:"nat64,omitempty"`
	Policyid                       *int64                                                  `json:"policyid,omitempty"`
	ProfileGroup                   *string                                                 `json:"profile-group,omitempty"`
	ProfileProtocolOptions         *string                                                 `json:"profile-protocol-options,omitempty"`
	ProfileType                    *string                                                 `json:"profile-type,omitempty"`
	Schedule                       *string                                                 `json:"schedule,omitempty"`
	SctpFilterProfile              *string                                                 `json:"sctp-filter-profile,omitempty"`
	SendDenyPacket                 *string                                                 `json:"send-deny-packet,omitempty"`
	Service                        *[]FirewallSecurityPolicyService                        `json:"service,omitempty"`
	ServiceNegate                  *string                                                 `json:"service-negate,omitempty"`
	Srcaddr                        *[]FirewallSecurityPolicySrcaddr                        `json:"srcaddr,omitempty"`
	SrcaddrNegate                  *string                                                 `json:"srcaddr-negate,omitempty"`
	Srcaddr4                       *[]FirewallSecurityPolicySrcaddr4                       `json:"srcaddr4,omitempty"`
	Srcaddr6                       *[]FirewallSecurityPolicySrcaddr6                       `json:"srcaddr6,omitempty"`
	Srcaddr6Negate                 *string                                                 `json:"srcaddr6-negate,omitempty"`
	Srcintf                        *[]FirewallSecurityPolicySrcintf                        `json:"srcintf,omitempty"`
	SshFilterProfile               *string                                                 `json:"ssh-filter-profile,omitempty"`
	SslSshProfile                  *string                                                 `json:"ssl-ssh-profile,omitempty"`
	Status                         *string                                                 `json:"status,omitempty"`
	UrlCategory                    *string                                                 `json:"url-category,omitempty"`
	Users                          *[]FirewallSecurityPolicyUsers                          `json:"users,omitempty"`
	Uuid                           *string                                                 `json:"uuid,omitempty"`
	VideofilterProfile             *string                                                 `json:"videofilter-profile,omitempty"`
	VoipProfile                    *string                                                 `json:"voip-profile,omitempty"`
	WebfilterProfile               *string                                                 `json:"webfilter-profile,omitempty"`
}

type FirewallSecurityPolicyAppCategory struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallSecurityPolicyAppGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyApplication struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallSecurityPolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyDstaddr4 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyDstaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyDstintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyFssoGroups struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyGroups struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetServiceCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetServiceGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetServiceId struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallSecurityPolicyInternetServiceName struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetServiceSrcCustom struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetServiceSrcCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetServiceSrcGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetServiceSrcId struct {
	Id *int64 `json:"id,omitempty"`
}

type FirewallSecurityPolicyInternetServiceSrcName struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetService6Custom struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetService6CustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetService6Group struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetService6Name struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetService6SrcCustom struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetService6SrcCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetService6SrcGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyInternetService6SrcName struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyService struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicySrcaddr4 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicySrcaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicySrcintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSecurityPolicyUsers struct {
	Name *string `json:"name,omitempty"`
}
