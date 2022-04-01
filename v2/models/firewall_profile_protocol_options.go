package models

const FirewallProfileProtocolOptionsPath = "firewall/profile-protocol-options/"

type FirewallProfileProtocolOptions struct {
	Cifs                  *FirewallProfileProtocolOptionsCifs          `json:"cifs,omitempty"`
	Comment               *string                                      `json:"comment,omitempty"`
	Dns                   *FirewallProfileProtocolOptionsDns           `json:"dns,omitempty"`
	FeatureSet            *string                                      `json:"feature-set,omitempty"`
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

type FirewallProfileProtocolOptionsCifsServerKeytab struct {
	Keytab    *string `json:"keytab,omitempty"`
	Principal *string `json:"principal,omitempty"`
}

type FirewallProfileProtocolOptionsDns struct {
	Ports  *int64  `json:"ports,omitempty"`
	Status *string `json:"status,omitempty"`
}

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

type FirewallProfileProtocolOptionsHttp struct {
	AddressIpRating              *string `json:"address-ip-rating,omitempty"`
	BlockPageStatusCode          *int64  `json:"block-page-status-code,omitempty"`
	ComfortAmount                *int64  `json:"comfort-amount,omitempty"`
	ComfortInterval              *int64  `json:"comfort-interval,omitempty"`
	FortinetBar                  *string `json:"fortinet-bar,omitempty"`
	FortinetBarPort              *int64  `json:"fortinet-bar-port,omitempty"`
	H2c                          *string `json:"h2c,omitempty"`
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

type FirewallProfileProtocolOptionsMailSignature struct {
	Signature *string `json:"signature,omitempty"`
	Status    *string `json:"status,omitempty"`
}

type FirewallProfileProtocolOptionsMapi struct {
	Options                   *string `json:"options,omitempty"`
	OversizeLimit             *int64  `json:"oversize-limit,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ScanBzip2                 *string `json:"scan-bzip2,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UncompressedNestLimit     *int64  `json:"uncompressed-nest-limit,omitempty"`
	UncompressedOversizeLimit *int64  `json:"uncompressed-oversize-limit,omitempty"`
}

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
