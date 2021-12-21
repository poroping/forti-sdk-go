package models

const SystemsdwanServicePath = "system.sdwan/service/"

type SystemsdwanService struct {
	AddrMode                    *string                                          `json:"addr-mode,omitempty"`
	BandwidthWeight             *int64                                           `json:"bandwidth-weight,omitempty"`
	Default                     *string                                          `json:"default,omitempty"`
	DscpForward                 *string                                          `json:"dscp-forward,omitempty"`
	DscpForwardTag              *string                                          `json:"dscp-forward-tag,omitempty"`
	DscpReverse                 *string                                          `json:"dscp-reverse,omitempty"`
	DscpReverseTag              *string                                          `json:"dscp-reverse-tag,omitempty"`
	Dst                         *[]SystemsdwanServiceDst                         `json:"dst,omitempty"`
	DstNegate                   *string                                          `json:"dst-negate,omitempty"`
	Dst6                        *[]SystemsdwanServiceDst6                        `json:"dst6,omitempty"`
	EndPort                     *int64                                           `json:"end-port,omitempty"`
	Gateway                     *string                                          `json:"gateway,omitempty"`
	Groups                      *[]SystemsdwanServiceGroups                      `json:"groups,omitempty"`
	HashMode                    *string                                          `json:"hash-mode,omitempty"`
	HealthCheck                 *[]SystemsdwanServiceHealthCheck                 `json:"health-check,omitempty"`
	HoldDownTime                *int64                                           `json:"hold-down-time,omitempty"`
	Fosid                       *int64                                           `json:"fosid,omitempty"`
	InputDevice                 *[]SystemsdwanServiceInputDevice                 `json:"input-device,omitempty"`
	InputDeviceNegate           *string                                          `json:"input-device-negate,omitempty"`
	InternetService             *string                                          `json:"internet-service,omitempty"`
	InternetServiceAppCtrl      *[]SystemsdwanServiceInternetServiceAppCtrl      `json:"internet-service-app-ctrl,omitempty"`
	InternetServiceAppCtrlGroup *[]SystemsdwanServiceInternetServiceAppCtrlGroup `json:"internet-service-app-ctrl-group,omitempty"`
	InternetServiceCustom       *[]SystemsdwanServiceInternetServiceCustom       `json:"internet-service-custom,omitempty"`
	InternetServiceCustomGroup  *[]SystemsdwanServiceInternetServiceCustomGroup  `json:"internet-service-custom-group,omitempty"`
	InternetServiceGroup        *[]SystemsdwanServiceInternetServiceGroup        `json:"internet-service-group,omitempty"`
	InternetServiceName         *[]SystemsdwanServiceInternetServiceName         `json:"internet-service-name,omitempty"`
	JitterWeight                *int64                                           `json:"jitter-weight,omitempty"`
	LatencyWeight               *int64                                           `json:"latency-weight,omitempty"`
	LinkCostFactor              *string                                          `json:"link-cost-factor,omitempty"`
	LinkCostThreshold           *int64                                           `json:"link-cost-threshold,omitempty"`
	MinimumSlaMeetMembers       *int64                                           `json:"minimum-sla-meet-members,omitempty"`
	Mode                        *string                                          `json:"mode,omitempty"`
	Name                        *string                                          `json:"name,omitempty"`
	PacketLossWeight            *int64                                           `json:"packet-loss-weight,omitempty"`
	PassiveMeasurement          *string                                          `json:"passive-measurement,omitempty"`
	PriorityMembers             *[]SystemsdwanServicePriorityMembers             `json:"priority-members,omitempty"`
	PriorityZone                *[]SystemsdwanServicePriorityZone                `json:"priority-zone,omitempty"`
	Protocol                    *int64                                           `json:"protocol,omitempty"`
	QualityLink                 *int64                                           `json:"quality-link,omitempty"`
	Role                        *string                                          `json:"role,omitempty"`
	RouteTag                    *int64                                           `json:"route-tag,omitempty"`
	Sla                         *[]SystemsdwanServiceSla                         `json:"sla,omitempty"`
	SlaCompareMethod            *string                                          `json:"sla-compare-method,omitempty"`
	Src                         *[]SystemsdwanServiceSrc                         `json:"src,omitempty"`
	SrcNegate                   *string                                          `json:"src-negate,omitempty"`
	Src6                        *[]SystemsdwanServiceSrc6                        `json:"src6,omitempty"`
	StandaloneAction            *string                                          `json:"standalone-action,omitempty"`
	StartPort                   *int64                                           `json:"start-port,omitempty"`
	Status                      *string                                          `json:"status,omitempty"`
	TieBreak                    *string                                          `json:"tie-break,omitempty"`
	Tos                         *string                                          `json:"tos,omitempty"`
	TosMask                     *string                                          `json:"tos-mask,omitempty"`
	UseShortcutSla              *string                                          `json:"use-shortcut-sla,omitempty"`
	Users                       *[]SystemsdwanServiceUsers                       `json:"users,omitempty"`
}

type SystemsdwanServiceDst struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanServiceDst6 struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanServiceGroups struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanServiceHealthCheck struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanServiceInputDevice struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanServiceInternetServiceAppCtrl struct {
	Id *int64 `json:"id,omitempty"`
}

type SystemsdwanServiceInternetServiceAppCtrlGroup struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanServiceInternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanServiceInternetServiceCustomGroup struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanServiceInternetServiceGroup struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanServiceInternetServiceName struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanServicePriorityMembers struct {
	SeqNum *int64 `json:"seq-num,omitempty"`
}

type SystemsdwanServicePriorityZone struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanServiceSla struct {
	HealthCheck *string `json:"health-check,omitempty"`
	Id          *int64  `json:"id,omitempty"`
}

type SystemsdwanServiceSrc struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanServiceSrc6 struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanServiceUsers struct {
	Name *string `json:"name,omitempty"`
}
