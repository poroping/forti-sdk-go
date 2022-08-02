package models

const FirewallSecurityPolicyPath = "firewall/security-policy/"

type FirewallSecurityPolicy struct {
	Action                        *string                                                `json:"action,omitempty"`
	AppCategory                   *[]FirewallSecurityPolicyAppCategory                   `json:"app-category,omitempty"`
	AppGroup                      *[]FirewallSecurityPolicyAppGroup                      `json:"app-group,omitempty"`
	Application                   *[]FirewallSecurityPolicyApplication                   `json:"application,omitempty"`
	ApplicationList               *string                                                `json:"application-list,omitempty"`
	AvProfile                     *string                                                `json:"av-profile,omitempty"`
	CifsProfile                   *string                                                `json:"cifs-profile,omitempty"`
	Comments                      *string                                                `json:"comments,omitempty"`
	DlpSensor                     *string                                                `json:"dlp-sensor,omitempty"`
	DnsfilterProfile              *string                                                `json:"dnsfilter-profile,omitempty"`
	Dstaddr                       *[]FirewallSecurityPolicyDstaddr                       `json:"dstaddr,omitempty"`
	DstaddrNegate                 *string                                                `json:"dstaddr-negate,omitempty"`
	Dstaddr6                      *[]FirewallSecurityPolicyDstaddr6                      `json:"dstaddr6,omitempty"`
	Dstintf                       *[]FirewallSecurityPolicyDstintf                       `json:"dstintf,omitempty"`
	EmailfilterProfile            *string                                                `json:"emailfilter-profile,omitempty"`
	EnforceDefaultAppPort         *string                                                `json:"enforce-default-app-port,omitempty"`
	FileFilterProfile             *string                                                `json:"file-filter-profile,omitempty"`
	FssoGroups                    *[]FirewallSecurityPolicyFssoGroups                    `json:"fsso-groups,omitempty"`
	Groups                        *[]FirewallSecurityPolicyGroups                        `json:"groups,omitempty"`
	IcapProfile                   *string                                                `json:"icap-profile,omitempty"`
	InternetService               *string                                                `json:"internet-service,omitempty"`
	InternetServiceCustom         *[]FirewallSecurityPolicyInternetServiceCustom         `json:"internet-service-custom,omitempty"`
	InternetServiceCustomGroup    *[]FirewallSecurityPolicyInternetServiceCustomGroup    `json:"internet-service-custom-group,omitempty"`
	InternetServiceGroup          *[]FirewallSecurityPolicyInternetServiceGroup          `json:"internet-service-group,omitempty"`
	InternetServiceName           *[]FirewallSecurityPolicyInternetServiceName           `json:"internet-service-name,omitempty"`
	InternetServiceNegate         *string                                                `json:"internet-service-negate,omitempty"`
	InternetServiceSrc            *string                                                `json:"internet-service-src,omitempty"`
	InternetServiceSrcCustom      *[]FirewallSecurityPolicyInternetServiceSrcCustom      `json:"internet-service-src-custom,omitempty"`
	InternetServiceSrcCustomGroup *[]FirewallSecurityPolicyInternetServiceSrcCustomGroup `json:"internet-service-src-custom-group,omitempty"`
	InternetServiceSrcGroup       *[]FirewallSecurityPolicyInternetServiceSrcGroup       `json:"internet-service-src-group,omitempty"`
	InternetServiceSrcName        *[]FirewallSecurityPolicyInternetServiceSrcName        `json:"internet-service-src-name,omitempty"`
	InternetServiceSrcNegate      *string                                                `json:"internet-service-src-negate,omitempty"`
	IpsSensor                     *string                                                `json:"ips-sensor,omitempty"`
	LearningMode                  *string                                                `json:"learning-mode,omitempty"`
	Logtraffic                    *string                                                `json:"logtraffic,omitempty"`
	Name                          *string                                                `json:"name,omitempty"`
	Nat46                         *string                                                `json:"nat46,omitempty"`
	Nat64                         *string                                                `json:"nat64,omitempty"`
	Policyid                      *int64                                                 `json:"policyid,omitempty"`
	ProfileGroup                  *string                                                `json:"profile-group,omitempty"`
	ProfileProtocolOptions        *string                                                `json:"profile-protocol-options,omitempty"`
	ProfileType                   *string                                                `json:"profile-type,omitempty"`
	Schedule                      *string                                                `json:"schedule,omitempty"`
	SctpFilterProfile             *string                                                `json:"sctp-filter-profile,omitempty"`
	SendDenyPacket                *string                                                `json:"send-deny-packet,omitempty"`
	Service                       *[]FirewallSecurityPolicyService                       `json:"service,omitempty"`
	ServiceNegate                 *string                                                `json:"service-negate,omitempty"`
	Srcaddr                       *[]FirewallSecurityPolicySrcaddr                       `json:"srcaddr,omitempty"`
	SrcaddrNegate                 *string                                                `json:"srcaddr-negate,omitempty"`
	Srcaddr6                      *[]FirewallSecurityPolicySrcaddr6                      `json:"srcaddr6,omitempty"`
	Srcintf                       *[]FirewallSecurityPolicySrcintf                       `json:"srcintf,omitempty"`
	SshFilterProfile              *string                                                `json:"ssh-filter-profile,omitempty"`
	SslSshProfile                 *string                                                `json:"ssl-ssh-profile,omitempty"`
	Status                        *string                                                `json:"status,omitempty"`
	UrlCategory                   *[]FirewallSecurityPolicyUrlCategory                   `json:"url-category,omitempty"`
	Users                         *[]FirewallSecurityPolicyUsers                         `json:"users,omitempty"`
	Uuid                          *string                                                `json:"uuid,omitempty"`
	VideofilterProfile            *string                                                `json:"videofilter-profile,omitempty"`
	VoipProfile                   *string                                                `json:"voip-profile,omitempty"`
	WebfilterProfile              *string                                                `json:"webfilter-profile,omitempty"`
}

