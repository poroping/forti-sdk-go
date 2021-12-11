package models

const EmailfilterProfilePath = "emailfilter/profile/"

type EmailfilterProfile struct {
	Comment                   *string                           `json:"comment,omitempty"`
	External                  *string                           `json:"external,omitempty"`
	FeatureSet                *string                           `json:"feature-set,omitempty"`
	Gmail                     []EmailfilterProfileGmail         `json:"gmail,omitempty"`
	Imap                      []EmailfilterProfileImap          `json:"imap,omitempty"`
	Mapi                      []EmailfilterProfileMapi          `json:"mapi,omitempty"`
	MsnHotmail                []EmailfilterProfileMsnHotmail    `json:"msn-hotmail,omitempty"`
	Name                      *string                           `json:"name,omitempty"`
	Options                   *string                           `json:"options,omitempty"`
	OtherWebmails             []EmailfilterProfileOtherWebmails `json:"other-webmails,omitempty"`
	Pop3                      []EmailfilterProfilePop3          `json:"pop3,omitempty"`
	ReplacemsgGroup           *string                           `json:"replacemsg-group,omitempty"`
	Smtp                      []EmailfilterProfileSmtp          `json:"smtp,omitempty"`
	SpamBalTable              *int64                            `json:"spam-bal-table,omitempty"`
	SpamBwordTable            *int64                            `json:"spam-bword-table,omitempty"`
	SpamBwordThreshold        *int64                            `json:"spam-bword-threshold,omitempty"`
	SpamFiltering             *string                           `json:"spam-filtering,omitempty"`
	SpamIptrustTable          *int64                            `json:"spam-iptrust-table,omitempty"`
	SpamLog                   *string                           `json:"spam-log,omitempty"`
	SpamLogFortiguardResponse *string                           `json:"spam-log-fortiguard-response,omitempty"`
	SpamMheaderTable          *int64                            `json:"spam-mheader-table,omitempty"`
	SpamRblTable              *int64                            `json:"spam-rbl-table,omitempty"`
	YahooMail                 []EmailfilterProfileYahooMail     `json:"yahoo-mail,omitempty"`
}

type EmailfilterProfileGmail struct {
	LogAll *string `json:"log-all,omitempty"`
}

type EmailfilterProfileImap struct {
	Action  *string `json:"action,omitempty"`
	LogAll  *string `json:"log-all,omitempty"`
	TagMsg  *string `json:"tag-msg,omitempty"`
	TagType *string `json:"tag-type,omitempty"`
}

type EmailfilterProfileMapi struct {
	Action *string `json:"action,omitempty"`
	LogAll *string `json:"log-all,omitempty"`
}

type EmailfilterProfileMsnHotmail struct {
	LogAll *string `json:"log-all,omitempty"`
}

type EmailfilterProfileOtherWebmails struct {
	LogAll *string `json:"log-all,omitempty"`
}

type EmailfilterProfilePop3 struct {
	Action  *string `json:"action,omitempty"`
	LogAll  *string `json:"log-all,omitempty"`
	TagMsg  *string `json:"tag-msg,omitempty"`
	TagType *string `json:"tag-type,omitempty"`
}

type EmailfilterProfileSmtp struct {
	Action        *string `json:"action,omitempty"`
	Hdrip         *string `json:"hdrip,omitempty"`
	LocalOverride *string `json:"local-override,omitempty"`
	LogAll        *string `json:"log-all,omitempty"`
	TagMsg        *string `json:"tag-msg,omitempty"`
	TagType       *string `json:"tag-type,omitempty"`
}

type EmailfilterProfileYahooMail struct {
	LogAll *string `json:"log-all,omitempty"`
}
