package models

const SystemNpuPath = "system/npu/"

type SystemNpu struct {
	CapwapOffload               *string                         `json:"capwap-offload,omitempty"`
	DedicatedManagementAffinity *string                         `json:"dedicated-management-affinity,omitempty"`
	DedicatedManagementCpu      *string                         `json:"dedicated-management-cpu,omitempty"`
	DefaultQosType              *string                         `json:"default-qos-type,omitempty"`
	DosOptions                  *[]SystemNpuDosOptions          `json:"dos-options,omitempty"`
	DoubleLevelMcastOffload     *string                         `json:"double-level-mcast-offload,omitempty"`
	DseTimeout                  *float64                        `json:"dse-timeout,omitempty"`
	DswDtsProfile               *[]SystemNpuDswDtsProfile       `json:"dsw-dts-profile,omitempty"`
	DswQueueDtsProfile          *[]SystemNpuDswQueueDtsProfile  `json:"dsw-queue-dts-profile,omitempty"`
	Fastpath                    *string                         `json:"fastpath,omitempty"`
	FpAnomaly                   *[]SystemNpuFpAnomaly           `json:"fp-anomaly,omitempty"`
	GtpSupport                  *string                         `json:"gtp-support,omitempty"`
	HashTblSpread               *string                         `json:"hash-tbl-spread,omitempty"`
	Hpe                         *[]SystemNpuHpe                 `json:"hpe,omitempty"`
	HtabDediQueueNr             *float64                        `json:"htab-dedi-queue-nr,omitempty"`
	HtabMsgQueue                *string                         `json:"htab-msg-queue,omitempty"`
	InboundDscpCopyPort         *[]SystemNpuInboundDscpCopyPort `json:"inbound-dscp-copy-port,omitempty"`
	IpReassembly                *[]SystemNpuIpReassembly        `json:"ip-reassembly,omitempty"`
	IppoolOverloadHigh          *float64                        `json:"ippool-overload-high,omitempty"`
	IppoolOverloadLow           *float64                        `json:"ippool-overload-low,omitempty"`
	IpsecDecSubengineMask       *string                         `json:"ipsec-dec-subengine-mask,omitempty"`
	IpsecEncSubengineMask       *string                         `json:"ipsec-enc-subengine-mask,omitempty"`
	IpsecInboundCache           *string                         `json:"ipsec-inbound-cache,omitempty"`
	IpsecMtuOverride            *string                         `json:"ipsec-mtu-override,omitempty"`
	IpsecObNpSel                *string                         `json:"ipsec-ob-np-sel,omitempty"`
	IpsecOverVlink              *string                         `json:"ipsec-over-vlink,omitempty"`
	IsfNpQueues                 *[]SystemNpuIsfNpQueues         `json:"isf-np-queues,omitempty"`
	MaxSessionTimeout           *float64                        `json:"max-session-timeout,omitempty"`
	McastSessionAccounting      *string                         `json:"mcast-session-accounting,omitempty"`
	NapiBreakInterval           *float64                        `json:"napi-break-interval,omitempty"`
	NpQueues                    *[]SystemNpuNpQueues            `json:"np-queues,omitempty"`
	Np6CpsOptimizationMode      *string                         `json:"np6-cps-optimization-mode,omitempty"`
	PbaEim                      *string                         `json:"pba-eim,omitempty"`
	PerSessionAccounting        *string                         `json:"per-session-accounting,omitempty"`
	PolicyOffloadLevel          *string                         `json:"policy-offload-level,omitempty"`
	PortCpuMap                  *[]SystemNpuPortCpuMap          `json:"port-cpu-map,omitempty"`
	PortNpuMap                  *[]SystemNpuPortNpuMap          `json:"port-npu-map,omitempty"`
	PriorityProtocol            *[]SystemNpuPriorityProtocol    `json:"priority-protocol,omitempty"`
	QosMode                     *string                         `json:"qos-mode,omitempty"`
	RdpOffload                  *string                         `json:"rdp-offload,omitempty"`
	RecoverNp6Link              *string                         `json:"recover-np6-link,omitempty"`
	SessionAcctInterval         *float64                        `json:"session-acct-interval,omitempty"`
	SseBackpressure             *string                         `json:"sse-backpressure,omitempty"`
	StripClearTextPadding       *string                         `json:"strip-clear-text-padding,omitempty"`
	StripEspPadding             *string                         `json:"strip-esp-padding,omitempty"`
	SwNpBandwidth               *string                         `json:"sw-np-bandwidth,omitempty"`
	TcpRstTimeout               *float64                        `json:"tcp-rst-timeout,omitempty"`
	TcpTimeoutProfile           *[]SystemNpuTcpTimeoutProfile   `json:"tcp-timeout-profile,omitempty"`
	UdpTimeoutProfile           *[]SystemNpuUdpTimeoutProfile   `json:"udp-timeout-profile,omitempty"`
	VlanLookupCache             *string                         `json:"vlan-lookup-cache,omitempty"`
}

