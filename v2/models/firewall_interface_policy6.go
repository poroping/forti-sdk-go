package models

const FirewallInterfacePolicy6Path = "firewall/interface-policy6/"

type FirewallInterfacePolicy6 struct {
	ApplicationList          *string                             `json:"application-list,omitempty"`
	ApplicationListStatus    *string                             `json:"application-list-status,omitempty"`
	AvProfile                *string                             `json:"av-profile,omitempty"`
	AvProfileStatus          *string                             `json:"av-profile-status,omitempty"`
	Comments                 *string                             `json:"comments,omitempty"`
	DlpSensor                *string                             `json:"dlp-sensor,omitempty"`
	DlpSensorStatus          *string                             `json:"dlp-sensor-status,omitempty"`
	Dsri                     *string                             `json:"dsri,omitempty"`
	Dstaddr6                 *[]FirewallInterfacePolicy6Dstaddr6 `json:"dstaddr6,omitempty"`
	EmailfilterProfile       *string                             `json:"emailfilter-profile,omitempty"`
	EmailfilterProfileStatus *string                             `json:"emailfilter-profile-status,omitempty"`
	Interface                *string                             `json:"interface,omitempty"`
	IpsSensor                *string                             `json:"ips-sensor,omitempty"`
	IpsSensorStatus          *string                             `json:"ips-sensor-status,omitempty"`
	Logtraffic               *string                             `json:"logtraffic,omitempty"`
	Policyid                 *int64                              `json:"policyid,omitempty"`
	Service6                 *[]FirewallInterfacePolicy6Service6 `json:"service6,omitempty"`
	Srcaddr6                 *[]FirewallInterfacePolicy6Srcaddr6 `json:"srcaddr6,omitempty"`
	Status                   *string                             `json:"status,omitempty"`
	WebfilterProfile         *string                             `json:"webfilter-profile,omitempty"`
	WebfilterProfileStatus   *string                             `json:"webfilter-profile-status,omitempty"`
}

const FirewallInterfacePolicy6Dstaddr6Path = "firewall/interface-policy6/dstaddr6/"

type FirewallInterfacePolicy6Dstaddr6 struct {
	Name *string `json:"name,omitempty"`
}

const FirewallInterfacePolicy6Service6Path = "firewall/interface-policy6/service6/"

type FirewallInterfacePolicy6Service6 struct {
	Name *string `json:"name,omitempty"`
}

const FirewallInterfacePolicy6Srcaddr6Path = "firewall/interface-policy6/srcaddr6/"

type FirewallInterfacePolicy6Srcaddr6 struct {
	Name *string `json:"name,omitempty"`
}
