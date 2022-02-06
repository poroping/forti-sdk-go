package models

const WafProfilePath = "waf/profile/"

type WafProfile struct {
	AddressList *WafProfileAddressList `json:"address-list,omitempty"`
	Comment     *string                `json:"comment,omitempty"`
	Constraint  *WafProfileConstraint  `json:"constraint,omitempty"`
	ExtendedLog *string                `json:"extended-log,omitempty"`
	External    *string                `json:"external,omitempty"`
	Method      *WafProfileMethod      `json:"method,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Signature   *WafProfileSignature   `json:"signature,omitempty"`
	UrlAccess   *[]WafProfileUrlAccess `json:"url-access,omitempty"`
}

type WafProfileAddressList struct {
	BlockedAddress *[]WafProfileAddressListBlockedAddress `json:"blocked-address,omitempty"`
	BlockedLog     *string                                `json:"blocked-log,omitempty"`
	Severity       *string                                `json:"severity,omitempty"`
	Status         *string                                `json:"status,omitempty"`
	TrustedAddress *[]WafProfileAddressListTrustedAddress `json:"trusted-address,omitempty"`
}

type WafProfileAddressListBlockedAddress struct {
	Name *string `json:"name,omitempty"`
}

type WafProfileAddressListTrustedAddress struct {
	Name *string `json:"name,omitempty"`
}

type WafProfileConstraint struct {
	ContentLength   *WafProfileConstraintContentLength   `json:"content-length,omitempty"`
	Exception       *[]WafProfileConstraintException     `json:"exception,omitempty"`
	HeaderLength    *WafProfileConstraintHeaderLength    `json:"header-length,omitempty"`
	Hostname        *WafProfileConstraintHostname        `json:"hostname,omitempty"`
	LineLength      *WafProfileConstraintLineLength      `json:"line-length,omitempty"`
	Malformed       *WafProfileConstraintMalformed       `json:"malformed,omitempty"`
	MaxCookie       *WafProfileConstraintMaxCookie       `json:"max-cookie,omitempty"`
	MaxHeaderLine   *WafProfileConstraintMaxHeaderLine   `json:"max-header-line,omitempty"`
	MaxRangeSegment *WafProfileConstraintMaxRangeSegment `json:"max-range-segment,omitempty"`
	MaxUrlParam     *WafProfileConstraintMaxUrlParam     `json:"max-url-param,omitempty"`
	Method          *WafProfileConstraintMethod          `json:"method,omitempty"`
	ParamLength     *WafProfileConstraintParamLength     `json:"param-length,omitempty"`
	UrlParamLength  *WafProfileConstraintUrlParamLength  `json:"url-param-length,omitempty"`
	Version         *WafProfileConstraintVersion         `json:"version,omitempty"`
}

type WafProfileConstraintContentLength struct {
	Action   *string `json:"action,omitempty"`
	Length   *int64  `json:"length,omitempty"`
	Log      *string `json:"log,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status   *string `json:"status,omitempty"`
}

type WafProfileConstraintException struct {
	Address         *string `json:"address,omitempty"`
	ContentLength   *string `json:"content-length,omitempty"`
	HeaderLength    *string `json:"header-length,omitempty"`
	Hostname        *string `json:"hostname,omitempty"`
	Id              *int64  `json:"id,omitempty"`
	LineLength      *string `json:"line-length,omitempty"`
	Malformed       *string `json:"malformed,omitempty"`
	MaxCookie       *string `json:"max-cookie,omitempty"`
	MaxHeaderLine   *string `json:"max-header-line,omitempty"`
	MaxRangeSegment *string `json:"max-range-segment,omitempty"`
	MaxUrlParam     *string `json:"max-url-param,omitempty"`
	Method          *string `json:"method,omitempty"`
	ParamLength     *string `json:"param-length,omitempty"`
	Pattern         *string `json:"pattern,omitempty"`
	Regex           *string `json:"regex,omitempty"`
	UrlParamLength  *string `json:"url-param-length,omitempty"`
	Version         *string `json:"version,omitempty"`
}

type WafProfileConstraintHeaderLength struct {
	Action   *string `json:"action,omitempty"`
	Length   *int64  `json:"length,omitempty"`
	Log      *string `json:"log,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status   *string `json:"status,omitempty"`
}

type WafProfileConstraintHostname struct {
	Action   *string `json:"action,omitempty"`
	Log      *string `json:"log,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status   *string `json:"status,omitempty"`
}

type WafProfileConstraintLineLength struct {
	Action   *string `json:"action,omitempty"`
	Length   *int64  `json:"length,omitempty"`
	Log      *string `json:"log,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status   *string `json:"status,omitempty"`
}

type WafProfileConstraintMalformed struct {
	Action   *string `json:"action,omitempty"`
	Log      *string `json:"log,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status   *string `json:"status,omitempty"`
}

