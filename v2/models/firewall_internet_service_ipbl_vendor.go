package models

const FirewallInternetServiceIpblVendorPath = "firewall/internet-service-ipbl-vendor/"

type FirewallInternetServiceIpblVendor struct {
	Fosid *int64  `json:"fosid,omitempty"`
	Name  *string `json:"name,omitempty"`
}
