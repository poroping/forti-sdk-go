package models

const FirewallInternetServiceSldPath = "firewall/internet-service-sld/"

type FirewallInternetServiceSld struct {
	Fosid *float64 `json:"fosid,omitempty"`
	Name  *string  `json:"name,omitempty"`
}
