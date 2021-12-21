package models

const SystemPppoeInterfacePath = "system/pppoe-interface/"

type SystemPppoeInterface struct {
	AcName                   *string  `json:"ac-name,omitempty"`
	AuthType                 *string  `json:"auth-type,omitempty"`
	Device                   *string  `json:"device,omitempty"`
	DialOnDemand             *string  `json:"dial-on-demand,omitempty"`
	DiscRetryTimeout         *float64 `json:"disc-retry-timeout,omitempty"`
	IdleTimeout              *float64 `json:"idle-timeout,omitempty"`
	Ipunnumbered             *string  `json:"ipunnumbered,omitempty"`
	Ipv6                     *string  `json:"ipv6,omitempty"`
	LcpEchoInterval          *float64 `json:"lcp-echo-interval,omitempty"`
	LcpMaxEchoFails          *float64 `json:"lcp-max-echo-fails,omitempty"`
	Name                     *string  `json:"name,omitempty"`
	PadtRetryTimeout         *float64 `json:"padt-retry-timeout,omitempty"`
	Password                 *string  `json:"password,omitempty"`
	PppoeUnnumberedNegotiate *string  `json:"pppoe-unnumbered-negotiate,omitempty"`
	ServiceName              *string  `json:"service-name,omitempty"`
	Username                 *string  `json:"username,omitempty"`
}
