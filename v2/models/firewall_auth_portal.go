package models

const FirewallAuthPortalPath = "firewall/auth-portal/"

type FirewallAuthPortal struct {
	Groups             *[]FirewallAuthPortalGroups `json:"groups,omitempty"`
	IdentityBasedRoute *string                     `json:"identity-based-route,omitempty"`
	PortalAddr         *string                     `json:"portal-addr,omitempty"`
	PortalAddr6        *string                     `json:"portal-addr6,omitempty"`
	ProxyAuth          *string                     `json:"proxy-auth,omitempty"`
}

type FirewallAuthPortalGroups struct {
	Name *string `json:"name,omitempty"`
}
