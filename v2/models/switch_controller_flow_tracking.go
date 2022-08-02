package models

const SwitchControllerFlowTrackingPath = "switch-controller/flow-tracking/"

type SwitchControllerFlowTracking struct {
	Aggregates       *[]SwitchControllerFlowTrackingAggregates `json:"aggregates,omitempty"`
	CollectorIp      *string                                   `json:"collector-ip,omitempty"`
	CollectorPort    *int64                                    `json:"collector-port,omitempty"`
	Format           *string                                   `json:"format,omitempty"`
	Level            *string                                   `json:"level,omitempty"`
	MaxExportPktSize *int64                                    `json:"max-export-pkt-size,omitempty"`
	SampleMode       *string                                   `json:"sample-mode,omitempty"`
	SampleRate       *int64                                    `json:"sample-rate,omitempty"`
	TimeoutGeneral   *int64                                    `json:"timeout-general,omitempty"`
	TimeoutIcmp      *int64                                    `json:"timeout-icmp,omitempty"`
	TimeoutMax       *int64                                    `json:"timeout-max,omitempty"`
	TimeoutTcp       *int64                                    `json:"timeout-tcp,omitempty"`
	TimeoutTcpFin    *int64                                    `json:"timeout-tcp-fin,omitempty"`
	TimeoutTcpRst    *int64                                    `json:"timeout-tcp-rst,omitempty"`
	TimeoutUdp       *int64                                    `json:"timeout-udp,omitempty"`
	Transport        *string                                   `json:"transport,omitempty"`
}

const SwitchControllerFlowTrackingAggregatesPath = "switch-controller/flow-tracking/aggregates/"

type SwitchControllerFlowTrackingAggregates struct {
	Id *int64  `json:"id,omitempty"`
	Ip *string `json:"ip,omitempty"`
}
