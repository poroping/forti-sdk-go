package models

const VoipProfilePath = "voip/profile/"

type VoipProfile struct {
	Comment    *string          `json:"comment,omitempty"`
	FeatureSet *string          `json:"feature-set,omitempty"`
	Msrp       *VoipProfileMsrp `json:"msrp,omitempty"`
	Name       *string          `json:"name,omitempty"`
	Sccp       *VoipProfileSccp `json:"sccp,omitempty"`
	Sip        *VoipProfileSip  `json:"sip,omitempty"`
}

const VoipProfileMsrpPath = "voip/profile/msrp/"

type VoipProfileMsrp struct {
	LogViolations    *string `json:"log-violations,omitempty"`
	MaxMsgSize       *int64  `json:"max-msg-size,omitempty"`
	MaxMsgSizeAction *string `json:"max-msg-size-action,omitempty"`
	Status           *string `json:"status,omitempty"`
}

const VoipProfileSccpPath = "voip/profile/sccp/"

type VoipProfileSccp struct {
	BlockMcast     *string `json:"block-mcast,omitempty"`
	LogCallSummary *string `json:"log-call-summary,omitempty"`
	LogViolations  *string `json:"log-violations,omitempty"`
	MaxCalls       *int64  `json:"max-calls,omitempty"`
	Status         *string `json:"status,omitempty"`
	VerifyHeader   *string `json:"verify-header,omitempty"`
}

const VoipProfileSipPath = "voip/profile/sip/"

