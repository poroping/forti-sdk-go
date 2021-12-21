package models

const IpsSensorPath = "ips/sensor/"

type IpsSensor struct {
	BlockMaliciousUrl     *string              `json:"block-malicious-url,omitempty"`
	Comment               *string              `json:"comment,omitempty"`
	Entries               *[]IpsSensorEntries  `json:"entries,omitempty"`
	ExtendedLog           *string              `json:"extended-log,omitempty"`
	Filter                *[]IpsSensorFilter   `json:"filter,omitempty"`
	Name                  *string              `json:"name,omitempty"`
	Override              *[]IpsSensorOverride `json:"override,omitempty"`
	ReplacemsgGroup       *string              `json:"replacemsg-group,omitempty"`
	ScanBotnetConnections *string              `json:"scan-botnet-connections,omitempty"`
}

type IpsSensorEntries struct {
	Action           *string                     `json:"action,omitempty"`
	Application      *string                     `json:"application,omitempty"`
	Cve              *[]IpsSensorEntriesCve      `json:"cve,omitempty"`
	ExemptIp         *[]IpsSensorEntriesExemptIp `json:"exempt-ip,omitempty"`
	Id               *float64                    `json:"id,omitempty"`
	Location         *string                     `json:"location,omitempty"`
	Log              *string                     `json:"log,omitempty"`
	LogAttackContext *string                     `json:"log-attack-context,omitempty"`
	LogPacket        *string                     `json:"log-packet,omitempty"`
	Os               *string                     `json:"os,omitempty"`
	Protocol         *string                     `json:"protocol,omitempty"`
	Quarantine       *string                     `json:"quarantine,omitempty"`
	QuarantineExpiry *string                     `json:"quarantine-expiry,omitempty"`
	QuarantineLog    *string                     `json:"quarantine-log,omitempty"`
	RateCount        *float64                    `json:"rate-count,omitempty"`
	RateDuration     *float64                    `json:"rate-duration,omitempty"`
	RateMode         *string                     `json:"rate-mode,omitempty"`
	RateTrack        *string                     `json:"rate-track,omitempty"`
	Rule             *[]IpsSensorEntriesRule     `json:"rule,omitempty"`
	Severity         *string                     `json:"severity,omitempty"`
	Status           *string                     `json:"status,omitempty"`
}

type IpsSensorEntriesCve struct {
	CveEntry *string `json:"cve-entry,omitempty"`
}

type IpsSensorEntriesExemptIp struct {
	DstIp *string  `json:"dst-ip,omitempty"`
	Id    *float64 `json:"id,omitempty"`
	SrcIp *string  `json:"src-ip,omitempty"`
}

type IpsSensorEntriesRule struct {
	Id *float64 `json:"id,omitempty"`
}

type IpsSensorFilter struct {
	Action           *string  `json:"action,omitempty"`
	Application      *string  `json:"application,omitempty"`
	Location         *string  `json:"location,omitempty"`
	Log              *string  `json:"log,omitempty"`
	LogPacket        *string  `json:"log-packet,omitempty"`
	Name             *string  `json:"name,omitempty"`
	Os               *string  `json:"os,omitempty"`
	Protocol         *string  `json:"protocol,omitempty"`
	Quarantine       *string  `json:"quarantine,omitempty"`
	QuarantineExpiry *float64 `json:"quarantine-expiry,omitempty"`
	QuarantineLog    *string  `json:"quarantine-log,omitempty"`
	Severity         *string  `json:"severity,omitempty"`
	Status           *string  `json:"status,omitempty"`
}

type IpsSensorOverride struct {
	Action           *string                      `json:"action,omitempty"`
	ExemptIp         *[]IpsSensorOverrideExemptIp `json:"exempt-ip,omitempty"`
	Log              *string                      `json:"log,omitempty"`
	LogPacket        *string                      `json:"log-packet,omitempty"`
	Quarantine       *string                      `json:"quarantine,omitempty"`
	QuarantineExpiry *float64                     `json:"quarantine-expiry,omitempty"`
	QuarantineLog    *string                      `json:"quarantine-log,omitempty"`
	RuleId           *float64                     `json:"rule-id,omitempty"`
	Status           *string                      `json:"status,omitempty"`
}

type IpsSensorOverrideExemptIp struct {
	DstIp *string  `json:"dst-ip,omitempty"`
	Id    *float64 `json:"id,omitempty"`
	SrcIp *string  `json:"src-ip,omitempty"`
}
