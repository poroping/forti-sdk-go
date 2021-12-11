package models

const WirelessControllerUtmProfilePath = "wireless-controller/utm-profile/"

type WirelessControllerUtmProfile struct {
	AntivirusProfile      *string `json:"antivirus-profile,omitempty"`
	ApplicationList       *string `json:"application-list,omitempty"`
	Comment               *string `json:"comment,omitempty"`
	IpsSensor             *string `json:"ips-sensor,omitempty"`
	Name                  *string `json:"name,omitempty"`
	ScanBotnetConnections *string `json:"scan-botnet-connections,omitempty"`
	UtmLog                *string `json:"utm-log,omitempty"`
	WebfilterProfile      *string `json:"webfilter-profile,omitempty"`
}
