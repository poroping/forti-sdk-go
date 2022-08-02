package models

const LogNpuServerPath = "log/npu-server/"

type LogNpuServer struct {
	LogProcessing  *string                    `json:"log-processing,omitempty"`
	LogProcessor   *string                    `json:"log-processor,omitempty"`
	NetflowVer     *string                    `json:"netflow-ver,omitempty"`
	ServerGroup    *[]LogNpuServerServerGroup `json:"server-group,omitempty"`
	ServerInfo     *[]LogNpuServerServerInfo  `json:"server-info,omitempty"`
	SyslogFacility *int64                     `json:"syslog-facility,omitempty"`
	SyslogSeverity *int64                     `json:"syslog-severity,omitempty"`
}

const LogNpuServerServerGroupPath = "log/npu-server/server-group/"

type LogNpuServerServerGroup struct {
	GroupName     *string `json:"group-name,omitempty"`
	LogFormat     *string `json:"log-format,omitempty"`
	LogMode       *string `json:"log-mode,omitempty"`
	LogTxMode     *string `json:"log-tx-mode,omitempty"`
	ServerNumber  *int64  `json:"server-number,omitempty"`
	ServerStartId *int64  `json:"server-start-id,omitempty"`
	SwLogFlags    *int64  `json:"sw-log-flags,omitempty"`
}

const LogNpuServerServerInfoPath = "log/npu-server/server-info/"

type LogNpuServerServerInfo struct {
	DestPort          *int64  `json:"dest-port,omitempty"`
	Id                *int64  `json:"id,omitempty"`
	IpFamily          *string `json:"ip-family,omitempty"`
	Ipv4Server        *string `json:"ipv4-server,omitempty"`
	Ipv6Server        *string `json:"ipv6-server,omitempty"`
	SourcePort        *int64  `json:"source-port,omitempty"`
	TemplateTxTimeout *int64  `json:"template-tx-timeout,omitempty"`
	Vdom              *string `json:"vdom,omitempty"`
}
