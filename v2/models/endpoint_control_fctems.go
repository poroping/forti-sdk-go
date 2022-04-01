package models

const EndpointControlFctemsPath = "endpoint-control/fctems/"

type EndpointControlFctems struct {
	AdminPassword                  *string `json:"admin-password,omitempty"`
	AdminUsername                  *string `json:"admin-username,omitempty"`
	CallTimeout                    *int64  `json:"call-timeout,omitempty"`
	Capabilities                   *string `json:"capabilities,omitempty"`
	Certificate                    *string `json:"certificate,omitempty"`
	CloudServerType                *string `json:"cloud-server-type,omitempty"`
	FortinetoneCloudAuthentication *string `json:"fortinetone-cloud-authentication,omitempty"`
	HttpsPort                      *int64  `json:"https-port,omitempty"`
	Name                           *string `json:"name,omitempty"`
	OutOfSyncThreshold             *int64  `json:"out-of-sync-threshold,omitempty"`
	PreserveSslSession             *string `json:"preserve-ssl-session,omitempty"`
	PullAvatars                    *string `json:"pull-avatars,omitempty"`
	PullMalwareHash                *string `json:"pull-malware-hash,omitempty"`
	PullSysinfo                    *string `json:"pull-sysinfo,omitempty"`
	PullTags                       *string `json:"pull-tags,omitempty"`
	PullVulnerabilities            *string `json:"pull-vulnerabilities,omitempty"`
	SerialNumber                   *string `json:"serial-number,omitempty"`
	Server                         *string `json:"server,omitempty"`
	SourceIp                       *string `json:"source-ip,omitempty"`
	StatusCheckInterval            *int64  `json:"status-check-interval,omitempty"`
	WebsocketOverride              *string `json:"websocket-override,omitempty"`
}
