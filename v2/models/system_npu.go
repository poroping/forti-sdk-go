package models

const SystemNpuPath = "system/npu/"

type SystemNpu struct {
	CapwapOffload               *string                         `json:"capwap-offload,omitempty"`
	DedicatedManagementAffinity *string                         `json:"dedicated-management-affinity,omitempty"`
	DedicatedManagementCpu      *string                         `json:"dedicated-management-cpu,omitempty"`
	DefaultQosType              *string                         `json:"default-qos-type,omitempty"`
	DosOptions                  *SystemNpuDosOptions            `json:"dos-options,omitempty"`
	DswDtsProfile               *[]SystemNpuDswDtsProfile       `json:"dsw-dts-profile,omitempty"`
	DswQueueDtsProfile          *[]SystemNpuDswQueueDtsProfile  `json:"dsw-queue-dts-profile,omitempty"`
	FpAnomaly                   *SystemNpuFpAnomaly             `json:"fp-anomaly,omitempty"`
	GtpSupport                  *string                         `json:"gtp-support,omitempty"`
	HashTblSpread               *string                         `json:"hash-tbl-spread,omitempty"`
	Hpe                         *SystemNpuHpe                   `json:"hpe,omitempty"`
	HtabDediQueueNr             *int64                          `json:"htab-dedi-queue-nr,omitempty"`
	HtabMsgQueue                *string                         `json:"htab-msg-queue,omitempty"`
	HtxIcmpCsumChk              *string                         `json:"htx-icmp-csum-chk,omitempty"`
	InboundDscpCopyPort         *[]SystemNpuInboundDscpCopyPort `json:"inbound-dscp-copy-port,omitempty"`
	IpReassembly                *SystemNpuIpReassembly          `json:"ip-reassembly,omitempty"`
	IpsecMtuOverride            *string                         `json:"ipsec-mtu-override,omitempty"`
	IpsecObNpSel                *string                         `json:"ipsec-ob-np-sel,omitempty"`
	MaxSessionTimeout           *int64                          `json:"max-session-timeout,omitempty"`
	NapiBreakInterval           *int64                          `json:"napi-break-interval,omitempty"`
	NpQueues                    *SystemNpuNpQueues              `json:"np-queues,omitempty"`
	PerPolicyAccounting         *string                         `json:"per-policy-accounting,omitempty"`
	PerSessionAccounting        *string                         `json:"per-session-accounting,omitempty"`
	PolicyOffloadLevel          *string                         `json:"policy-offload-level,omitempty"`
	PortNpuMap                  *[]SystemNpuPortNpuMap          `json:"port-npu-map,omitempty"`
	QtmBufMode                  *string                         `json:"qtm-buf-mode,omitempty"`
	SessionAcctInterval         *int64                          `json:"session-acct-interval,omitempty"`
	VlanLookupCache             *string                         `json:"vlan-lookup-cache,omitempty"`
}

const SystemNpuDosOptionsPath = "system/npu/dos-options/"

type SystemNpuDosOptions struct {
	NpuDosMeterMode *string `json:"npu-dos-meter-mode,omitempty"`
	NpuDosTpeMode   *string `json:"npu-dos-tpe-mode,omitempty"`
}

const SystemNpuDswDtsProfilePath = "system/npu/dsw-dts-profile/"

type SystemNpuDswDtsProfile struct {
	Action    *string `json:"action,omitempty"`
	MinLimit  *int64  `json:"min-limit,omitempty"`
	ProfileId *int64  `json:"profile-id,omitempty"`
	Step      *int64  `json:"step,omitempty"`
}

const SystemNpuDswQueueDtsProfilePath = "system/npu/dsw-queue-dts-profile/"

type SystemNpuDswQueueDtsProfile struct {
	Iport       *string `json:"iport,omitempty"`
	Name        *string `json:"name,omitempty"`
	Oport       *string `json:"oport,omitempty"`
	ProfileId   *int64  `json:"profile-id,omitempty"`
	QueueSelect *int64  `json:"queue-select,omitempty"`
}

const SystemNpuFpAnomalyPath = "system/npu/fp-anomaly/"

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

const SystemNpuHpePath = "system/npu/hpe/"

