package models

const FirewallProfileProtocolOptionsPath = "firewall/profile-protocol-options/"

type FirewallProfileProtocolOptions struct {
	Cifs                  *FirewallProfileProtocolOptionsCifs          `json:"cifs,omitempty"`
	Comment               *string                                      `json:"comment,omitempty"`
	Dns                   *FirewallProfileProtocolOptionsDns           `json:"dns,omitempty"`
	Ftp                   *FirewallProfileProtocolOptionsFtp           `json:"ftp,omitempty"`
	Http                  *FirewallProfileProtocolOptionsHttp          `json:"http,omitempty"`
	Imap                  *FirewallProfileProtocolOptionsImap          `json:"imap,omitempty"`
	MailSignature         *FirewallProfileProtocolOptionsMailSignature `json:"mail-signature,omitempty"`
	Mapi                  *FirewallProfileProtocolOptionsMapi          `json:"mapi,omitempty"`
	Name                  *string                                      `json:"name,omitempty"`
	Nntp                  *FirewallProfileProtocolOptionsNntp          `json:"nntp,omitempty"`
	OversizeLog           *string                                      `json:"oversize-log,omitempty"`
	Pop3                  *FirewallProfileProtocolOptionsPop3          `json:"pop3,omitempty"`
	ReplacemsgGroup       *string                                      `json:"replacemsg-group,omitempty"`
	RpcOverHttp           *string                                      `json:"rpc-over-http,omitempty"`
	Smtp                  *FirewallProfileProtocolOptionsSmtp          `json:"smtp,omitempty"`
	Ssh                   *FirewallProfileProtocolOptionsSsh           `json:"ssh,omitempty"`
	SwitchingProtocolsLog *string                                      `json:"switching-protocols-log,omitempty"`
}

const FirewallProfileProtocolOptionsCifsPath = "firewall/profile-protocol-options/cifs/"

type FirewallProfileProtocolOptionsCifs struct {
	DomainController          *string                                           `json:"domain-controller,omitempty"`
	Options                   *string                                           `json:"options,omitempty"`
	OversizeLimit             *int64                                            `json:"oversize-limit,omitempty"`
	Ports                     *int64                                            `json:"ports,omitempty"`
	ScanBzip2                 *string                                           `json:"scan-bzip2,omitempty"`
	ServerCredentialType      *string                                           `json:"server-credential-type,omitempty"`
	ServerKeytab              *[]FirewallProfileProtocolOptionsCifsServerKeytab `json:"server-keytab,omitempty"`
	Status                    *string                                           `json:"status,omitempty"`
	TcpWindowMaximum          *int64                                            `json:"tcp-window-maximum,omitempty"`
	TcpWindowMinimum          *int64                                            `json:"tcp-window-minimum,omitempty"`
	TcpWindowSize             *int64                                            `json:"tcp-window-size,omitempty"`
	TcpWindowType             *string                                           `json:"tcp-window-type,omitempty"`
	UncompressedNestLimit     *int64                                            `json:"uncompressed-nest-limit,omitempty"`
	UncompressedOversizeLimit *int64                                            `json:"uncompressed-oversize-limit,omitempty"`
}

const FirewallProfileProtocolOptionsCifsServerKeytabPath = "firewall/profile-protocol-options/cifs/server-keytab/"

type FirewallProfileProtocolOptionsCifsServerKeytab struct {
	Keytab    *string `json:"keytab,omitempty"`
	Principal *string `json:"principal,omitempty"`
}

const FirewallProfileProtocolOptionsDnsPath = "firewall/profile-protocol-options/dns/"

type FirewallProfileProtocolOptionsDns struct {
	Ports  *int64  `json:"ports,omitempty"`
	Status *string `json:"status,omitempty"`
}

const FirewallProfileProtocolOptionsFtpPath = "firewall/profile-protocol-options/ftp/"

