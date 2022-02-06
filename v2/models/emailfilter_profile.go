package models

const EmailfilterProfilePath = "emailfilter/profile/"

type EmailfilterProfile struct {
	Comment                   *string                          `json:"comment,omitempty"`
	External                  *string                          `json:"external,omitempty"`
	FeatureSet                *string                          `json:"feature-set,omitempty"`
	FileFilter                *EmailfilterProfileFileFilter    `json:"file-filter,omitempty"`
	Gmail                     *EmailfilterProfileGmail         `json:"gmail,omitempty"`
	Imap                      *EmailfilterProfileImap          `json:"imap,omitempty"`
	Mapi                      *EmailfilterProfileMapi          `json:"mapi,omitempty"`
	MsnHotmail                *EmailfilterProfileMsnHotmail    `json:"msn-hotmail,omitempty"`
	Name                      *string                          `json:"name,omitempty"`
	Options                   *string                          `json:"options,omitempty"`
	OtherWebmails             *EmailfilterProfileOtherWebmails `json:"other-webmails,omitempty"`
	Pop3                      *EmailfilterProfilePop3          `json:"pop3,omitempty"`
	ReplacemsgGroup           *string                          `json:"replacemsg-group,omitempty"`
	Smtp                      *EmailfilterProfileSmtp          `json:"smtp,omitempty"`
	SpamBalTable              *int64                           `json:"spam-bal-table,omitempty"`
	SpamBwlTable              *int64                           `json:"spam-bwl-table,omitempty"`
	SpamBwordTable            *int64                           `json:"spam-bword-table,omitempty"`
	SpamBwordThreshold        *int64                           `json:"spam-bword-threshold,omitempty"`
	SpamFiltering             *string                          `json:"spam-filtering,omitempty"`
	SpamIptrustTable          *int64                           `json:"spam-iptrust-table,omitempty"`
	SpamLog                   *string                          `json:"spam-log,omitempty"`
	SpamLogFortiguardResponse *string                          `json:"spam-log-fortiguard-response,omitempty"`
	SpamMheaderTable          *int64                           `json:"spam-mheader-table,omitempty"`
	SpamRblTable              *int64                           `json:"spam-rbl-table,omitempty"`
	YahooMail                 *EmailfilterProfileYahooMail     `json:"yahoo-mail,omitempty"`
}

type EmailfilterProfileFileFilter struct {
	Entries             *[]EmailfilterProfileFileFilterEntries `json:"entries,omitempty"`
	Log                 *string                                `json:"log,omitempty"`
	ScanArchiveContents *string                                `json:"scan-archive-contents,omitempty"`
	Status              *string                                `json:"status,omitempty"`
}

type EmailfilterProfileFileFilterEntries struct {
	Action            *string                                        `json:"action,omitempty"`
	Comment           *string                                        `json:"comment,omitempty"`
	FileType          *[]EmailfilterProfileFileFilterEntriesFileType `json:"file-type,omitempty"`
	Filter            *string                                        `json:"filter,omitempty"`
	PasswordProtected *string                                        `json:"password-protected,omitempty"`
	Protocol          *string                                        `json:"protocol,omitempty"`
}

type EmailfilterProfileFileFilterEntriesFileType struct {
	Name *string `json:"name,omitempty"`
}

type EmailfilterProfileGmail struct {
	Log    *string `json:"log,omitempty"`
	LogAll *string `json:"log-all,omitempty"`
}

type EmailfilterProfileImap struct {
	Action  *string `json:"action,omitempty"`
	Log     *string `json:"log,omitempty"`
	LogAll  *string `json:"log-all,omitempty"`
	TagMsg  *string `json:"tag-msg,omitempty"`
	TagType *string `json:"tag-type,omitempty"`
}

type EmailfilterProfileMapi struct {
	Action *string `json:"action,omitempty"`
	Log    *string `json:"log,omitempty"`
	LogAll *string `json:"log-all,omitempty"`
}

type EmailfilterProfileMsnHotmail struct {
	Log    *string `json:"log,omitempty"`
	LogAll *string `json:"log-all,omitempty"`
}

type EmailfilterProfileOtherWebmails struct {
	LogAll *string `json:"log-all,omitempty"`
}

type EmailfilterProfilePop3 struct {
	Action  *string `json:"action,omitempty"`
	Log     *string `json:"log,omitempty"`
	LogAll  *string `json:"log-all,omitempty"`
	TagMsg  *string `json:"tag-msg,omitempty"`
	TagType *string `json:"tag-type,omitempty"`
}

type EmailfilterProfileSmtp struct {
	Action        *string `json:"action,omitempty"`
	Hdrip         *string `json:"hdrip,omitempty"`
	LocalOverride *string `json:"local-override,omitempty"`
	Log           *string `json:"log,omitempty"`
	LogAll        *string `json:"log-all,omitempty"`
	TagMsg        *string `json:"tag-msg,omitempty"`
	TagType       *string `json:"tag-type,omitempty"`
}

type EmailfilterProfileYahooMail struct {
	Log    *string `json:"log,omitempty"`
	LogAll *string `json:"log-all,omitempty"`
}