type SystemNpuHpe struct {
	AllProtocol  *int64  `json:"all-protocol,omitempty"`
	ArpMax       *int64  `json:"arp-max,omitempty"`
	EnableShaper *string `json:"enable-shaper,omitempty"`
	EspMax       *int64  `json:"esp-max,omitempty"`
	HighPriority *int64  `json:"high-priority,omitempty"`
	IcmpMax      *int64  `json:"icmp-max,omitempty"`
	IpFragMax    *int64  `json:"ip-frag-max,omitempty"`
	IpOthersMax  *int64  `json:"ip-others-max,omitempty"`
	L2OthersMax  *int64  `json:"l2-others-max,omitempty"`
	SctpMax      *int64  `json:"sctp-max,omitempty"`
	TcpMax       *int64  `json:"tcp-max,omitempty"`
	TcpfinRstMax *int64  `json:"tcpfin-rst-max,omitempty"`
	TcpsynAckMax *int64  `json:"tcpsyn-ack-max,omitempty"`
	TcpsynMax    *int64  `json:"tcpsyn-max,omitempty"`
	UdpMax       *int64  `json:"udp-max,omitempty"`
}

const SystemNpuInboundDscpCopyPortPath = "system/npu/inbound-dscp-copy-port/"

type SystemNpuInboundDscpCopyPort struct {
	Interface *string `json:"interface,omitempty"`
}

const SystemNpuIpReassemblyPath = "system/npu/ip-reassembly/"

type SystemNpuIpReassembly struct {
	MaxTimeout *int64  `json:"max-timeout,omitempty"`
	MinTimeout *int64  `json:"min-timeout,omitempty"`
	Status     *string `json:"status,omitempty"`
}

const SystemNpuNpQueuesPath = "system/npu/np-queues/"

type SystemNpuNpQueues struct {
	EthernetType *[]SystemNpuNpQueuesEthernetType `json:"ethernet-type,omitempty"`
	IpProtocol   *[]SystemNpuNpQueuesIpProtocol   `json:"ip-protocol,omitempty"`
	IpService    *[]SystemNpuNpQueuesIpService    `json:"ip-service,omitempty"`
	Profile      *[]SystemNpuNpQueuesProfile      `json:"profile,omitempty"`
	Scheduler    *[]SystemNpuNpQueuesScheduler    `json:"scheduler,omitempty"`
}

const SystemNpuNpQueuesEthernetTypePath = "system/npu/np-queues/ethernet-type/"

type SystemNpuNpQueuesEthernetType struct {
	Name   *string `json:"name,omitempty"`
	Queue  *int64  `json:"queue,omitempty"`
	Type   *string `json:"type,omitempty"`
	Weight *int64  `json:"weight,omitempty"`
}

const SystemNpuNpQueuesIpProtocolPath = "system/npu/np-queues/ip-protocol/"

type SystemNpuNpQueuesIpProtocol struct {
	Name     *string `json:"name,omitempty"`
	Protocol *int64  `json:"protocol,omitempty"`
	Queue    *int64  `json:"queue,omitempty"`
	Weight   *int64  `json:"weight,omitempty"`
}

const SystemNpuNpQueuesIpServicePath = "system/npu/np-queues/ip-service/"

type SystemNpuNpQueuesIpService struct {
	Dport    *int64  `json:"dport,omitempty"`
	Name     *string `json:"name,omitempty"`
	Protocol *int64  `json:"protocol,omitempty"`
	Queue    *int64  `json:"queue,omitempty"`
	Sport    *int64  `json:"sport,omitempty"`
	Weight   *int64  `json:"weight,omitempty"`
}

const SystemNpuNpQueuesProfilePath = "system/npu/np-queues/profile/"

