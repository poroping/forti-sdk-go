package models

const FirewallTrafficClassPath = "firewall/traffic-class/"

type FirewallTrafficClass struct {
	ClassId   *float64 `json:"class-id,omitempty"`
	ClassName *string  `json:"class-name,omitempty"`
}
