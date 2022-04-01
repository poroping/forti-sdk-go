package models

const FirewallSnifferPath = "firewall/sniffer/"

type FirewallSniffer struct {
	Anomaly                  *[]FirewallSnifferAnomaly      `json:"anomaly,omitempty"`
	ApplicationList          *string                        `json:"application-list,omitempty"`
	ApplicationListStatus    *string                        `json:"application-list-status,omitempty"`
	AvProfile                *string                        `json:"av-profile,omitempty"`
	AvProfileStatus          *string                        `json:"av-profile-status,omitempty"`
	DlpProfile               *string                        `json:"dlp-profile,omitempty"`
	DlpProfileStatus         *string                        `json:"dlp-profile-status,omitempty"`
	DlpSensor                *string                        `json:"dlp-sensor,omitempty"`
	DlpSensorStatus          *string                        `json:"dlp-sensor-status,omitempty"`
	Dsri                     *string                        `json:"dsri,omitempty"`
	EmailfilterProfile       *string                        `json:"emailfilter-profile,omitempty"`
	EmailfilterProfileStatus *string                        `json:"emailfilter-profile-status,omitempty"`
	FileFilterProfile        *string                        `json:"file-filter-profile,omitempty"`
	FileFilterProfileStatus  *string                        `json:"file-filter-profile-status,omitempty"`
	Host                     *string                        `json:"host,omitempty"`
	Id                       *int64                         `json:"id,omitempty"`
	Interface                *string                        `json:"interface,omitempty"`
	IpThreatfeed             *[]FirewallSnifferIpThreatfeed `json:"ip-threatfeed,omitempty"`
	IpThreatfeedStatus       *string                        `json:"ip-threatfeed-status,omitempty"`
	IpsDosStatus             *string                        `json:"ips-dos-status,omitempty"`
	IpsSensor                *string                        `json:"ips-sensor,omitempty"`
	IpsSensorStatus          *string                        `json:"ips-sensor-status,omitempty"`
	Ipv6                     *string                        `json:"ipv6,omitempty"`
	Logtraffic               *string                        `json:"logtraffic,omitempty"`
	MaxPacketCount           *int64                         `json:"max-packet-count,omitempty"`
	NonIp                    *string                        `json:"non-ip,omitempty"`
	Port                     *string                        `json:"port,omitempty"`
	Protocol                 *string                        `json:"protocol,omitempty"`
	Status                   *string                        `json:"status,omitempty"`
	Vlan                     *string                        `json:"vlan,omitempty"`
	WebfilterProfile         *string                        `json:"webfilter-profile,omitempty"`
	WebfilterProfileStatus   *string                        `json:"webfilter-profile-status,omitempty"`
}

type FirewallSnifferAnomaly struct {
	Action           *string `json:"action,omitempty"`
	Log              *string `json:"log,omitempty"`
	Name             *string `json:"name,omitempty"`
	Quarantine       *string `json:"quarantine,omitempty"`
	QuarantineExpiry *string `json:"quarantine-expiry,omitempty"`
	QuarantineLog    *string `json:"quarantine-log,omitempty"`
	Status           *string `json:"status,omitempty"`
	Threshold        *int64  `json:"threshold,omitempty"`
	Thresholddefault *int64  `json:"threshold(default),omitempty"`
}

type FirewallSnifferIpThreatfeed struct {
	Name *string `json:"name,omitempty"`
}
