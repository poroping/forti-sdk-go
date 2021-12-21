package models

const FirewallPolicy46Path = "firewall/policy46/"

type FirewallPolicy46 struct {
	Action               *string                     `json:"action,omitempty"`
	Comments             *string                     `json:"comments,omitempty"`
	Dstaddr              *[]FirewallPolicy46Dstaddr  `json:"dstaddr,omitempty"`
	Dstintf              *string                     `json:"dstintf,omitempty"`
	Fixedport            *string                     `json:"fixedport,omitempty"`
	Ippool               *string                     `json:"ippool,omitempty"`
	Logtraffic           *string                     `json:"logtraffic,omitempty"`
	LogtrafficStart      *string                     `json:"logtraffic-start,omitempty"`
	Name                 *string                     `json:"name,omitempty"`
	PerIpShaper          *string                     `json:"per-ip-shaper,omitempty"`
	PermitAnyHost        *string                     `json:"permit-any-host,omitempty"`
	Policyid             *int64                      `json:"policyid,omitempty"`
	Poolname             *[]FirewallPolicy46Poolname `json:"poolname,omitempty"`
	Schedule             *string                     `json:"schedule,omitempty"`
	Service              *[]FirewallPolicy46Service  `json:"service,omitempty"`
	Srcaddr              *[]FirewallPolicy46Srcaddr  `json:"srcaddr,omitempty"`
	Srcintf              *string                     `json:"srcintf,omitempty"`
	Status               *string                     `json:"status,omitempty"`
	TcpMssReceiver       *int64                      `json:"tcp-mss-receiver,omitempty"`
	TcpMssSender         *int64                      `json:"tcp-mss-sender,omitempty"`
	TrafficShaper        *string                     `json:"traffic-shaper,omitempty"`
	TrafficShaperReverse *string                     `json:"traffic-shaper-reverse,omitempty"`
	Uuid                 *string                     `json:"uuid,omitempty"`
}

type FirewallPolicy46Dstaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy46Poolname struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy46Service struct {
	Name *string `json:"name,omitempty"`
}

type FirewallPolicy46Srcaddr struct {
	Name *string `json:"name,omitempty"`
}
