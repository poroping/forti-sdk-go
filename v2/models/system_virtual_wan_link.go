package models

const SystemVirtualWanLinkPath = "system/virtual-wan-link/"

type SystemVirtualWanLink struct {
	FailAlertInterfaces  *[]SystemVirtualWanLinkFailAlertInterfaces `json:"fail-alert-interfaces,omitempty"`
	FailDetect           *string                                    `json:"fail-detect,omitempty"`
	HealthCheck          *[]SystemVirtualWanLinkHealthCheck         `json:"health-check,omitempty"`
	LoadBalanceMode      *string                                    `json:"load-balance-mode,omitempty"`
	Members              *[]SystemVirtualWanLinkMembers             `json:"members,omitempty"`
	Neighbor             *[]SystemVirtualWanLinkNeighbor            `json:"neighbor,omitempty"`
	NeighborHoldBootTime *int64                                     `json:"neighbor-hold-boot-time,omitempty"`
	NeighborHoldDown     *string                                    `json:"neighbor-hold-down,omitempty"`
	NeighborHoldDownTime *int64                                     `json:"neighbor-hold-down-time,omitempty"`
	Service              *[]SystemVirtualWanLinkService             `json:"service,omitempty"`
	Status               *string                                    `json:"status,omitempty"`
	Zone                 *[]SystemVirtualWanLinkZone                `json:"zone,omitempty"`
}

type SystemVirtualWanLinkFailAlertInterfaces struct {
	Name *string `json:"name,omitempty"`
}

type SystemVirtualWanLinkHealthCheck struct {
	AddrMode                   *string                                   `json:"addr-mode,omitempty"`
	Diffservcode               *string                                   `json:"diffservcode,omitempty"`
	Failtime                   *int64                                    `json:"failtime,omitempty"`
	HaPriority                 *int64                                    `json:"ha-priority,omitempty"`
	HttpAgent                  *string                                   `json:"http-agent,omitempty"`
	HttpGet                    *string                                   `json:"http-get,omitempty"`
	HttpMatch                  *string                                   `json:"http-match,omitempty"`
	Interval                   *int64                                    `json:"interval,omitempty"`
	Members                    *[]SystemVirtualWanLinkHealthCheckMembers `json:"members,omitempty"`
	Name                       *string                                   `json:"name,omitempty"`
	PacketSize                 *int64                                    `json:"packet-size,omitempty"`
	Password                   *string                                   `json:"password,omitempty"`
	Port                       *int64                                    `json:"port,omitempty"`
	ProbePackets               *string                                   `json:"probe-packets,omitempty"`
	ProbeTimeout               *int64                                    `json:"probe-timeout,omitempty"`
	Protocol                   *string                                   `json:"protocol,omitempty"`
	Recoverytime               *int64                                    `json:"recoverytime,omitempty"`
	SecurityMode               *string                                   `json:"security-mode,omitempty"`
	Server                     *string                                   `json:"server,omitempty"`
	Sla                        *[]SystemVirtualWanLinkHealthCheckSla     `json:"sla,omitempty"`
	SlaFailLogPeriod           *int64                                    `json:"sla-fail-log-period,omitempty"`
	SlaPassLogPeriod           *int64                                    `json:"sla-pass-log-period,omitempty"`
	ThresholdAlertJitter       *int64                                    `json:"threshold-alert-jitter,omitempty"`
	ThresholdAlertLatency      *int64                                    `json:"threshold-alert-latency,omitempty"`
	ThresholdAlertPacketloss   *int64                                    `json:"threshold-alert-packetloss,omitempty"`
	ThresholdWarningJitter     *int64                                    `json:"threshold-warning-jitter,omitempty"`
	ThresholdWarningLatency    *int64                                    `json:"threshold-warning-latency,omitempty"`
	ThresholdWarningPacketloss *int64                                    `json:"threshold-warning-packetloss,omitempty"`
	UpdateCascadeInterface     *string                                   `json:"update-cascade-interface,omitempty"`
	UpdateStaticRoute          *string                                   `json:"update-static-route,omitempty"`
}

