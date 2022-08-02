package models

const SystemNp6xlitePath = "system/np6xlite/"

type SystemNp6xlite struct {
	Fastpath                  *string                  `json:"fastpath,omitempty"`
	FpAnomaly                 *SystemNp6xliteFpAnomaly `json:"fp-anomaly,omitempty"`
	GarbageSessionCollector   *string                  `json:"garbage-session-collector,omitempty"`
	Hpe                       *SystemNp6xliteHpe       `json:"hpe,omitempty"`
	IpsecInnerFragment        *string                  `json:"ipsec-inner-fragment,omitempty"`
	Name                      *string                  `json:"name,omitempty"`
	PerSessionAccounting      *string                  `json:"per-session-accounting,omitempty"`
	SessionCollectorInterval  *int64                   `json:"session-collector-interval,omitempty"`
	SessionTimeoutFixed       *string                  `json:"session-timeout-fixed,omitempty"`
	SessionTimeoutInterval    *int64                   `json:"session-timeout-interval,omitempty"`
	SessionTimeoutRandomRange *int64                   `json:"session-timeout-random-range,omitempty"`
}

const SystemNp6xliteFpAnomalyPath = "system/np6xlite/fp-anomaly/"

type SystemNp6xliteFpAnomaly struct {
	IcmpCsumErr      *string `json:"icmp-csum-err,omitempty"`
	IcmpFrag         *string `json:"icmp-frag,omitempty"`
	IcmpLand         *string `json:"icmp-land,omitempty"`
	Ipv4CsumErr      *string `json:"ipv4-csum-err,omitempty"`
	Ipv4Land         *string `json:"ipv4-land,omitempty"`
	Ipv4Optlsrr      *string `json:"ipv4-optlsrr,omitempty"`
	Ipv4Optrr        *string `json:"ipv4-optrr,omitempty"`
	Ipv4Optsecurity  *string `json:"ipv4-optsecurity,omitempty"`
	Ipv4Optssrr      *string `json:"ipv4-optssrr,omitempty"`
	Ipv4Optstream    *string `json:"ipv4-optstream,omitempty"`
	Ipv4Opttimestamp *string `json:"ipv4-opttimestamp,omitempty"`
	Ipv4ProtoErr     *string `json:"ipv4-proto-err,omitempty"`
	Ipv4Unknopt      *string `json:"ipv4-unknopt,omitempty"`
	Ipv6DaddrErr     *string `json:"ipv6-daddr-err,omitempty"`
	Ipv6Land         *string `json:"ipv6-land,omitempty"`
	Ipv6Optendpid    *string `json:"ipv6-optendpid,omitempty"`
	Ipv6Opthomeaddr  *string `json:"ipv6-opthomeaddr,omitempty"`
	Ipv6Optinvld     *string `json:"ipv6-optinvld,omitempty"`
	Ipv6Optjumbo     *string `json:"ipv6-optjumbo,omitempty"`
	Ipv6Optnsap      *string `json:"ipv6-optnsap,omitempty"`
	Ipv6Optralert    *string `json:"ipv6-optralert,omitempty"`
	Ipv6Opttunnel    *string `json:"ipv6-opttunnel,omitempty"`
	Ipv6ProtoErr     *string `json:"ipv6-proto-err,omitempty"`
	Ipv6SaddrErr     *string `json:"ipv6-saddr-err,omitempty"`
	Ipv6Unknopt      *string `json:"ipv6-unknopt,omitempty"`
	TcpCsumErr       *string `json:"tcp-csum-err,omitempty"`
	TcpFinNoack      *string `json:"tcp-fin-noack,omitempty"`
	TcpFinOnly       *string `json:"tcp-fin-only,omitempty"`
	TcpLand          *string `json:"tcp-land,omitempty"`
	TcpNoFlag        *string `json:"tcp-no-flag,omitempty"`
	TcpSynData       *string `json:"tcp-syn-data,omitempty"`
	TcpSynFin        *string `json:"tcp-syn-fin,omitempty"`
	TcpWinnuke       *string `json:"tcp-winnuke,omitempty"`
	UdpCsumErr       *string `json:"udp-csum-err,omitempty"`
	UdpLand          *string `json:"udp-land,omitempty"`
}

const SystemNp6xliteHpePath = "system/np6xlite/hpe/"

type SystemNp6xliteHpe struct {
	ArpMax       *int64  `json:"arp-max,omitempty"`
	EnableShaper *string `json:"enable-shaper,omitempty"`
	EspMax       *int64  `json:"esp-max,omitempty"`
	IcmpMax      *int64  `json:"icmp-max,omitempty"`
	IpFragMax    *int64  `json:"ip-frag-max,omitempty"`
	IpOthersMax  *int64  `json:"ip-others-max,omitempty"`
	L2OthersMax  *int64  `json:"l2-others-max,omitempty"`
	SctpMax      *int64  `json:"sctp-max,omitempty"`
	TcpMax       *int64  `json:"tcp-max,omitempty"`
	TcpsynMax    *int64  `json:"tcpsyn-max,omitempty"`
	UdpMax       *int64  `json:"udp-max,omitempty"`
}
