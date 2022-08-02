package models

const FirewallShapingPolicyPath = "firewall/shaping-policy/"

type FirewallShapingPolicy struct {
	AppCategory                   *[]FirewallShapingPolicyAppCategory                   `json:"app-category,omitempty"`
	AppGroup                      *[]FirewallShapingPolicyAppGroup                      `json:"app-group,omitempty"`
	Application                   *[]FirewallShapingPolicyApplication                   `json:"application,omitempty"`
	ClassId                       *int64                                                `json:"class-id,omitempty"`
	Comment                       *string                                               `json:"comment,omitempty"`
	DiffservForward               *string                                               `json:"diffserv-forward,omitempty"`
	DiffservReverse               *string                                               `json:"diffserv-reverse,omitempty"`
	DiffservcodeForward           *string                                               `json:"diffservcode-forward,omitempty"`
	DiffservcodeRev               *string                                               `json:"diffservcode-rev,omitempty"`
	Dstaddr                       *[]FirewallShapingPolicyDstaddr                       `json:"dstaddr,omitempty"`
	Dstaddr6                      *[]FirewallShapingPolicyDstaddr6                      `json:"dstaddr6,omitempty"`
	Dstintf                       *[]FirewallShapingPolicyDstintf                       `json:"dstintf,omitempty"`
	Groups                        *[]FirewallShapingPolicyGroups                        `json:"groups,omitempty"`
	Id                            *int64                                                `json:"id,omitempty"`
	InternetService               *string                                               `json:"internet-service,omitempty"`
	InternetServiceCustom         *[]FirewallShapingPolicyInternetServiceCustom         `json:"internet-service-custom,omitempty"`
	InternetServiceCustomGroup    *[]FirewallShapingPolicyInternetServiceCustomGroup    `json:"internet-service-custom-group,omitempty"`
	InternetServiceGroup          *[]FirewallShapingPolicyInternetServiceGroup          `json:"internet-service-group,omitempty"`
	InternetServiceName           *[]FirewallShapingPolicyInternetServiceName           `json:"internet-service-name,omitempty"`
	InternetServiceSrc            *string                                               `json:"internet-service-src,omitempty"`
	InternetServiceSrcCustom      *[]FirewallShapingPolicyInternetServiceSrcCustom      `json:"internet-service-src-custom,omitempty"`
	InternetServiceSrcCustomGroup *[]FirewallShapingPolicyInternetServiceSrcCustomGroup `json:"internet-service-src-custom-group,omitempty"`
	InternetServiceSrcGroup       *[]FirewallShapingPolicyInternetServiceSrcGroup       `json:"internet-service-src-group,omitempty"`
	InternetServiceSrcName        *[]FirewallShapingPolicyInternetServiceSrcName        `json:"internet-service-src-name,omitempty"`
	IpVersion                     *string                                               `json:"ip-version,omitempty"`
	Name                          *string                                               `json:"name,omitempty"`
	PerIpShaper                   *string                                               `json:"per-ip-shaper,omitempty"`
	Schedule                      *string                                               `json:"schedule,omitempty"`
	Service                       *[]FirewallShapingPolicyService                       `json:"service,omitempty"`
	Srcaddr                       *[]FirewallShapingPolicySrcaddr                       `json:"srcaddr,omitempty"`
	Srcaddr6                      *[]FirewallShapingPolicySrcaddr6                      `json:"srcaddr6,omitempty"`
	Srcintf                       *[]FirewallShapingPolicySrcintf                       `json:"srcintf,omitempty"`
	Status                        *string                                               `json:"status,omitempty"`
	Tos                           *string                                               `json:"tos,omitempty"`
	TosMask                       *string                                               `json:"tos-mask,omitempty"`
	TosNegate                     *string                                               `json:"tos-negate,omitempty"`
	TrafficShaper                 *string                                               `json:"traffic-shaper,omitempty"`
	TrafficShaperReverse          *string                                               `json:"traffic-shaper-reverse,omitempty"`
	UrlCategory                   *[]FirewallShapingPolicyUrlCategory                   `json:"url-category,omitempty"`
	Users                         *[]FirewallShapingPolicyUsers                         `json:"users,omitempty"`
}

const FirewallShapingPolicyAppCategoryPath = "firewall/shaping-policy/app-category/"

type FirewallShapingPolicyAppCategory struct {
	Id *int64 `json:"id,omitempty"`
}

const FirewallShapingPolicyAppGroupPath = "firewall/shaping-policy/app-group/"

type FirewallShapingPolicyAppGroup struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyApplicationPath = "firewall/shaping-policy/application/"

type FirewallShapingPolicyApplication struct {
	Id *int64 `json:"id,omitempty"`
}

const FirewallShapingPolicyDstaddrPath = "firewall/shaping-policy/dstaddr/"

type FirewallShapingPolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyDstaddr6Path = "firewall/shaping-policy/dstaddr6/"

type FirewallShapingPolicyDstaddr6 struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyDstintfPath = "firewall/shaping-policy/dstintf/"

type FirewallShapingPolicyDstintf struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyGroupsPath = "firewall/shaping-policy/groups/"

type FirewallShapingPolicyGroups struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyInternetServiceCustomPath = "firewall/shaping-policy/internet-service-custom/"

type FirewallShapingPolicyInternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyInternetServiceCustomGroupPath = "firewall/shaping-policy/internet-service-custom-group/"

type FirewallShapingPolicyInternetServiceCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyInternetServiceGroupPath = "firewall/shaping-policy/internet-service-group/"

type FirewallShapingPolicyInternetServiceGroup struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyInternetServiceNamePath = "firewall/shaping-policy/internet-service-name/"

type FirewallShapingPolicyInternetServiceName struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyInternetServiceSrcCustomPath = "firewall/shaping-policy/internet-service-src-custom/"

type FirewallShapingPolicyInternetServiceSrcCustom struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyInternetServiceSrcCustomGroupPath = "firewall/shaping-policy/internet-service-src-custom-group/"

type FirewallShapingPolicyInternetServiceSrcCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyInternetServiceSrcGroupPath = "firewall/shaping-policy/internet-service-src-group/"

type FirewallShapingPolicyInternetServiceSrcGroup struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyInternetServiceSrcNamePath = "firewall/shaping-policy/internet-service-src-name/"

type FirewallShapingPolicyInternetServiceSrcName struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyServicePath = "firewall/shaping-policy/service/"

type FirewallShapingPolicyService struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicySrcaddrPath = "firewall/shaping-policy/srcaddr/"

type FirewallShapingPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicySrcaddr6Path = "firewall/shaping-policy/srcaddr6/"

type FirewallShapingPolicySrcaddr6 struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicySrcintfPath = "firewall/shaping-policy/srcintf/"

type FirewallShapingPolicySrcintf struct {
	Name *string `json:"name,omitempty"`
}

const FirewallShapingPolicyUrlCategoryPath = "firewall/shaping-policy/url-category/"

type FirewallShapingPolicyUrlCategory struct {
	Id *int64 `json:"id,omitempty"`
}

const FirewallShapingPolicyUsersPath = "firewall/shaping-policy/users/"

type FirewallShapingPolicyUsers struct {
	Name *string `json:"name,omitempty"`
}