type SystemNpuDosOptions struct {
	NpuDosMeterMode *string `json:"npu-dos-meter-mode,omitempty"`
	NpuDosTpeMode   *string `json:"npu-dos-tpe-mode,omitempty"`
}

type SystemNpuDswDtsProfile struct {
	Action    *string  `json:"action,omitempty"`
	MinLimit  *float64 `json:"min-limit,omitempty"`
	ProfileId *float64 `json:"profile-id,omitempty"`
	Step      *float64 `json:"step,omitempty"`
}

type SystemNpuDswQueueDtsProfile struct {
	Iport       *string  `json:"iport,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Oport       *string  `json:"oport,omitempty"`
	ProfileId   *float64 `json:"profile-id,omitempty"`
	QueueSelect *float64 `json:"queue-select,omitempty"`
}

type SystemNpuFpAnomaly struct {
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

type SystemNpuHpe struct {
	AllProtocol  *float64 `json:"all-protocol,omitempty"`
	ArpMax       *float64 `json:"arp-max,omitempty"`
	EnableShaper *string  `json:"enable-shaper,omitempty"`
	EspMax       *float64 `json:"esp-max,omitempty"`
	HighPriority *float64 `json:"high-priority,omitempty"`
	IcmpMax      *float64 `json:"icmp-max,omitempty"`
	IpFragMax    *float64 `json:"ip-frag-max,omitempty"`
	IpOthersMax  *float64 `json:"ip-others-max,omitempty"`
	L2OthersMax  *float64 `json:"l2-others-max,omitempty"`
	SctpMax      *float64 `json:"sctp-max,omitempty"`
	TcpMax       *float64 `json:"tcp-max,omitempty"`
	TcpfinRstMax *float64 `json:"tcpfin-rst-max,omitempty"`
	TcpsynAckMax *float64 `json:"tcpsyn-ack-max,omitempty"`
	TcpsynMax    *float64 `json:"tcpsyn-max,omitempty"`
	UdpMax       *float64 `json:"udp-max,omitempty"`
}

type SystemNpuInboundDscpCopyPort struct {
	Interface *string `json:"interface,omitempty"`
}

type SystemNpuIpReassembly struct {
	MaxTimeout *float64 `json:"max-timeout,omitempty"`
	MinTimeout *float64 `json:"min-timeout,omitempty"`
	Status     *string  `json:"status,omitempty"`
}

type SystemNpuIsfNpQueues struct {
	Cos0 *string `json:"cos0,omitempty"`
	Cos1 *string `json:"cos1,omitempty"`
	Cos2 *string `json:"cos2,omitempty"`
	Cos3 *string `json:"cos3,omitempty"`
	Cos4 *string `json:"cos4,omitempty"`
	Cos5 *string `json:"cos5,omitempty"`
	Cos6 *string `json:"cos6,omitempty"`
	Cos7 *string `json:"cos7,omitempty"`
}

type SystemNpuNpQueues struct {
	EthernetType *[]SystemNpuNpQueuesEthernetType `json:"ethernet-type,omitempty"`
	IpProtocol   *[]SystemNpuNpQueuesIpProtocol   `json:"ip-protocol,omitempty"`
	IpService    *[]SystemNpuNpQueuesIpService    `json:"ip-service,omitempty"`
	Profile      *[]SystemNpuNpQueuesProfile      `json:"profile,omitempty"`
	Scheduler    *[]SystemNpuNpQueuesScheduler    `json:"scheduler,omitempty"`
}

type SystemNpuNpQueuesEthernetType struct {
	Name   *string  `json:"name,omitempty"`
	Queue  *float64 `json:"queue,omitempty"`
	Type   *string  `json:"type,omitempty"`
	Weight *float64 `json:"weight,omitempty"`
}

type SystemNpuNpQueuesIpProtocol struct {
	Name     *string  `json:"name,omitempty"`
	Protocol *float64 `json:"protocol,omitempty"`
	Queue    *float64 `json:"queue,omitempty"`
	Weight   *float64 `json:"weight,omitempty"`
}

type SystemNpuNpQueuesIpService struct {
	Dport    *float64 `json:"dport,omitempty"`
	Name     *string  `json:"name,omitempty"`
	Protocol *float64 `json:"protocol,omitempty"`
	Queue    *float64 `json:"queue,omitempty"`
	Sport    *float64 `json:"sport,omitempty"`
	Weight   *float64 `json:"weight,omitempty"`
}

type SystemNpuNpQueuesProfile struct {
	Cos0   *string  `json:"cos0,omitempty"`
	Cos1   *string  `json:"cos1,omitempty"`
	Cos2   *string  `json:"cos2,omitempty"`
	Cos3   *string  `json:"cos3,omitempty"`
	Cos4   *string  `json:"cos4,omitempty"`
	Cos5   *string  `json:"cos5,omitempty"`
	Cos6   *string  `json:"cos6,omitempty"`
	Cos7   *string  `json:"cos7,omitempty"`
	Dscp0  *string  `json:"dscp0,omitempty"`
	Dscp1  *string  `json:"dscp1,omitempty"`
	Dscp10 *string  `json:"dscp10,omitempty"`
	Dscp11 *string  `json:"dscp11,omitempty"`
	Dscp12 *string  `json:"dscp12,omitempty"`
	Dscp13 *string  `json:"dscp13,omitempty"`
	Dscp14 *string  `json:"dscp14,omitempty"`
	Dscp15 *string  `json:"dscp15,omitempty"`
	Dscp16 *string  `json:"dscp16,omitempty"`
	Dscp17 *string  `json:"dscp17,omitempty"`
	Dscp18 *string  `json:"dscp18,omitempty"`
	Dscp19 *string  `json:"dscp19,omitempty"`
	Dscp2  *string  `json:"dscp2,omitempty"`
	Dscp20 *string  `json:"dscp20,omitempty"`
	Dscp21 *string  `json:"dscp21,omitempty"`
	Dscp22 *string  `json:"dscp22,omitempty"`
	Dscp23 *string  `json:"dscp23,omitempty"`
	Dscp24 *string  `json:"dscp24,omitempty"`
	Dscp25 *string  `json:"dscp25,omitempty"`
	Dscp26 *string  `json:"dscp26,omitempty"`
	Dscp27 *string  `json:"dscp27,omitempty"`
	Dscp28 *string  `json:"dscp28,omitempty"`
	Dscp29 *string  `json:"dscp29,omitempty"`
	Dscp3  *string  `json:"dscp3,omitempty"`
	Dscp30 *string  `json:"dscp30,omitempty"`
	Dscp31 *string  `json:"dscp31,omitempty"`
	Dscp32 *string  `json:"dscp32,omitempty"`
	Dscp33 *string  `json:"dscp33,omitempty"`
	Dscp34 *string  `json:"dscp34,omitempty"`
	Dscp35 *string  `json:"dscp35,omitempty"`
	Dscp36 *string  `json:"dscp36,omitempty"`
	Dscp37 *string  `json:"dscp37,omitempty"`
	Dscp38 *string  `json:"dscp38,omitempty"`
	Dscp39 *string  `json:"dscp39,omitempty"`
	Dscp4  *string  `json:"dscp4,omitempty"`
	Dscp40 *string  `json:"dscp40,omitempty"`
	Dscp41 *string  `json:"dscp41,omitempty"`
	Dscp42 *string  `json:"dscp42,omitempty"`
	Dscp43 *string  `json:"dscp43,omitempty"`
	Dscp44 *string  `json:"dscp44,omitempty"`
	Dscp45 *string  `json:"dscp45,omitempty"`
	Dscp46 *string  `json:"dscp46,omitempty"`
	Dscp47 *string  `json:"dscp47,omitempty"`
	Dscp48 *string  `json:"dscp48,omitempty"`
	Dscp49 *string  `json:"dscp49,omitempty"`
	Dscp5  *string  `json:"dscp5,omitempty"`
	Dscp50 *string  `json:"dscp50,omitempty"`
	Dscp51 *string  `json:"dscp51,omitempty"`
	Dscp52 *string  `json:"dscp52,omitempty"`
	Dscp53 *string  `json:"dscp53,omitempty"`
	Dscp54 *string  `json:"dscp54,omitempty"`
	Dscp55 *string  `json:"dscp55,omitempty"`
	Dscp56 *string  `json:"dscp56,omitempty"`
	Dscp57 *string  `json:"dscp57,omitempty"`
	Dscp58 *string  `json:"dscp58,omitempty"`
	Dscp59 *string  `json:"dscp59,omitempty"`
	Dscp6  *string  `json:"dscp6,omitempty"`
	Dscp60 *string  `json:"dscp60,omitempty"`
	Dscp61 *string  `json:"dscp61,omitempty"`
	Dscp62 *string  `json:"dscp62,omitempty"`
	Dscp63 *string  `json:"dscp63,omitempty"`
	Dscp7  *string  `json:"dscp7,omitempty"`
	Dscp8  *string  `json:"dscp8,omitempty"`
	Dscp9  *string  `json:"dscp9,omitempty"`
	Id     *float64 `json:"id,omitempty"`
	Type   *string  `json:"type,omitempty"`
	Weight *float64 `json:"weight,omitempty"`
}

type SystemNpuNpQueuesScheduler struct {
	Mode *string `json:"mode,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SystemNpuPortCpuMap struct {
	CpuCore   *string `json:"cpu-core,omitempty"`
	Interface *string `json:"interface,omitempty"`
}

type SystemNpuPortNpuMap struct {
	Interface     *string  `json:"interface,omitempty"`
	NpuGroupIndex *float64 `json:"npu-group-index,omitempty"`
}

type SystemNpuPriorityProtocol struct {
	Bfd  *string `json:"bfd,omitempty"`
	Bgp  *string `json:"bgp,omitempty"`
	Slbc *string `json:"slbc,omitempty"`
}

type SystemNpuTcpTimeoutProfile struct {
	CloseWait *float64 `json:"close-wait,omitempty"`
	FinWait   *float64 `json:"fin-wait,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	SynSent   *float64 `json:"syn-sent,omitempty"`
	SynWait   *float64 `json:"syn-wait,omitempty"`
	TcpIdle   *float64 `json:"tcp-idle,omitempty"`
	TimeWait  *float64 `json:"time-wait,omitempty"`
}

type SystemNpuUdpTimeoutProfile struct {
	Id      *float64 `json:"id,omitempty"`
	UdpIdle *float64 `json:"udp-idle,omitempty"`
}
