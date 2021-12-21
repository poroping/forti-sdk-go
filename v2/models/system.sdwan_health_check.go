package models

const SystemsdwanHealthCheckPath = "system.sdwan/health-check/"

type SystemsdwanHealthCheck struct {
	AddrMode                   *string                          `json:"addr-mode,omitempty"`
	DetectMode                 *string                          `json:"detect-mode,omitempty"`
	Diffservcode               *string                          `json:"diffservcode,omitempty"`
	DnsMatchIp                 *string                          `json:"dns-match-ip,omitempty"`
	DnsRequestDomain           *string                          `json:"dns-request-domain,omitempty"`
	Failtime                   *float64                         `json:"failtime,omitempty"`
	FtpFile                    *string                          `json:"ftp-file,omitempty"`
	FtpMode                    *string                          `json:"ftp-mode,omitempty"`
	HaPriority                 *float64                         `json:"ha-priority,omitempty"`
	HttpAgent                  *string                          `json:"http-agent,omitempty"`
	HttpGet                    *string                          `json:"http-get,omitempty"`
	HttpMatch                  *string                          `json:"http-match,omitempty"`
	Interval                   *float64                         `json:"interval,omitempty"`
	Members                    *[]SystemsdwanHealthCheckMembers `json:"members,omitempty"`
	Name                       *string                          `json:"name,omitempty"`
	PacketSize                 *float64                         `json:"packet-size,omitempty"`
	Password                   *string                          `json:"password,omitempty"`
	Port                       *float64                         `json:"port,omitempty"`
	ProbeCount                 *float64                         `json:"probe-count,omitempty"`
	ProbePackets               *string                          `json:"probe-packets,omitempty"`
	ProbeTimeout               *float64                         `json:"probe-timeout,omitempty"`
	Protocol                   *string                          `json:"protocol,omitempty"`
	QualityMeasuredMethod      *string                          `json:"quality-measured-method,omitempty"`
	Recoverytime               *float64                         `json:"recoverytime,omitempty"`
	SecurityMode               *string                          `json:"security-mode,omitempty"`
	Server                     *string                          `json:"server,omitempty"`
	Sla                        *[]SystemsdwanHealthCheckSla     `json:"sla,omitempty"`
	SlaFailLogPeriod           *float64                         `json:"sla-fail-log-period,omitempty"`
	SlaPassLogPeriod           *float64                         `json:"sla-pass-log-period,omitempty"`
	SystemDns                  *string                          `json:"system-dns,omitempty"`
	ThresholdAlertJitter       *float64                         `json:"threshold-alert-jitter,omitempty"`
	ThresholdAlertLatency      *float64                         `json:"threshold-alert-latency,omitempty"`
	ThresholdAlertPacketloss   *float64                         `json:"threshold-alert-packetloss,omitempty"`
	ThresholdWarningJitter     *float64                         `json:"threshold-warning-jitter,omitempty"`
	ThresholdWarningLatency    *float64                         `json:"threshold-warning-latency,omitempty"`
	ThresholdWarningPacketloss *float64                         `json:"threshold-warning-packetloss,omitempty"`
	UpdateCascadeInterface     *string                          `json:"update-cascade-interface,omitempty"`
	UpdateStaticRoute          *string                          `json:"update-static-route,omitempty"`
	User                       *string                          `json:"user,omitempty"`
}

type SystemsdwanHealthCheckMembers struct {
	SeqNum *float64 `json:"seq-num,omitempty"`
}

type SystemsdwanHealthCheckSla struct {
	Id                  *float64 `json:"id,omitempty"`
	JitterThreshold     *float64 `json:"jitter-threshold,omitempty"`
	LatencyThreshold    *float64 `json:"latency-threshold,omitempty"`
	LinkCostFactor      *string  `json:"link-cost-factor,omitempty"`
	PacketlossThreshold *float64 `json:"packetloss-threshold,omitempty"`
}