const FirewallSecurityPolicyAppCategoryPath = "firewall/security-policy/app-category/"

type FirewallSecurityPolicyAppCategory struct {
	Id *int64 `json:"id,omitempty"`
}

const FirewallSecurityPolicyAppGroupPath = "firewall/security-policy/app-group/"

type FirewallSecurityPolicyAppGroup struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyApplicationPath = "firewall/security-policy/application/"

type FirewallSecurityPolicyApplication struct {
	Id *int64 `json:"id,omitempty"`
}

const FirewallSecurityPolicyDstaddrPath = "firewall/security-policy/dstaddr/"

type FirewallSecurityPolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyDstaddr6Path = "firewall/security-policy/dstaddr6/"

type FirewallSecurityPolicyDstaddr6 struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyDstintfPath = "firewall/security-policy/dstintf/"

type FirewallSecurityPolicyDstintf struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyFssoGroupsPath = "firewall/security-policy/fsso-groups/"

type FirewallSecurityPolicyFssoGroups struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyGroupsPath = "firewall/security-policy/groups/"

type FirewallSecurityPolicyGroups struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyInternetServiceCustomPath = "firewall/security-policy/internet-service-custom/"

type FirewallSecurityPolicyInternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyInternetServiceCustomGroupPath = "firewall/security-policy/internet-service-custom-group/"

type FirewallSecurityPolicyInternetServiceCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyInternetServiceGroupPath = "firewall/security-policy/internet-service-group/"

type FirewallSecurityPolicyInternetServiceGroup struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyInternetServiceNamePath = "firewall/security-policy/internet-service-name/"

type FirewallSecurityPolicyInternetServiceName struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyInternetServiceSrcCustomPath = "firewall/security-policy/internet-service-src-custom/"

type FirewallSecurityPolicyInternetServiceSrcCustom struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyInternetServiceSrcCustomGroupPath = "firewall/security-policy/internet-service-src-custom-group/"

type FirewallSecurityPolicyInternetServiceSrcCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyInternetServiceSrcGroupPath = "firewall/security-policy/internet-service-src-group/"

type FirewallSecurityPolicyInternetServiceSrcGroup struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyInternetServiceSrcNamePath = "firewall/security-policy/internet-service-src-name/"

type FirewallSecurityPolicyInternetServiceSrcName struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyServicePath = "firewall/security-policy/service/"

type FirewallSecurityPolicyService struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicySrcaddrPath = "firewall/security-policy/srcaddr/"

type FirewallSecurityPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicySrcaddr6Path = "firewall/security-policy/srcaddr6/"

type FirewallSecurityPolicySrcaddr6 struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicySrcintfPath = "firewall/security-policy/srcintf/"

type FirewallSecurityPolicySrcintf struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSecurityPolicyUrlCategoryPath = "firewall/security-policy/url-category/"

type FirewallSecurityPolicyUrlCategory struct {
	Id *int64 `json:"id,omitempty"`
}

const FirewallSecurityPolicyUsersPath = "firewall/security-policy/users/"

type FirewallSecurityPolicyUsers struct {
	Name *string `json:"name,omitempty"`
}
