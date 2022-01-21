package models

const LogThreatWeightPath = "log/threat-weight/"

type LogThreatWeight struct {
	Application              *[]LogThreatWeightApplication `json:"application,omitempty"`
	BlockedConnection        *string                       `json:"blocked-connection,omitempty"`
	BotnetConnectionDetected *string                       `json:"botnet-connection-detected,omitempty"`
	FailedConnection         *string                       `json:"failed-connection,omitempty"`
	Geolocation              *[]LogThreatWeightGeolocation `json:"geolocation,omitempty"`
	Ips                      *[]LogThreatWeightIps         `json:"ips,omitempty"`
	Level                    *[]LogThreatWeightLevel       `json:"level,omitempty"`
	Malware                  *[]LogThreatWeightMalware     `json:"malware,omitempty"`
	Status                   *string                       `json:"status,omitempty"`
	UrlBlockDetected         *string                       `json:"url-block-detected,omitempty"`
	Web                      *[]LogThreatWeightWeb         `json:"web,omitempty"`
}

type LogThreatWeightApplication struct {
	Category *int64  `json:"category,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Level    *string `json:"level,omitempty"`
}

type LogThreatWeightGeolocation struct {
	Country *string `json:"country,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	Level   *string `json:"level,omitempty"`
}

type LogThreatWeightIps struct {
	CriticalSeverity *string `json:"critical-severity,omitempty"`
	HighSeverity     *string `json:"high-severity,omitempty"`
	InfoSeverity     *string `json:"info-severity,omitempty"`
	LowSeverity      *string `json:"low-severity,omitempty"`
	MediumSeverity   *string `json:"medium-severity,omitempty"`
}

type LogThreatWeightLevel struct {
	Critical *int64 `json:"critical,omitempty"`
	High     *int64 `json:"high,omitempty"`
	Low      *int64 `json:"low,omitempty"`
	Medium   *int64 `json:"medium,omitempty"`
}

type LogThreatWeightMalware struct {
	CommandBlocked          *string `json:"command-blocked,omitempty"`
	ContentDisarm           *string `json:"content-disarm,omitempty"`
	EmsThreatFeed           *string `json:"ems-threat-feed,omitempty"`
	FileBlocked             *string `json:"file-blocked,omitempty"`
	Fortiai                 *string `json:"fortiai,omitempty"`
	FsaHighRisk             *string `json:"fsa-high-risk,omitempty"`
	FsaMalicious            *string `json:"fsa-malicious,omitempty"`
	FsaMediumRisk           *string `json:"fsa-medium-risk,omitempty"`
	MalwareList             *string `json:"malware-list,omitempty"`
	Mimefragmented          *string `json:"mimefragmented,omitempty"`
	Oversized               *string `json:"oversized,omitempty"`
	SwitchProto             *string `json:"switch-proto,omitempty"`
	VirusFileTypeExecutable *string `json:"virus-file-type-executable,omitempty"`
	VirusInfected           *string `json:"virus-infected,omitempty"`
	VirusOutbreakPrevention *string `json:"virus-outbreak-prevention,omitempty"`
	VirusScanError          *string `json:"virus-scan-error,omitempty"`
}

type LogThreatWeightWeb struct {
	Category *int64  `json:"category,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Level    *string `json:"level,omitempty"`
}
