package models

const FirewallAccessProxySshClientCertPath = "firewall/access-proxy-ssh-client-cert/"

type FirewallAccessProxySshClientCert struct {
	AuthCa                *string                                          `json:"auth-ca,omitempty"`
	CertExtension         *[]FirewallAccessProxySshClientCertCertExtension `json:"cert-extension,omitempty"`
	Name                  *string                                          `json:"name,omitempty"`
	PermitAgentForwarding *string                                          `json:"permit-agent-forwarding,omitempty"`
	PermitPortForwarding  *string                                          `json:"permit-port-forwarding,omitempty"`
	PermitPty             *string                                          `json:"permit-pty,omitempty"`
	PermitUserRc          *string                                          `json:"permit-user-rc,omitempty"`
	PermitX11Forwarding   *string                                          `json:"permit-x11-forwarding,omitempty"`
	SourceAddress         *string                                          `json:"source-address,omitempty"`
}

const FirewallAccessProxySshClientCertCertExtensionPath = "firewall/access-proxy-ssh-client-cert/cert-extension/"

type FirewallAccessProxySshClientCertCertExtension struct {
	Critical *string `json:"critical,omitempty"`
	Data     *string `json:"data,omitempty"`
	Name     *string `json:"name,omitempty"`
	Type     *string `json:"type,omitempty"`
}