type FirewallProfileProtocolOptionsFtp struct {
	ComfortAmount                *int64  `json:"comfort-amount,omitempty"`
	ComfortInterval              *int64  `json:"comfort-interval,omitempty"`
	InspectAll                   *string `json:"inspect-all,omitempty"`
	Options                      *string `json:"options,omitempty"`
	OversizeLimit                *int64  `json:"oversize-limit,omitempty"`
	Ports                        *int64  `json:"ports,omitempty"`
	ScanBzip2                    *string `json:"scan-bzip2,omitempty"`
	SslOffloaded                 *string `json:"ssl-offloaded,omitempty"`
	Status                       *string `json:"status,omitempty"`
	StreamBasedUncompressedLimit *int64  `json:"stream-based-uncompressed-limit,omitempty"`
	TcpWindowMaximum             *int64  `json:"tcp-window-maximum,omitempty"`
	TcpWindowMinimum             *int64  `json:"tcp-window-minimum,omitempty"`
	TcpWindowSize                *int64  `json:"tcp-window-size,omitempty"`
	TcpWindowType                *string `json:"tcp-window-type,omitempty"`
	UncompressedNestLimit        *int64  `json:"uncompressed-nest-limit,omitempty"`
	UncompressedOversizeLimit    *int64  `json:"uncompressed-oversize-limit,omitempty"`
}

const FirewallProfileProtocolOptionsHttpPath = "firewall/profile-protocol-options/http/"

type FirewallProfileProtocolOptionsHttp struct {
	BlockPageStatusCode          *int64  `json:"block-page-status-code,omitempty"`
	ComfortAmount                *int64  `json:"comfort-amount,omitempty"`
	ComfortInterval              *int64  `json:"comfort-interval,omitempty"`
	InspectAll                   *string `json:"inspect-all,omitempty"`
	Options                      *string `json:"options,omitempty"`
	OversizeLimit                *int64  `json:"oversize-limit,omitempty"`
	Ports                        *int64  `json:"ports,omitempty"`
	PostLang                     *string `json:"post-lang,omitempty"`
	ProxyAfterTcpHandshake       *string `json:"proxy-after-tcp-handshake,omitempty"`
	RangeBlock                   *string `json:"range-block,omitempty"`
	RetryCount                   *int64  `json:"retry-count,omitempty"`
	ScanBzip2                    *string `json:"scan-bzip2,omitempty"`
	SslOffloaded                 *string `json:"ssl-offloaded,omitempty"`
	Status                       *string `json:"status,omitempty"`
	StreamBasedUncompressedLimit *int64  `json:"stream-based-uncompressed-limit,omitempty"`
	StreamingContentBypass       *string `json:"streaming-content-bypass,omitempty"`
	StripXForwardedFor           *string `json:"strip-x-forwarded-for,omitempty"`
	SwitchingProtocols           *string `json:"switching-protocols,omitempty"`
	TcpWindowMaximum             *int64  `json:"tcp-window-maximum,omitempty"`
	TcpWindowMinimum             *int64  `json:"tcp-window-minimum,omitempty"`
	TcpWindowSize                *int64  `json:"tcp-window-size,omitempty"`
	TcpWindowType                *string `json:"tcp-window-type,omitempty"`
	TunnelNonHttp                *string `json:"tunnel-non-http,omitempty"`
	UncompressedNestLimit        *int64  `json:"uncompressed-nest-limit,omitempty"`
	UncompressedOversizeLimit    *int64  `json:"uncompressed-oversize-limit,omitempty"`
	UnknownHttpVersion           *string `json:"unknown-http-version,omitempty"`
}

const FirewallProfileProtocolOptionsImapPath = "firewall/profile-protocol-options/imap/"

type FirewallProfileProtocolOptionsImap struct {
	InspectAll                *string `json:"inspect-all,omitempty"`
	Options                   *string `json:"options,omitempty"`
	OversizeLimit             *int64  `json:"oversize-limit,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake    *string `json:"proxy-after-tcp-handshake,omitempty"`
	ScanBzip2                 *string `json:"scan-bzip2,omitempty"`
	SslOffloaded              *string `json:"ssl-offloaded,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UncompressedNestLimit     *int64  `json:"uncompressed-nest-limit,omitempty"`
	UncompressedOversizeLimit *int64  `json:"uncompressed-oversize-limit,omitempty"`
}

const FirewallProfileProtocolOptionsMailSignaturePath = "firewall/profile-protocol-options/mail-signature/"