type SystemVirtualWanLinkHealthCheckMembers struct {
	SeqNum *int64 `json:"seq-num,omitempty"`
}

type SystemVirtualWanLinkHealthCheckSla struct {
	Id                  *int64  `json:"id,omitempty"`
	JitterThreshold     *int64  `json:"jitter-threshold,omitempty"`
	LatencyThreshold    *int64  `json:"latency-threshold,omitempty"`
	LinkCostFactor      *string `json:"link-cost-factor,omitempty"`
	PacketlossThreshold *int64  `json:"packetloss-threshold,omitempty"`
}

type SystemVirtualWanLinkMembers struct {
	Comment                   *string `json:"comment,omitempty"`
	Cost                      *int64  `json:"cost,omitempty"`
	Gateway                   *string `json:"gateway,omitempty"`
	Gateway6                  *string `json:"gateway6,omitempty"`
	IngressSpilloverThreshold *int64  `json:"ingress-spillover-threshold,omitempty"`
	Interface                 *string `json:"interface,omitempty"`
	Priority                  *int64  `json:"priority,omitempty"`
	SeqNum                    *int64  `json:"seq-num,omitempty"`
	Source                    *string `json:"source,omitempty"`
	Source6                   *string `json:"source6,omitempty"`
	SpilloverThreshold        *int64  `json:"spillover-threshold,omitempty"`
	Status                    *string `json:"status,omitempty"`
	VolumeRatio               *int64  `json:"volume-ratio,omitempty"`
	Weight                    *int64  `json:"weight,omitempty"`
}

type SystemVirtualWanLinkNeighbor struct {
	HealthCheck *string `json:"health-check,omitempty"`
	Ip          *string `json:"ip,omitempty"`
	Member      *int64  `json:"member,omitempty"`
	Role        *string `json:"role,omitempty"`
	SlaId       *int64  `json:"sla-id,omitempty"`
}

