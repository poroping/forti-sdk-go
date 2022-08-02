package models

const SystemLinkMonitorPath = "system/link-monitor/"

type SystemLinkMonitor struct {
	AddrMode               *string                        `json:"addr-mode,omitempty"`
	ClassId                *int64                         `json:"class-id,omitempty"`
	Diffservcode           *string                        `json:"diffservcode,omitempty"`
	FailWeight             *int64                         `json:"fail-weight,omitempty"`
	Failtime               *int64                         `json:"failtime,omitempty"`
	GatewayIp              *string                        `json:"gateway-ip,omitempty"`
	GatewayIp6             *string                        `json:"gateway-ip6,omitempty"`
	HaPriority             *int64                         `json:"ha-priority,omitempty"`
	HttpAgent              *string                        `json:"http-agent,omitempty"`
	HttpGet                *string                        `json:"http-get,omitempty"`
	HttpMatch              *string                        `json:"http-match,omitempty"`
	Interval               *int64                         `json:"interval,omitempty"`
	Name                   *string                        `json:"name,omitempty"`
	PacketSize             *int64                         `json:"packet-size,omitempty"`
	Password               *string                        `json:"password,omitempty"`
	Port                   *int64                         `json:"port,omitempty"`
	ProbeCount             *int64                         `json:"probe-count,omitempty"`
	ProbeTimeout           *int64                         `json:"probe-timeout,omitempty"`
	Protocol               *string                        `json:"protocol,omitempty"`
	Recoverytime           *int64                         `json:"recoverytime,omitempty"`
	Route                  *[]SystemLinkMonitorRoute      `json:"route,omitempty"`
	SecurityMode           *string                        `json:"security-mode,omitempty"`
	Server                 *[]SystemLinkMonitorServer     `json:"server,omitempty"`
	ServerConfig           *string                        `json:"server-config,omitempty"`
	ServerList             *[]SystemLinkMonitorServerList `json:"server-list,omitempty"`
	ServiceDetection       *string                        `json:"service-detection,omitempty"`
	SourceIp               *string                        `json:"source-ip,omitempty"`
	SourceIp6              *string                        `json:"source-ip6,omitempty"`
	Srcintf                *string                        `json:"srcintf,omitempty"`
	Status                 *string                        `json:"status,omitempty"`
	UpdateCascadeInterface *string                        `json:"update-cascade-interface,omitempty"`
	UpdatePolicyRoute      *string                        `json:"update-policy-route,omitempty"`
	UpdateStaticRoute      *string                        `json:"update-static-route,omitempty"`
}

const SystemLinkMonitorRoutePath = "system/link-monitor/route/"

type SystemLinkMonitorRoute struct {
	Subnet *string `json:"subnet,omitempty"`
}

const SystemLinkMonitorServerPath = "system/link-monitor/server/"

type SystemLinkMonitorServer struct {
	Address *string `json:"address,omitempty"`
}

const SystemLinkMonitorServerListPath = "system/link-monitor/server-list/"

type SystemLinkMonitorServerList struct {
	Dst      *string `json:"dst,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Port     *int64  `json:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Weight   *int64  `json:"weight,omitempty"`
}