type WafProfileConstraintMaxCookie struct {
	Action    *string `json:"action,omitempty"`
	Log       *string `json:"log,omitempty"`
	MaxCookie *int64  `json:"max-cookie,omitempty"`
	Severity  *string `json:"severity,omitempty"`
	Status    *string `json:"status,omitempty"`
}

type WafProfileConstraintMaxHeaderLine struct {
	Action        *string `json:"action,omitempty"`
	Log           *string `json:"log,omitempty"`
	MaxHeaderLine *int64  `json:"max-header-line,omitempty"`
	Severity      *string `json:"severity,omitempty"`
	Status        *string `json:"status,omitempty"`
}

type WafProfileConstraintMaxRangeSegment struct {
	Action          *string `json:"action,omitempty"`
	Log             *string `json:"log,omitempty"`
	MaxRangeSegment *int64  `json:"max-range-segment,omitempty"`
	Severity        *string `json:"severity,omitempty"`
	Status          *string `json:"status,omitempty"`
}

type WafProfileConstraintMaxUrlParam struct {
	Action      *string `json:"action,omitempty"`
	Log         *string `json:"log,omitempty"`
	MaxUrlParam *int64  `json:"max-url-param,omitempty"`
	Severity    *string `json:"severity,omitempty"`
	Status      *string `json:"status,omitempty"`
}

type WafProfileConstraintMethod struct {
	Action   *string `json:"action,omitempty"`
	Log      *string `json:"log,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status   *string `json:"status,omitempty"`
}

type WafProfileConstraintParamLength struct {
	Action   *string `json:"action,omitempty"`
	Length   *int64  `json:"length,omitempty"`
	Log      *string `json:"log,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status   *string `json:"status,omitempty"`
}

type WafProfileConstraintUrlParamLength struct {
	Action   *string `json:"action,omitempty"`
	Length   *int64  `json:"length,omitempty"`
	Log      *string `json:"log,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status   *string `json:"status,omitempty"`
}

type WafProfileConstraintVersion struct {
	Action   *string `json:"action,omitempty"`
	Log      *string `json:"log,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status   *string `json:"status,omitempty"`
}

type WafProfileMethod struct {
	DefaultAllowedMethods *string                         `json:"default-allowed-methods,omitempty"`
	Log                   *string                         `json:"log,omitempty"`
	MethodPolicy          *[]WafProfileMethodMethodPolicy `json:"method-policy,omitempty"`
	Severity              *string                         `json:"severity,omitempty"`
	Status                *string                         `json:"status,omitempty"`
}

type WafProfileMethodMethodPolicy struct {
	Address        *string `json:"address,omitempty"`
	AllowedMethods *string `json:"allowed-methods,omitempty"`
	Id             *int64  `json:"id,omitempty"`
	Pattern        *string `json:"pattern,omitempty"`
	Regex          *string `json:"regex,omitempty"`
}

type WafProfileSignature struct {
	CreditCardDetectionThreshold *int64                                  `json:"credit-card-detection-threshold,omitempty"`
	CustomSignature              *[]WafProfileSignatureCustomSignature   `json:"custom-signature,omitempty"`
	DisabledSignature            *[]WafProfileSignatureDisabledSignature `json:"disabled-signature,omitempty"`
	DisabledSubClass             *[]WafProfileSignatureDisabledSubClass  `json:"disabled-sub-class,omitempty"`
	MainClass                    *[]WafProfileSignatureMainClass         `json:"main-class,omitempty"`
}

type WafProfileSignatureCustomSignature struct {
	Action          *string `json:"action,omitempty"`
	CaseSensitivity *string `json:"case-sensitivity,omitempty"`
	Direction       *string `json:"direction,omitempty"`
	Log             *string `json:"log,omitempty"`
	Name            *string `json:"name,omitempty"`
	Pattern         *string `json:"pattern,omitempty"`
	Severity        *string `json:"severity,omitempty"`
	Status          *string `json:"status,omitempty"`
	Target          *string `json:"target,omitempty"`
}

type WafProfileSignatureDisabledSignature struct {
	Id *int64 `json:"id,omitempty"`
}

type WafProfileSignatureDisabledSubClass struct {
	Id *int64 `json:"id,omitempty"`
}

type WafProfileSignatureMainClass struct {
	Action   *string `json:"action,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Log      *string `json:"log,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status   *string `json:"status,omitempty"`
}

type WafProfileUrlAccess struct {
	AccessPattern *[]WafProfileUrlAccessAccessPattern `json:"access-pattern,omitempty"`
	Action        *string                             `json:"action,omitempty"`
	Address       *string                             `json:"address,omitempty"`
	Id            *int64                              `json:"id,omitempty"`
	Log           *string                             `json:"log,omitempty"`
	Severity      *string                             `json:"severity,omitempty"`
}

type WafProfileUrlAccessAccessPattern struct {
	Id      *int64  `json:"id,omitempty"`
	Negate  *string `json:"negate,omitempty"`
	Pattern *string `json:"pattern,omitempty"`
	Regex   *string `json:"regex,omitempty"`
	Srcaddr *string `json:"srcaddr,omitempty"`
}
