package models

const EndpointControlFctemsPath = "endpoint-control/fctems/"

type EndpointControlFctems struct {
	CallTimeout                    *int64  `json:"call-timeout,omitempty"`
	Capabilities                   *string `json:"capabilities,omitempty"`
	CloudServerType                *string `json:"cloud-server-type,omitempty"`
	FortinetoneCloudAuthentication *string `json:"fortinetone-cloud-authentication,omitempty"`
	HttpsPort                      *int64  `json:"https-port,omitempty"`
	Name                           *string `json:"name,omitempty"`
	PreserveSslSession             *string `json:"preserve-ssl-session,omitempty"`
	PullAvatars                    *string `json:"pull-avatars,omitempty"`
	PullMalwareHash                *string `json:"pull-malware-hash,omitempty"`
	PullSysinfo                    *string `json:"pull-sysinfo,omitempty"`
	PullTags                       *string `json:"pull-tags,omitempty"`
	PullVulnerabilities            *string `json:"pull-vulnerabilities,omitempty"`
	Server                         *string `json:"server,omitempty"`
	SourceIp                       *string `json:"source-ip,omitempty"`
	WebsocketOverride              *string `json:"websocket-override,omitempty"`
}
