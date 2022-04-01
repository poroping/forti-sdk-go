package models

const IcapProfilePath = "icap/profile/"

type IcapProfile struct {
	The204Response         *string                           `json:"204-response,omitempty"`
	The204SizeLimit        *int64                            `json:"204-size-limit,omitempty"`
	ChunkEncap             *string                           `json:"chunk-encap,omitempty"`
	ExtensionFeature       *string                           `json:"extension-feature,omitempty"`
	FileTransfer           *string                           `json:"file-transfer,omitempty"`
	FileTransferFailure    *string                           `json:"file-transfer-failure,omitempty"`
	FileTransferPath       *string                           `json:"file-transfer-path,omitempty"`
	FileTransferServer     *string                           `json:"file-transfer-server,omitempty"`
	IcapBlockLog           *string                           `json:"icap-block-log,omitempty"`
	IcapHeaders            *[]IcapProfileIcapHeaders         `json:"icap-headers,omitempty"`
	Methods                *string                           `json:"methods,omitempty"`
	Name                   *string                           `json:"name,omitempty"`
	Preview                *string                           `json:"preview,omitempty"`
	PreviewDataLength      *int64                            `json:"preview-data-length,omitempty"`
	ReplacemsgGroup        *string                           `json:"replacemsg-group,omitempty"`
	Request                *string                           `json:"request,omitempty"`
	RequestFailure         *string                           `json:"request-failure,omitempty"`
	RequestPath            *string                           `json:"request-path,omitempty"`
	RequestServer          *string                           `json:"request-server,omitempty"`
	RespmodDefaultAction   *string                           `json:"respmod-default-action,omitempty"`
	RespmodForwardRules    *[]IcapProfileRespmodForwardRules `json:"respmod-forward-rules,omitempty"`
	Response               *string                           `json:"response,omitempty"`
	ResponseFailure        *string                           `json:"response-failure,omitempty"`
	ResponsePath           *string                           `json:"response-path,omitempty"`
	ResponseReqHdr         *string                           `json:"response-req-hdr,omitempty"`
	ResponseServer         *string                           `json:"response-server,omitempty"`
	ScanProgressInterval   *int64                            `json:"scan-progress-interval,omitempty"`
	StreamingContentBypass *string                           `json:"streaming-content-bypass,omitempty"`
	Timeout                *int64                            `json:"timeout,omitempty"`
}

type IcapProfileIcapHeaders struct {
	Base64Encoding *string `json:"base64-encoding,omitempty"`
	Content        *string `json:"content,omitempty"`
	Id             *int64  `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
}

type IcapProfileRespmodForwardRules struct {
	Action             *string                                             `json:"action,omitempty"`
	HeaderGroup        *[]IcapProfileRespmodForwardRulesHeaderGroup        `json:"header-group,omitempty"`
	Host               *string                                             `json:"host,omitempty"`
	HttpRespStatusCode *[]IcapProfileRespmodForwardRulesHttpRespStatusCode `json:"http-resp-status-code,omitempty"`
	Name               *string                                             `json:"name,omitempty"`
}

type IcapProfileRespmodForwardRulesHeaderGroup struct {
	CaseSensitivity *string `json:"case-sensitivity,omitempty"`
	Header          *string `json:"header,omitempty"`
	HeaderName      *string `json:"header-name,omitempty"`
	Id              *int64  `json:"id,omitempty"`
}

type IcapProfileRespmodForwardRulesHttpRespStatusCode struct {
	Code *int64 `json:"code,omitempty"`
}
