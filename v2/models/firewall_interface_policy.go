package models

const FirewallInterfacePolicyPath = "firewall/interface-policy/"

type FirewallInterfacePolicy struct {
	ApplicationList          *string                           `json:"application-list,omitempty"`
	ApplicationListStatus    *string                           `json:"application-list-status,omitempty"`
	AvProfile                *string                           `json:"av-profile,omitempty"`
	AvProfileStatus          *string                           `json:"av-profile-status,omitempty"`
	Comments                 *string                           `json:"comments,omitempty"`
	DlpSensor                *string                           `json:"dlp-sensor,omitempty"`
	DlpSensorStatus          *string                           `json:"dlp-sensor-status,omitempty"`
	Dsri                     *string                           `json:"dsri,omitempty"`
	Dstaddr                  *[]FirewallInterfacePolicyDstaddr `json:"dstaddr,omitempty"`
	EmailfilterProfile       *string                           `json:"emailfilter-profile,omitempty"`
	EmailfilterProfileStatus *string                           `json:"emailfilter-profile-status,omitempty"`
	Interface                *string                           `json:"interface,omitempty"`
	IpsSensor                *string                           `json:"ips-sensor,omitempty"`
	IpsSensorStatus          *string                           `json:"ips-sensor-status,omitempty"`
	Logtraffic               *string                           `json:"logtraffic,omitempty"`
	Policyid                 *int64                            `json:"policyid,omitempty"`
	Service                  *[]FirewallInterfacePolicyService `json:"service,omitempty"`
	Srcaddr                  *[]FirewallInterfacePolicySrcaddr `json:"srcaddr,omitempty"`
	Status                   *string                           `json:"status,omitempty"`
	WebfilterProfile         *string                           `json:"webfilter-profile,omitempty"`
	WebfilterProfileStatus   *string                           `json:"webfilter-profile-status,omitempty"`
}

const FirewallInterfacePolicyDstaddrPath = "firewall/interface-policy/dstaddr/"

type FirewallInterfacePolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

const FirewallInterfacePolicyServicePath = "firewall/interface-policy/service/"

type FirewallInterfacePolicyService struct {
	Name *string `json:"name,omitempty"`
}

const FirewallInterfacePolicySrcaddrPath = "firewall/interface-policy/srcaddr/"

type FirewallInterfacePolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}
