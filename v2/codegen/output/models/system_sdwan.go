package models

const SystemSdwanPath = "system/sdwan/"

type SystemSdwan struct {
	Duplication            []SystemSdwanDuplication         `json:"duplication,omitempty"`
	DuplicationMaxNum      *int64                           `json:"duplication-max-num,omitempty"`
	FailAlertInterfaces    []SystemSdwanFailAlertInterfaces `json:"fail-alert-interfaces,omitempty"`
	FailDetect             *string                          `json:"fail-detect,omitempty"`
	HealthCheck            []SystemSdwanHealthCheck         `json:"health-check,omitempty"`
	LoadBalanceMode        *string                          `json:"load-balance-mode,omitempty"`
	Members                []SystemSdwanMembers             `json:"members,omitempty"`
	Neighbor               []SystemSdwanNeighbor            `json:"neighbor,omitempty"`
	NeighborHoldBootTime   *int64                           `json:"neighbor-hold-boot-time,omitempty"`
	NeighborHoldDown       *string                          `json:"neighbor-hold-down,omitempty"`
	NeighborHoldDownTime   *int64                           `json:"neighbor-hold-down-time,omitempty"`
	Service                []SystemSdwanService             `json:"service,omitempty"`
	SpeedtestBypassRouting *string                          `json:"speedtest-bypass-routing,omitempty"`
	Status                 *string                          `json:"status,omitempty"`
	Zone                   []SystemSdwanZone                `json:"zone,omitempty"`
}

type SystemSdwanDuplication struct {
	Dstaddr             []SystemSdwanDuplicationDstaddr   `json:"dstaddr,omitempty"`
	Dstaddr6            []SystemSdwanDuplicationDstaddr6  `json:"dstaddr6,omitempty"`
	Dstintf             []SystemSdwanDuplicationDstintf   `json:"dstintf,omitempty"`
	Id                  *int64                            `json:"id,omitempty"`
	PacketDeDuplication *string                           `json:"packet-de-duplication,omitempty"`
	PacketDuplication   *string                           `json:"packet-duplication,omitempty"`
	Service             []SystemSdwanDuplicationService   `json:"service,omitempty"`
	ServiceId           []SystemSdwanDuplicationServiceId `json:"service-id,omitempty"`
	Srcaddr             []SystemSdwanDuplicationSrcaddr   `json:"srcaddr,omitempty"`
	Srcaddr6            []SystemSdwanDuplicationSrcaddr6  `json:"srcaddr6,omitempty"`
	Srcintf             []SystemSdwanDuplicationSrcintf   `json:"srcintf,omitempty"`
}

type SystemSdwanDuplicationDstaddr struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanDuplicationDstaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanDuplicationDstintf struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanDuplicationService struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanDuplicationServiceId struct {
	Id *int64 `json:"id,omitempty"`
}

type SystemSdwanDuplicationSrcaddr struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanDuplicationSrcaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanDuplicationSrcintf struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanFailAlertInterfaces struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanHealthCheck struct {
	AddrMode                   *string                         `json:"addr-mode,omitempty"`
	DetectMode                 *string                         `json:"detect-mode,omitempty"`
	Diffservcode               *string                         `json:"diffservcode,omitempty"`
	DnsMatchIp                 *string                         `json:"dns-match-ip,omitempty"`
	DnsRequestDomain           *string                         `json:"dns-request-domain,omitempty"`
	Failtime                   *int64                          `json:"failtime,omitempty"`
	FtpFile                    *string                         `json:"ftp-file,omitempty"`
	FtpMode                    *string                         `json:"ftp-mode,omitempty"`
	HaPriority                 *int64                          `json:"ha-priority,omitempty"`
	HttpAgent                  *string                         `json:"http-agent,omitempty"`
	HttpGet                    *string                         `json:"http-get,omitempty"`
	HttpMatch                  *string                         `json:"http-match,omitempty"`
	Interval                   *int64                          `json:"interval,omitempty"`
	Members                    []SystemSdwanHealthCheckMembers `json:"members,omitempty"`
	Name                       *string                         `json:"name,omitempty"`
	PacketSize                 *int64                          `json:"packet-size,omitempty"`
	Password                   *string                         `json:"password,omitempty"`
	Port                       *int64                          `json:"port,omitempty"`
	ProbeCount                 *int64                          `json:"probe-count,omitempty"`
	ProbePackets               *string                         `json:"probe-packets,omitempty"`
	ProbeTimeout               *int64                          `json:"probe-timeout,omitempty"`
	Protocol                   *string                         `json:"protocol,omitempty"`
	QualityMeasuredMethod      *string                         `json:"quality-measured-method,omitempty"`
	Recoverytime               *int64                          `json:"recoverytime,omitempty"`
	SecurityMode               *string                         `json:"security-mode,omitempty"`
	Server                     *string                         `json:"server,omitempty"`
	Sla                        []SystemSdwanHealthCheckSla     `json:"sla,omitempty"`
	SlaFailLogPeriod           *int64                          `json:"sla-fail-log-period,omitempty"`
	SlaPassLogPeriod           *int64                          `json:"sla-pass-log-period,omitempty"`
	SystemDns                  *string                         `json:"system-dns,omitempty"`
	ThresholdAlertJitter       *int64                          `json:"threshold-alert-jitter,omitempty"`
	ThresholdAlertLatency      *int64                          `json:"threshold-alert-latency,omitempty"`
	ThresholdAlertPacketloss   *int64                          `json:"threshold-alert-packetloss,omitempty"`
	ThresholdWarningJitter     *int64                          `json:"threshold-warning-jitter,omitempty"`
	ThresholdWarningLatency    *int64                          `json:"threshold-warning-latency,omitempty"`
	ThresholdWarningPacketloss *int64                          `json:"threshold-warning-packetloss,omitempty"`
	UpdateCascadeInterface     *string                         `json:"update-cascade-interface,omitempty"`
	UpdateStaticRoute          *string                         `json:"update-static-route,omitempty"`
	User                       *string                         `json:"user,omitempty"`
}

