package models

const LogGuiDisplayPath = "log/gui-display/"

type LogGuiDisplay struct {
	FortiviewUnscannedApps *string `json:"fortiview-unscanned-apps,omitempty"`
	ResolveApps            *string `json:"resolve-apps,omitempty"`
	ResolveHosts           *string `json:"resolve-hosts,omitempty"`
}
