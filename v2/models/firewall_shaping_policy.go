package models

const FirewallShapingPolicyPath = "firewall/shaping-policy/"

type FirewallShapingPolicy struct {
	AppCategory                   *[]FirewallShapingPolicyAppCategory                   `json:"app-category,omitempty"`
	AppGroup                      *[]FirewallShapingPolicyAppGroup                      `json:"app-group,omitempty"`
	Application                   *[]FirewallShapingPolicyApplication                   `json:"application,omitempty"`
	ClassId                       *float64                                              `json:"class-id,omitempty"`
	Comment                       *string                                               `json:"comment,omitempty"`
	DiffservForward               *string                                               `json:"diffserv-forward,omitempty"`
	DiffservReverse               *string                                               `json:"diffserv-reverse,omitempty"`
	DiffservcodeForward           *string                                               `json:"diffservcode-forward,omitempty"`
	DiffservcodeRev               *string                                               `json:"diffservcode-rev,omitempty"`
	Dstaddr                       *[]FirewallShapingPolicyDstaddr                       `json:"dstaddr,omitempty"`
	Dstaddr6                      *[]FirewallShapingPolicyDstaddr6                      `json:"dstaddr6,omitempty"`
	Dstintf                       *[]FirewallShapingPolicyDstintf                       `json:"dstintf,omitempty"`
	Groups                        *[]FirewallShapingPolicyGroups                        `json:"groups,omitempty"`
	Fosid                         *float64                                              `json:"fosid,omitempty"`
	InternetService               *string                                               `json:"internet-service,omitempty"`
	InternetServiceCustom         *[]FirewallShapingPolicyInternetServiceCustom         `json:"internet-service-custom,omitempty"`
	InternetServiceCustomGroup    *[]FirewallShapingPolicyInternetServiceCustomGroup    `json:"internet-service-custom-group,omitempty"`
	InternetServiceGroup          *[]FirewallShapingPolicyInternetServiceGroup          `json:"internet-service-group,omitempty"`
	InternetServiceId             *[]FirewallShapingPolicyInternetServiceId             `json:"internet-service-id,omitempty"`
	InternetServiceName           *[]FirewallShapingPolicyInternetServiceName           `json:"internet-service-name,omitempty"`
	InternetServiceSrc            *string                                               `json:"internet-service-src,omitempty"`
	InternetServiceSrcCustom      *[]FirewallShapingPolicyInternetServiceSrcCustom      `json:"internet-service-src-custom,omitempty"`
	InternetServiceSrcCustomGroup *[]FirewallShapingPolicyInternetServiceSrcCustomGroup `json:"internet-service-src-custom-group,omitempty"`
	InternetServiceSrcGroup       *[]FirewallShapingPolicyInternetServiceSrcGroup       `json:"internet-service-src-group,omitempty"`
	InternetServiceSrcId          *[]FirewallShapingPolicyInternetServiceSrcId          `json:"internet-service-src-id,omitempty"`
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

type FirewallShapingPolicyAppCategory struct {
	Id *float64 `json:"id,omitempty"`
}

type FirewallShapingPolicyAppGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyApplication struct {
	Id *float64 `json:"id,omitempty"`
}

type FirewallShapingPolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyDstaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyDstintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyGroups struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyInternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyInternetServiceCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyInternetServiceGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyInternetServiceId struct {
	Id *float64 `json:"id,omitempty"`
}

type FirewallShapingPolicyInternetServiceName struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyInternetServiceSrcCustom struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyInternetServiceSrcCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyInternetServiceSrcGroup struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyInternetServiceSrcId struct {
	Id *float64 `json:"id,omitempty"`
}

type FirewallShapingPolicyInternetServiceSrcName struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyService struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicySrcaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicySrcintf struct {
	Name *string `json:"name,omitempty"`
}

type FirewallShapingPolicyUrlCategory struct {
	Id *float64 `json:"id,omitempty"`
}

type FirewallShapingPolicyUsers struct {
	Name *string `json:"name,omitempty"`
}