type FirewallProfileProtocolOptionsMailSignature struct {
	Signature *string `json:"signature,omitempty"`
	Status    *string `json:"status,omitempty"`
}

const FirewallProfileProtocolOptionsMapiPath = "firewall/profile-protocol-options/mapi/"

type FirewallProfileProtocolOptionsMapi struct {
	Options                   *string `json:"options,omitempty"`
	OversizeLimit             *int64  `json:"oversize-limit,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ScanBzip2                 *string `json:"scan-bzip2,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UncompressedNestLimit     *int64  `json:"uncompressed-nest-limit,omitempty"`
	UncompressedOversizeLimit *int64  `json:"uncompressed-oversize-limit,omitempty"`
}

const FirewallProfileProtocolOptionsNntpPath = "firewall/profile-protocol-options/nntp/"

type FirewallProfileProtocolOptionsNntp struct {
	InspectAll                *string `json:"inspect-all,omitempty"`
	Options                   *string `json:"options,omitempty"`
	OversizeLimit             *int64  `json:"oversize-limit,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake    *string `json:"proxy-after-tcp-handshake,omitempty"`
	ScanBzip2                 *string `json:"scan-bzip2,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UncompressedNestLimit     *int64  `json:"uncompressed-nest-limit,omitempty"`
	UncompressedOversizeLimit *int64  `json:"uncompressed-oversize-limit,omitempty"`
}

const FirewallProfileProtocolOptionsPop3Path = "firewall/profile-protocol-options/pop3/"

type FirewallProfileProtocolOptionsPop3 struct {
	InspectAll                *string `json:"inspect-all,omitempty"`
	Options                   *string `json:"options,omitempty"`
	OversizeLimit             *int64  `json:"oversize-limit,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake    *string `json:"proxy-after-tcp-handshake,omitempty"`
	ScanBzip2                 *string `json:"scan-bzip2,omitempty"`
	SslOffloaded              *string `json:"ssl-offloaded,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UncompressedNestLimit     *int64  `json:"uncompressed-nest-limit,omitempty"`
	UncompressedOversizeLimit *int64  `json:"uncompressed-oversize-limit,omitempty"`
}

const FirewallProfileProtocolOptionsSmtpPath = "firewall/profile-protocol-options/smtp/"

type FirewallProfileProtocolOptionsSmtp struct {
	InspectAll                *string `json:"inspect-all,omitempty"`
	Options                   *string `json:"options,omitempty"`
	OversizeLimit             *int64  `json:"oversize-limit,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake    *string `json:"proxy-after-tcp-handshake,omitempty"`
	ScanBzip2                 *string `json:"scan-bzip2,omitempty"`
	ServerBusy                *string `json:"server-busy,omitempty"`
	SslOffloaded              *string `json:"ssl-offloaded,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UncompressedNestLimit     *int64  `json:"uncompressed-nest-limit,omitempty"`
	UncompressedOversizeLimit *int64  `json:"uncompressed-oversize-limit,omitempty"`
}

const FirewallProfileProtocolOptionsSshPath = "firewall/profile-protocol-options/ssh/"

type FirewallProfileProtocolOptionsSsh struct {
	ComfortAmount                *int64  `json:"comfort-amount,omitempty"`
	ComfortInterval              *int64  `json:"comfort-interval,omitempty"`
	Options                      *string `json:"options,omitempty"`
	OversizeLimit                *int64  `json:"oversize-limit,omitempty"`
	ScanBzip2                    *string `json:"scan-bzip2,omitempty"`
	SslOffloaded                 *string `json:"ssl-offloaded,omitempty"`
	StreamBasedUncompressedLimit *int64  `json:"stream-based-uncompressed-limit,omitempty"`
	TcpWindowMaximum             *int64  `json:"tcp-window-maximum,omitempty"`
	TcpWindowMinimum             *int64  `json:"tcp-window-minimum,omitempty"`
	TcpWindowSize                *int64  `json:"tcp-window-size,omitempty"`
	TcpWindowType                *string `json:"tcp-window-type,omitempty"`
	UncompressedNestLimit        *int64  `json:"uncompressed-nest-limit,omitempty"`
	UncompressedOversizeLimit    *int64  `json:"uncompressed-oversize-limit,omitempty"`
}