type VoipProfileSip struct {
	AckRate                          *int64  `json:"ack-rate,omitempty"`
	AckRateTrack                     *string `json:"ack-rate-track,omitempty"`
	BlockAck                         *string `json:"block-ack,omitempty"`
	BlockBye                         *string `json:"block-bye,omitempty"`
	BlockCancel                      *string `json:"block-cancel,omitempty"`
	BlockGeoRedOptions               *string `json:"block-geo-red-options,omitempty"`
	BlockInfo                        *string `json:"block-info,omitempty"`
	BlockInvite                      *string `json:"block-invite,omitempty"`
	BlockLongLines                   *string `json:"block-long-lines,omitempty"`
	BlockMessage                     *string `json:"block-message,omitempty"`
	BlockNotify                      *string `json:"block-notify,omitempty"`
	BlockOptions                     *string `json:"block-options,omitempty"`
	BlockPrack                       *string `json:"block-prack,omitempty"`
	BlockPublish                     *string `json:"block-publish,omitempty"`
	BlockRefer                       *string `json:"block-refer,omitempty"`
	BlockRegister                    *string `json:"block-register,omitempty"`
	BlockSubscribe                   *string `json:"block-subscribe,omitempty"`
	BlockUnknown                     *string `json:"block-unknown,omitempty"`
	BlockUpdate                      *string `json:"block-update,omitempty"`
	ByeRate                          *int64  `json:"bye-rate,omitempty"`
	ByeRateTrack                     *string `json:"bye-rate-track,omitempty"`
	CallKeepalive                    *int64  `json:"call-keepalive,omitempty"`
	CancelRate                       *int64  `json:"cancel-rate,omitempty"`
	CancelRateTrack                  *string `json:"cancel-rate-track,omitempty"`
	ContactFixup                     *string `json:"contact-fixup,omitempty"`
	HntRestrictSourceIp              *string `json:"hnt-restrict-source-ip,omitempty"`
	HostedNatTraversal               *string `json:"hosted-nat-traversal,omitempty"`
	InfoRate                         *int64  `json:"info-rate,omitempty"`
	InfoRateTrack                    *string `json:"info-rate-track,omitempty"`
	InviteRate                       *int64  `json:"invite-rate,omitempty"`
	InviteRateTrack                  *string `json:"invite-rate-track,omitempty"`
	IpsRtp                           *string `json:"ips-rtp,omitempty"`
	LogCallSummary                   *string `json:"log-call-summary,omitempty"`
	LogViolations                    *string `json:"log-violations,omitempty"`
	MalformedHeaderAllow             *string `json:"malformed-header-allow,omitempty"`
	MalformedHeaderCallId            *string `json:"malformed-header-call-id,omitempty"`
	MalformedHeaderContact           *string `json:"malformed-header-contact,omitempty"`
	MalformedHeaderContentLength     *string `json:"malformed-header-content-length,omitempty"`
	MalformedHeaderContentType       *string `json:"malformed-header-content-type,omitempty"`
	MalformedHeaderCseq              *string `json:"malformed-header-cseq,omitempty"`
	MalformedHeaderExpires           *string `json:"malformed-header-expires,omitempty"`
	MalformedHeaderFrom              *string `json:"malformed-header-from,omitempty"`
	MalformedHeaderMaxForwards       *string `json:"malformed-header-max-forwards,omitempty"`
	MalformedHeaderNoProxyRequire    *string `json:"malformed-header-no-proxy-require,omitempty"`
	MalformedHeaderNoRequire         *string `json:"malformed-header-no-require,omitempty"`
	MalformedHeaderPAssertedIdentity *string `json:"malformed-header-p-asserted-identity,omitempty"`
	MalformedHeaderRack              *string `json:"malformed-header-rack,omitempty"`
	MalformedHeaderRecordRoute       *string `json:"malformed-header-record-route,omitempty"`
	MalformedHeaderRoute             *string `json:"malformed-header-route,omitempty"`
	MalformedHeaderRseq              *string `json:"malformed-header-rseq,omitempty"`
	MalformedHeaderSdpA              *string `json:"malformed-header-sdp-a,omitempty"`
	MalformedHeaderSdpB              *string `json:"malformed-header-sdp-b,omitempty"`
	MalformedHeaderSdpC              *string `json:"malformed-header-sdp-c,omitempty"`
	MalformedHeaderSdpI              *string `json:"malformed-header-sdp-i,omitempty"`
	MalformedHeaderSdpK              *string `json:"malformed-header-sdp-k,omitempty"`
	MalformedHeaderSdpM              *string `json:"malformed-header-sdp-m,omitempty"`
	MalformedHeaderSdpO              *string `json:"malformed-header-sdp-o,omitempty"`
	MalformedHeaderSdpR              *string `json:"malformed-header-sdp-r,omitempty"`
	MalformedHeaderSdpS              *string `json:"malformed-header-sdp-s,omitempty"`
	MalformedHeaderSdpT              *string `json:"malformed-header-sdp-t,omitempty"`
	MalformedHeaderSdpV              *string `json:"malformed-header-sdp-v,omitempty"`
	MalformedHeaderSdpZ              *string `json:"malformed-header-sdp-z,omitempty"`
	MalformedHeaderTo                *string `json:"malformed-header-to,omitempty"`
	MalformedHeaderVia               *string `json:"malformed-header-via,omitempty"`
	MalformedRequestLine             *string `json:"malformed-request-line,omitempty"`
	MaxBodyLength                    *int64  `json:"max-body-length,omitempty"`
	MaxDialogs                       *int64  `json:"max-dialogs,omitempty"`
	MaxIdleDialogs                   *int64  `json:"max-idle-dialogs,omitempty"`
	MaxLineLength                    *int64  `json:"max-line-length,omitempty"`
	MessageRate                      *int64  `json:"message-rate,omitempty"`
	MessageRateTrack                 *string `json:"message-rate-track,omitempty"`
	NatPortRange                     *string `json:"nat-port-range,omitempty"`
	NatTrace                         *string `json:"nat-trace,omitempty"`
	NoSdpFixup                       *string `json:"no-sdp-fixup,omitempty"`
	NotifyRate                       *int64  `json:"notify-rate,omitempty"`
	NotifyRateTrack                  *string `json:"notify-rate-track,omitempty"`
	OpenContactPinhole               *string `json:"open-contact-pinhole,omitempty"`
	OpenRecordRoutePinhole           *string `json:"open-record-route-pinhole,omitempty"`
	OpenRegisterPinhole              *string `json:"open-register-pinhole,omitempty"`
	OpenViaPinhole                   *string `json:"open-via-pinhole,omitempty"`
	OptionsRate                      *int64  `json:"options-rate,omitempty"`
	OptionsRateTrack                 *string `json:"options-rate-track,omitempty"`
	PrackRate                        *int64  `json:"prack-rate,omitempty"`
	PrackRateTrack                   *string `json:"prack-rate-track,omitempty"`
	PreserveOverride                 *string `json:"preserve-override,omitempty"`
	ProvisionalInviteExpiryTime      *int64  `json:"provisional-invite-expiry-time,omitempty"`
	PublishRate                      *int64  `json:"publish-rate,omitempty"`
	PublishRateTrack                 *string `json:"publish-rate-track,omitempty"`
	ReferRate                        *int64  `json:"refer-rate,omitempty"`
	ReferRateTrack                   *string `json:"refer-rate-track,omitempty"`
	RegisterContactTrace             *string `json:"register-contact-trace,omitempty"`
	RegisterRate                     *int64  `json:"register-rate,omitempty"`
	RegisterRateTrack                *string `json:"register-rate-track,omitempty"`
	Rfc2543Branch                    *string `json:"rfc2543-branch,omitempty"`
	Rtp                              *string `json:"rtp,omitempty"`
	SslAlgorithm                     *string `json:"ssl-algorithm,omitempty"`
	SslAuthClient                    *string `json:"ssl-auth-client,omitempty"`
	SslAuthServer                    *string `json:"ssl-auth-server,omitempty"`
	SslClientCertificate             *string `json:"ssl-client-certificate,omitempty"`
	SslClientRenegotiation           *string `json:"ssl-client-renegotiation,omitempty"`
	SslMaxVersion                    *string `json:"ssl-max-version,omitempty"`
	SslMinVersion                    *string `json:"ssl-min-version,omitempty"`
	SslMode                          *string `json:"ssl-mode,omitempty"`
	SslPfs                           *string `json:"ssl-pfs,omitempty"`
	SslSendEmptyFrags                *string `json:"ssl-send-empty-frags,omitempty"`
	SslServerCertificate             *string `json:"ssl-server-certificate,omitempty"`
	Status                           *string `json:"status,omitempty"`
	StrictRegister                   *string `json:"strict-register,omitempty"`
	SubscribeRate                    *int64  `json:"subscribe-rate,omitempty"`
	SubscribeRateTrack               *string `json:"subscribe-rate-track,omitempty"`
	UnknownHeader                    *string `json:"unknown-header,omitempty"`
	UpdateRate                       *int64  `json:"update-rate,omitempty"`
	UpdateRateTrack                  *string `json:"update-rate-track,omitempty"`
}