type SystemSdwanHealthCheckMembers struct {
	SeqNum *int64 `json:"seq-num,omitempty"`
}

type SystemSdwanHealthCheckSla struct {
	Id                  *int64  `json:"id,omitempty"`
	JitterThreshold     *int64  `json:"jitter-threshold,omitempty"`
	LatencyThreshold    *int64  `json:"latency-threshold,omitempty"`
	LinkCostFactor      *string `json:"link-cost-factor,omitempty"`
	PacketlossThreshold *int64  `json:"packetloss-threshold,omitempty"`
}

type SystemSdwanMembers struct {
	Comment                   *string `json:"comment,omitempty"`
	Cost                      *int64  `json:"cost,omitempty"`
	Gateway                   *string `json:"gateway,omitempty"`
	Gateway6                  *string `json:"gateway6,omitempty"`
	IngressSpilloverThreshold *int64  `json:"ingress-spillover-threshold,omitempty"`
	Interface                 *string `json:"interface,omitempty"`
	Priority                  *int64  `json:"priority,omitempty"`
	Priority6                 *int64  `json:"priority6,omitempty"`
	SeqNum                    *int64  `json:"seq-num,omitempty"`
	Source                    *string `json:"source,omitempty"`
	Source6                   *string `json:"source6,omitempty"`
	SpilloverThreshold        *int64  `json:"spillover-threshold,omitempty"`
	Status                    *string `json:"status,omitempty"`
	VolumeRatio               *int64  `json:"volume-ratio,omitempty"`
	Weight                    *int64  `json:"weight,omitempty"`
	Zone                      *string `json:"zone,omitempty"`
}

type SystemSdwanNeighbor struct {
	HealthCheck *string `json:"health-check,omitempty"`
	Ip          *string `json:"ip,omitempty"`
	Member      *int64  `json:"member,omitempty"`
	Mode        *string `json:"mode,omitempty"`
	Role        *string `json:"role,omitempty"`
	SlaId       *int64  `json:"sla-id,omitempty"`
}

