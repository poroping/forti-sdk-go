package models

const FirewallshaperPerIpShaperPath = "firewall.shaper/per-ip-shaper/"

type FirewallshaperPerIpShaper struct {
	BandwidthUnit           *string `json:"bandwidth-unit,omitempty"`
	DiffservForward         *string `json:"diffserv-forward,omitempty"`
	DiffservReverse         *string `json:"diffserv-reverse,omitempty"`
	DiffservcodeForward     *string `json:"diffservcode-forward,omitempty"`
	DiffservcodeRev         *string `json:"diffservcode-rev,omitempty"`
	MaxBandwidth            *int64  `json:"max-bandwidth,omitempty"`
	MaxConcurrentSession    *int64  `json:"max-concurrent-session,omitempty"`
	MaxConcurrentTcpSession *int64  `json:"max-concurrent-tcp-session,omitempty"`
	MaxConcurrentUdpSession *int64  `json:"max-concurrent-udp-session,omitempty"`
	Name                    *string `json:"name,omitempty"`
}
