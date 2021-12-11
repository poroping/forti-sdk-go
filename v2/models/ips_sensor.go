package models

const IpsSensorPath = "ips/sensor/"

type IpsSensor struct {
	BlockMaliciousUrl     *string            `json:"block-malicious-url,omitempty"`
	Comment               *string            `json:"comment,omitempty"`
	Entries               []IpsSensorEntries `json:"entries,omitempty"`
	ExtendedLog           *string            `json:"extended-log,omitempty"`
	Name                  *string            `json:"name,omitempty"`
	ReplacemsgGroup       *string            `json:"replacemsg-group,omitempty"`
	ScanBotnetConnections *string            `json:"scan-botnet-connections,omitempty"`
}

type IpsSensorEntries struct {
	Action           *string                    `json:"action,omitempty"`
	Application      *string                    `json:"application,omitempty"`
	Cve              []IpsSensorEntriesCve      `json:"cve,omitempty"`
	ExemptIp         []IpsSensorEntriesExemptIp `json:"exempt-ip,omitempty"`
	Id               *int64                     `json:"id,omitempty"`
	Location         *string                    `json:"location,omitempty"`
	Log              *string                    `json:"log,omitempty"`
	LogAttackContext *string                    `json:"log-attack-context,omitempty"`
	LogPacket        *string                    `json:"log-packet,omitempty"`
	Os               *string                    `json:"os,omitempty"`
	Protocol         *string                    `json:"protocol,omitempty"`
	Quarantine       *string                    `json:"quarantine,omitempty"`
	QuarantineExpiry *string                    `json:"quarantine-expiry,omitempty"`
	QuarantineLog    *string                    `json:"quarantine-log,omitempty"`
	RateCount        *int64                     `json:"rate-count,omitempty"`
	RateDuration     *int64                     `json:"rate-duration,omitempty"`
	RateMode         *string                    `json:"rate-mode,omitempty"`
	RateTrack        *string                    `json:"rate-track,omitempty"`
	Rule             []IpsSensorEntriesRule     `json:"rule,omitempty"`
	Severity         *string                    `json:"severity,omitempty"`
	Status           *string                    `json:"status,omitempty"`
}

type IpsSensorEntriesCve struct {
	CveEntry *string `json:"cve-entry,omitempty"`
}

type IpsSensorEntriesExemptIp struct {
	DstIp *string `json:"dst-ip,omitempty"`
	Id    *int64  `json:"id,omitempty"`
	SrcIp *string `json:"src-ip,omitempty"`
}

type IpsSensorEntriesRule struct {
	Id *int64 `json:"id,omitempty"`
}