type SystemNpuNpQueuesProfile struct {
	Cos0   *string `json:"cos0,omitempty"`
	Cos1   *string `json:"cos1,omitempty"`
	Cos2   *string `json:"cos2,omitempty"`
	Cos3   *string `json:"cos3,omitempty"`
	Cos4   *string `json:"cos4,omitempty"`
	Cos5   *string `json:"cos5,omitempty"`
	Cos6   *string `json:"cos6,omitempty"`
	Cos7   *string `json:"cos7,omitempty"`
	Dscp0  *string `json:"dscp0,omitempty"`
	Dscp1  *string `json:"dscp1,omitempty"`
	Dscp10 *string `json:"dscp10,omitempty"`
	Dscp11 *string `json:"dscp11,omitempty"`
	Dscp12 *string `json:"dscp12,omitempty"`
	Dscp13 *string `json:"dscp13,omitempty"`
	Dscp14 *string `json:"dscp14,omitempty"`
	Dscp15 *string `json:"dscp15,omitempty"`
	Dscp16 *string `json:"dscp16,omitempty"`
	Dscp17 *string `json:"dscp17,omitempty"`
	Dscp18 *string `json:"dscp18,omitempty"`
	Dscp19 *string `json:"dscp19,omitempty"`
	Dscp2  *string `json:"dscp2,omitempty"`
	Dscp20 *string `json:"dscp20,omitempty"`
	Dscp21 *string `json:"dscp21,omitempty"`
	Dscp22 *string `json:"dscp22,omitempty"`
	Dscp23 *string `json:"dscp23,omitempty"`
	Dscp24 *string `json:"dscp24,omitempty"`
	Dscp25 *string `json:"dscp25,omitempty"`
	Dscp26 *string `json:"dscp26,omitempty"`
	Dscp27 *string `json:"dscp27,omitempty"`
	Dscp28 *string `json:"dscp28,omitempty"`
	Dscp29 *string `json:"dscp29,omitempty"`
	Dscp3  *string `json:"dscp3,omitempty"`
	Dscp30 *string `json:"dscp30,omitempty"`
	Dscp31 *string `json:"dscp31,omitempty"`
	Dscp32 *string `json:"dscp32,omitempty"`
	Dscp33 *string `json:"dscp33,omitempty"`
	Dscp34 *string `json:"dscp34,omitempty"`
	Dscp35 *string `json:"dscp35,omitempty"`
	Dscp36 *string `json:"dscp36,omitempty"`
	Dscp37 *string `json:"dscp37,omitempty"`
	Dscp38 *string `json:"dscp38,omitempty"`
	Dscp39 *string `json:"dscp39,omitempty"`
	Dscp4  *string `json:"dscp4,omitempty"`
	Dscp40 *string `json:"dscp40,omitempty"`
	Dscp41 *string `json:"dscp41,omitempty"`
	Dscp42 *string `json:"dscp42,omitempty"`
	Dscp43 *string `json:"dscp43,omitempty"`
	Dscp44 *string `json:"dscp44,omitempty"`
	Dscp45 *string `json:"dscp45,omitempty"`
	Dscp46 *string `json:"dscp46,omitempty"`
	Dscp47 *string `json:"dscp47,omitempty"`
	Dscp48 *string `json:"dscp48,omitempty"`
	Dscp49 *string `json:"dscp49,omitempty"`
	Dscp5  *string `json:"dscp5,omitempty"`
	Dscp50 *string `json:"dscp50,omitempty"`
	Dscp51 *string `json:"dscp51,omitempty"`
	Dscp52 *string `json:"dscp52,omitempty"`
	Dscp53 *string `json:"dscp53,omitempty"`
	Dscp54 *string `json:"dscp54,omitempty"`
	Dscp55 *string `json:"dscp55,omitempty"`
	Dscp56 *string `json:"dscp56,omitempty"`
	Dscp57 *string `json:"dscp57,omitempty"`
	Dscp58 *string `json:"dscp58,omitempty"`
	Dscp59 *string `json:"dscp59,omitempty"`
	Dscp6  *string `json:"dscp6,omitempty"`
	Dscp60 *string `json:"dscp60,omitempty"`
	Dscp61 *string `json:"dscp61,omitempty"`
	Dscp62 *string `json:"dscp62,omitempty"`
	Dscp63 *string `json:"dscp63,omitempty"`
	Dscp7  *string `json:"dscp7,omitempty"`
	Dscp8  *string `json:"dscp8,omitempty"`
	Dscp9  *string `json:"dscp9,omitempty"`
	Id     *int64  `json:"id,omitempty"`
	Type   *string `json:"type,omitempty"`
	Weight *int64  `json:"weight,omitempty"`
}

const SystemNpuNpQueuesSchedulerPath = "system/npu/np-queues/scheduler/"

type SystemNpuNpQueuesScheduler struct {
	Mode *string `json:"mode,omitempty"`
	Name *string `json:"name,omitempty"`
}

const SystemNpuPortNpuMapPath = "system/npu/port-npu-map/"

type SystemNpuPortNpuMap struct {
	Interface     *string `json:"interface,omitempty"`
	NpuGroupIndex *int64  `json:"npu-group-index,omitempty"`
}
