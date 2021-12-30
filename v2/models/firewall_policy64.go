package models

const FirewallPolicy64Path = "firewall/policy64/"

type FirewallPolicy64 struct {
	Action               *string                     `json:"action,omitempty"`
	Comments             *string                     `json:"comments,omitempty"`
	Dstaddr              *[]FirewallPolicy64Dstaddr  `json:"dstaddr,omitempty"`
	Dstintf              *string                     `json:"dstintf,omitempty"`
	Fixedport            *string                     `json:"fixedport,omitempty"`
	Ippool               *string                     `json:"ippool,omitempty"`
	Logtraffic           *string                     `json:"logtraffic,omitempty"`
	LogtrafficStart      *string                     `json:"logtraffic-start,omitempty"`
	Name                 *string                     `json:"name,omitempty"`
	PerIpShaper          *string                     `json:"per-ip-shaper,omitempty"`
	PermitAnyHost        *string                     `json:"permit-any-host,omitempty"`
	Policyid             *int64                      `json:"policyid,omitempty"`
	Poolname             *[]FirewallPolicy64Poolname `json:"poolname,omitempty"`
	Schedule             *string                     `json:"schedule,omitempty"`
	Service              *[]FirewallPolicy64Service  `json:"service,omitempty"`
	Srcaddr              *[]FirewallPolicy64Srcaddr  `json:"srcaddr,omitempty"`
	Srcintf              *string                     `json:"srcintf,omitempty"`
	Status               *string                     `json:"status,omitempty"`
	TcpMssReceiver       *int64                      `json:"tcp-mss-receiver,omitempty"`
	TcpMssSender         *int64                      `json:"tcp-mss-sender,omitempty"`
	TrafficShaper        *string                     `json:"traffic-shaper,omitempty"`
	TrafficShaperReverse *string                     `json:"traffic-shaper-reverse,omitempty"`
	Uuid                 *string                     `json:"uuid,omitempty"`
}

type FirewallPolicy64Dstaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy64Poolname struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy64Service struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy64Srcaddr struct {
	Name *string `json:"name,omitempty"`
}
