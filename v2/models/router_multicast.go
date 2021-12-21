package models

const RouterMulticastPath = "router/multicast/"

type RouterMulticast struct {
	Interface        *[]RouterMulticastInterface   `json:"interface,omitempty"`
	MulticastRouting *string                       `json:"multicast-routing,omitempty"`
	PimSmGlobal      *[]RouterMulticastPimSmGlobal `json:"pim-sm-global,omitempty"`
	RouteLimit       *float64                      `json:"route-limit,omitempty"`
	RouteThreshold   *float64                      `json:"route-threshold,omitempty"`
}

type RouterMulticastInterface struct {
	Bfd                  *string                              `json:"bfd,omitempty"`
	CiscoExcludeGenid    *string                              `json:"cisco-exclude-genid,omitempty"`
	DrPriority           *float64                             `json:"dr-priority,omitempty"`
	HelloHoldtime        *float64                             `json:"hello-holdtime,omitempty"`
	HelloInterval        *float64                             `json:"hello-interval,omitempty"`
	Igmp                 *[]RouterMulticastInterfaceIgmp      `json:"igmp,omitempty"`
	JoinGroup            *[]RouterMulticastInterfaceJoinGroup `json:"join-group,omitempty"`
	MulticastFlow        *string                              `json:"multicast-flow,omitempty"`
	Name                 *string                              `json:"name,omitempty"`
	NeighbourFilter      *string                              `json:"neighbour-filter,omitempty"`
	Passive              *string                              `json:"passive,omitempty"`
	PimMode              *string                              `json:"pim-mode,omitempty"`
	PropagationDelay     *float64                             `json:"propagation-delay,omitempty"`
	RpCandidate          *string                              `json:"rp-candidate,omitempty"`
	RpCandidateGroup     *string                              `json:"rp-candidate-group,omitempty"`
	RpCandidateInterval  *float64                             `json:"rp-candidate-interval,omitempty"`
	RpCandidatePriority  *float64                             `json:"rp-candidate-priority,omitempty"`
	RpfNbrFailBack       *string                              `json:"rpf-nbr-fail-back,omitempty"`
	RpfNbrFailBackFilter *string                              `json:"rpf-nbr-fail-back-filter,omitempty"`
	StateRefreshInterval *float64                             `json:"state-refresh-interval,omitempty"`
	StaticGroup          *string                              `json:"static-group,omitempty"`
	TtlThreshold         *float64                             `json:"ttl-threshold,omitempty"`
}

type RouterMulticastInterfaceIgmp struct {
	AccessGroup             *string  `json:"access-group,omitempty"`
	ImmediateLeaveGroup     *string  `json:"immediate-leave-group,omitempty"`
	LastMemberQueryCount    *float64 `json:"last-member-query-count,omitempty"`
	LastMemberQueryInterval *float64 `json:"last-member-query-interval,omitempty"`
	QueryInterval           *float64 `json:"query-interval,omitempty"`
	QueryMaxResponseTime    *float64 `json:"query-max-response-time,omitempty"`
	QueryTimeout            *float64 `json:"query-timeout,omitempty"`
	RouterAlertCheck        *string  `json:"router-alert-check,omitempty"`
	Version                 *string  `json:"version,omitempty"`
}

type RouterMulticastInterfaceJoinGroup struct {
	Address *string `json:"address,omitempty"`
}

type RouterMulticastPimSmGlobal struct {
	AcceptRegisterList         *string                                `json:"accept-register-list,omitempty"`
	AcceptSourceList           *string                                `json:"accept-source-list,omitempty"`
	BsrAllowQuickRefresh       *string                                `json:"bsr-allow-quick-refresh,omitempty"`
	BsrCandidate               *string                                `json:"bsr-candidate,omitempty"`
	BsrHash                    *float64                               `json:"bsr-hash,omitempty"`
	BsrInterface               *string                                `json:"bsr-interface,omitempty"`
	BsrPriority                *float64                               `json:"bsr-priority,omitempty"`
	CiscoCrpPrefix             *string                                `json:"cisco-crp-prefix,omitempty"`
	CiscoIgnoreRpSetPriority   *string                                `json:"cisco-ignore-rp-set-priority,omitempty"`
	CiscoRegisterChecksum      *string                                `json:"cisco-register-checksum,omitempty"`
	CiscoRegisterChecksumGroup *string                                `json:"cisco-register-checksum-group,omitempty"`
	JoinPruneHoldtime          *float64                               `json:"join-prune-holdtime,omitempty"`
	MessageInterval            *float64                               `json:"message-interval,omitempty"`
	NullRegisterRetries        *float64                               `json:"null-register-retries,omitempty"`
	RegisterRateLimit          *float64                               `json:"register-rate-limit,omitempty"`
	RegisterRpReachability     *string                                `json:"register-rp-reachability,omitempty"`
	RegisterSource             *string                                `json:"register-source,omitempty"`
	RegisterSourceInterface    *string                                `json:"register-source-interface,omitempty"`
	RegisterSourceIp           *string                                `json:"register-source-ip,omitempty"`
	RegisterSupression         *float64                               `json:"register-supression,omitempty"`
	RpAddress                  *[]RouterMulticastPimSmGlobalRpAddress `json:"rp-address,omitempty"`
	RpRegisterKeepalive        *float64                               `json:"rp-register-keepalive,omitempty"`
	SptThreshold               *string                                `json:"spt-threshold,omitempty"`
	SptThresholdGroup          *string                                `json:"spt-threshold-group,omitempty"`
	Ssm                        *string                                `json:"ssm,omitempty"`
	SsmRange                   *string                                `json:"ssm-range,omitempty"`
}

type RouterMulticastPimSmGlobalRpAddress struct {
	Group     *string  `json:"group,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	IpAddress *string  `json:"ip-address,omitempty"`
}