type SystemSdwanService struct {
	AddrMode                    *string                                         `json:"addr-mode,omitempty"`
	BandwidthWeight             *int64                                          `json:"bandwidth-weight,omitempty"`
	Default                     *string                                         `json:"default,omitempty"`
	DscpForward                 *string                                         `json:"dscp-forward,omitempty"`
	DscpForwardTag              *string                                         `json:"dscp-forward-tag,omitempty"`
	DscpReverse                 *string                                         `json:"dscp-reverse,omitempty"`
	DscpReverseTag              *string                                         `json:"dscp-reverse-tag,omitempty"`
	Dst                         []SystemSdwanServiceDst                         `json:"dst,omitempty"`
	DstNegate                   *string                                         `json:"dst-negate,omitempty"`
	Dst6                        []SystemSdwanServiceDst6                        `json:"dst6,omitempty"`
	EndPort                     *int64                                          `json:"end-port,omitempty"`
	Gateway                     *string                                         `json:"gateway,omitempty"`
	Groups                      []SystemSdwanServiceGroups                      `json:"groups,omitempty"`
	HashMode                    *string                                         `json:"hash-mode,omitempty"`
	HealthCheck                 []SystemSdwanServiceHealthCheck                 `json:"health-check,omitempty"`
	HoldDownTime                *int64                                          `json:"hold-down-time,omitempty"`
	Id                          *int64                                          `json:"id,omitempty"`
	InputDevice                 []SystemSdwanServiceInputDevice                 `json:"input-device,omitempty"`
	InputDeviceNegate           *string                                         `json:"input-device-negate,omitempty"`
	InternetService             *string                                         `json:"internet-service,omitempty"`
	InternetServiceAppCtrl      []SystemSdwanServiceInternetServiceAppCtrl      `json:"internet-service-app-ctrl,omitempty"`
	InternetServiceAppCtrlGroup []SystemSdwanServiceInternetServiceAppCtrlGroup `json:"internet-service-app-ctrl-group,omitempty"`
	InternetServiceCustom       []SystemSdwanServiceInternetServiceCustom       `json:"internet-service-custom,omitempty"`
	InternetServiceCustomGroup  []SystemSdwanServiceInternetServiceCustomGroup  `json:"internet-service-custom-group,omitempty"`
	InternetServiceGroup        []SystemSdwanServiceInternetServiceGroup        `json:"internet-service-group,omitempty"`
	InternetServiceName         []SystemSdwanServiceInternetServiceName         `json:"internet-service-name,omitempty"`
	JitterWeight                *int64                                          `json:"jitter-weight,omitempty"`
	LatencyWeight               *int64                                          `json:"latency-weight,omitempty"`
	LinkCostFactor              *string                                         `json:"link-cost-factor,omitempty"`
	LinkCostThreshold           *int64                                          `json:"link-cost-threshold,omitempty"`
	MinimumSlaMeetMembers       *int64                                          `json:"minimum-sla-meet-members,omitempty"`
	Mode                        *string                                         `json:"mode,omitempty"`
	Name                        *string                                         `json:"name,omitempty"`
	PacketLossWeight            *int64                                          `json:"packet-loss-weight,omitempty"`
	PassiveMeasurement          *string                                         `json:"passive-measurement,omitempty"`
	PriorityMembers             []SystemSdwanServicePriorityMembers             `json:"priority-members,omitempty"`
	PriorityZone                []SystemSdwanServicePriorityZone                `json:"priority-zone,omitempty"`
	Protocol                    *int64                                          `json:"protocol,omitempty"`
	QualityLink                 *int64                                          `json:"quality-link,omitempty"`
	Role                        *string                                         `json:"role,omitempty"`
	RouteTag                    *int64                                          `json:"route-tag,omitempty"`
	Sla                         []SystemSdwanServiceSla                         `json:"sla,omitempty"`
	SlaCompareMethod            *string                                         `json:"sla-compare-method,omitempty"`
	Src                         []SystemSdwanServiceSrc                         `json:"src,omitempty"`
	SrcNegate                   *string                                         `json:"src-negate,omitempty"`
	Src6                        []SystemSdwanServiceSrc6                        `json:"src6,omitempty"`
	StandaloneAction            *string                                         `json:"standalone-action,omitempty"`
	StartPort                   *int64                                          `json:"start-port,omitempty"`
	Status                      *string                                         `json:"status,omitempty"`
	TieBreak                    *string                                         `json:"tie-break,omitempty"`
	Tos                         *string                                         `json:"tos,omitempty"`
	TosMask                     *string                                         `json:"tos-mask,omitempty"`
	UseShortcutSla              *string                                         `json:"use-shortcut-sla,omitempty"`
	Users                       []SystemSdwanServiceUsers                       `json:"users,omitempty"`
}

type SystemSdwanServiceDst struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanServiceDst6 struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanServiceGroups struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanServiceHealthCheck struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanServiceInputDevice struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanServiceInternetServiceAppCtrl struct {
	Id *int64 `json:"id,omitempty"`
}

type SystemSdwanServiceInternetServiceAppCtrlGroup struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanServiceInternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanServiceInternetServiceCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanServiceInternetServiceGroup struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanServiceInternetServiceName struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanServicePriorityMembers struct {
	SeqNum *int64 `json:"seq-num,omitempty"`
}

type SystemSdwanServicePriorityZone struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanServiceSla struct {
	HealthCheck *string `json:"health-check,omitempty"`
	Id          *int64  `json:"id,omitempty"`
}

type SystemSdwanServiceSrc struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanServiceSrc6 struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanServiceUsers struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdwanZone struct {
	Name               *string `json:"name,omitempty"`
	ServiceSlaTieBreak *string `json:"service-sla-tie-break,omitempty"`
}