type SystemVirtualWanLinkService struct {
	AddrMode                    *string                                                   `json:"addr-mode,omitempty"`
	BandwidthWeight             *int64                                                    `json:"bandwidth-weight,omitempty"`
	Default                     *string                                                   `json:"default,omitempty"`
	DscpForward                 *string                                                   `json:"dscp-forward,omitempty"`
	DscpForwardTag              *string                                                   `json:"dscp-forward-tag,omitempty"`
	DscpReverse                 *string                                                   `json:"dscp-reverse,omitempty"`
	DscpReverseTag              *string                                                   `json:"dscp-reverse-tag,omitempty"`
	Dst                         *[]SystemVirtualWanLinkServiceDst                         `json:"dst,omitempty"`
	DstNegate                   *string                                                   `json:"dst-negate,omitempty"`
	Dst6                        *[]SystemVirtualWanLinkServiceDst6                        `json:"dst6,omitempty"`
	EndPort                     *int64                                                    `json:"end-port,omitempty"`
	Gateway                     *string                                                   `json:"gateway,omitempty"`
	Groups                      *[]SystemVirtualWanLinkServiceGroups                      `json:"groups,omitempty"`
	HealthCheck                 *string                                                   `json:"health-check,omitempty"`
	HoldDownTime                *int64                                                    `json:"hold-down-time,omitempty"`
	Id                          *int64                                                    `json:"id,omitempty"`
	InputDevice                 *[]SystemVirtualWanLinkServiceInputDevice                 `json:"input-device,omitempty"`
	InputDeviceNegate           *string                                                   `json:"input-device-negate,omitempty"`
	InternetService             *string                                                   `json:"internet-service,omitempty"`
	InternetServiceAppCtrl      *[]SystemVirtualWanLinkServiceInternetServiceAppCtrl      `json:"internet-service-app-ctrl,omitempty"`
	InternetServiceAppCtrlGroup *[]SystemVirtualWanLinkServiceInternetServiceAppCtrlGroup `json:"internet-service-app-ctrl-group,omitempty"`
	InternetServiceCustom       *[]SystemVirtualWanLinkServiceInternetServiceCustom       `json:"internet-service-custom,omitempty"`
	InternetServiceCustomGroup  *[]SystemVirtualWanLinkServiceInternetServiceCustomGroup  `json:"internet-service-custom-group,omitempty"`
	InternetServiceGroup        *[]SystemVirtualWanLinkServiceInternetServiceGroup        `json:"internet-service-group,omitempty"`
	InternetServiceId           *[]SystemVirtualWanLinkServiceInternetServiceId           `json:"internet-service-id,omitempty"`
	JitterWeight                *int64                                                    `json:"jitter-weight,omitempty"`
	LatencyWeight               *int64                                                    `json:"latency-weight,omitempty"`
	LinkCostFactor              *string                                                   `json:"link-cost-factor,omitempty"`
	LinkCostThreshold           *int64                                                    `json:"link-cost-threshold,omitempty"`
	Mode                        *string                                                   `json:"mode,omitempty"`
	Name                        *string                                                   `json:"name,omitempty"`
	PacketLossWeight            *int64                                                    `json:"packet-loss-weight,omitempty"`
	PriorityMembers             *[]SystemVirtualWanLinkServicePriorityMembers             `json:"priority-members,omitempty"`
	Protocol                    *int64                                                    `json:"protocol,omitempty"`
	QualityLink                 *int64                                                    `json:"quality-link,omitempty"`
	Role                        *string                                                   `json:"role,omitempty"`
	RouteTag                    *int64                                                    `json:"route-tag,omitempty"`
	Sla                         *[]SystemVirtualWanLinkServiceSla                         `json:"sla,omitempty"`
	SlaCompareMethod            *string                                                   `json:"sla-compare-method,omitempty"`
	Src                         *[]SystemVirtualWanLinkServiceSrc                         `json:"src,omitempty"`
	SrcNegate                   *string                                                   `json:"src-negate,omitempty"`
	Src6                        *[]SystemVirtualWanLinkServiceSrc6                        `json:"src6,omitempty"`
	StandaloneAction            *string                                                   `json:"standalone-action,omitempty"`
	StartPort                   *int64                                                    `json:"start-port,omitempty"`
	Status                      *string                                                   `json:"status,omitempty"`
	Tos                         *string                                                   `json:"tos,omitempty"`
	TosMask                     *string                                                   `json:"tos-mask,omitempty"`
	Users                       *[]SystemVirtualWanLinkServiceUsers                       `json:"users,omitempty"`
}

type SystemVirtualWanLinkServiceDst struct {
	Name *string `json:"name,omitempty"`
}

type SystemVirtualWanLinkServiceDst6 struct {
	Name *string `json:"name,omitempty"`
}

type SystemVirtualWanLinkServiceGroups struct {
	Name *string `json:"name,omitempty"`
}

type SystemVirtualWanLinkServiceInputDevice struct {
	Name *string `json:"name,omitempty"`
}

type SystemVirtualWanLinkServiceInternetServiceAppCtrl struct {
	Id *int64 `json:"id,omitempty"`
}

type SystemVirtualWanLinkServiceInternetServiceAppCtrlGroup struct {
	Name *string `json:"name,omitempty"`
}

type SystemVirtualWanLinkServiceInternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

type SystemVirtualWanLinkServiceInternetServiceCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type SystemVirtualWanLinkServiceInternetServiceGroup struct {
	Name *string `json:"name,omitempty"`
}

type SystemVirtualWanLinkServiceInternetServiceId struct {
	Id *int64 `json:"id,omitempty"`
}

type SystemVirtualWanLinkServicePriorityMembers struct {
	SeqNum *int64 `json:"seq-num,omitempty"`
}

type SystemVirtualWanLinkServiceSla struct {
	HealthCheck *string `json:"health-check,omitempty"`
	Id          *int64  `json:"id,omitempty"`
}

type SystemVirtualWanLinkServiceSrc struct {
	Name *string `json:"name,omitempty"`
}

type SystemVirtualWanLinkServiceSrc6 struct {
	Name *string `json:"name,omitempty"`
}

type SystemVirtualWanLinkServiceUsers struct {
	Name *string `json:"name,omitempty"`
}

type SystemVirtualWanLinkZone struct {
	Name *string `json:"name,omitempty"`
}
