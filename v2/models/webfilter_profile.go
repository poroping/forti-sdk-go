package models

const WebfilterProfilePath = "webfilter/profile/"

type WebfilterProfile struct {
	Antiphish                 *[]WebfilterProfileAntiphish            `json:"antiphish,omitempty"`
	Comment                   *string                                 `json:"comment,omitempty"`
	ExtendedLog               *string                                 `json:"extended-log,omitempty"`
	FeatureSet                *string                                 `json:"feature-set,omitempty"`
	FileFilter                *[]WebfilterProfileFileFilter           `json:"file-filter,omitempty"`
	FtgdWf                    *[]WebfilterProfileFtgdWf               `json:"ftgd-wf,omitempty"`
	HttpsReplacemsg           *string                                 `json:"https-replacemsg,omitempty"`
	LogAllUrl                 *string                                 `json:"log-all-url,omitempty"`
	Name                      *string                                 `json:"name,omitempty"`
	Options                   *string                                 `json:"options,omitempty"`
	Override                  *[]WebfilterProfileOverride             `json:"override,omitempty"`
	OvrdPerm                  *string                                 `json:"ovrd-perm,omitempty"`
	PostAction                *string                                 `json:"post-action,omitempty"`
	ReplacemsgGroup           *string                                 `json:"replacemsg-group,omitempty"`
	Web                       *[]WebfilterProfileWeb                  `json:"web,omitempty"`
	WebAntiphishingLog        *string                                 `json:"web-antiphishing-log,omitempty"`
	WebContentLog             *string                                 `json:"web-content-log,omitempty"`
	WebExtendedAllActionLog   *string                                 `json:"web-extended-all-action-log,omitempty"`
	WebFilterActivexLog       *string                                 `json:"web-filter-activex-log,omitempty"`
	WebFilterAppletLog        *string                                 `json:"web-filter-applet-log,omitempty"`
	WebFilterCommandBlockLog  *string                                 `json:"web-filter-command-block-log,omitempty"`
	WebFilterCookieLog        *string                                 `json:"web-filter-cookie-log,omitempty"`
	WebFilterCookieRemovalLog *string                                 `json:"web-filter-cookie-removal-log,omitempty"`
	WebFilterJsLog            *string                                 `json:"web-filter-js-log,omitempty"`
	WebFilterJscriptLog       *string                                 `json:"web-filter-jscript-log,omitempty"`
	WebFilterRefererLog       *string                                 `json:"web-filter-referer-log,omitempty"`
	WebFilterUnknownLog       *string                                 `json:"web-filter-unknown-log,omitempty"`
	WebFilterVbsLog           *string                                 `json:"web-filter-vbs-log,omitempty"`
	WebFtgdErrLog             *string                                 `json:"web-ftgd-err-log,omitempty"`
	WebFtgdQuotaUsage         *string                                 `json:"web-ftgd-quota-usage,omitempty"`
	WebInvalidDomainLog       *string                                 `json:"web-invalid-domain-log,omitempty"`
	WebUrlLog                 *string                                 `json:"web-url-log,omitempty"`
	Wisp                      *string                                 `json:"wisp,omitempty"`
	WispAlgorithm             *string                                 `json:"wisp-algorithm,omitempty"`
	WispServers               *[]WebfilterProfileWispServers          `json:"wisp-servers,omitempty"`
	YoutubeChannelFilter      *[]WebfilterProfileYoutubeChannelFilter `json:"youtube-channel-filter,omitempty"`
	YoutubeChannelStatus      *string                                 `json:"youtube-channel-status,omitempty"`
}

type WebfilterProfileAntiphish struct {
	Authentication    *string                                       `json:"authentication,omitempty"`
	CheckBasicAuth    *string                                       `json:"check-basic-auth,omitempty"`
	CheckUri          *string                                       `json:"check-uri,omitempty"`
	CheckUsernameOnly *string                                       `json:"check-username-only,omitempty"`
	CustomPatterns    *[]WebfilterProfileAntiphishCustomPatterns    `json:"custom-patterns,omitempty"`
	DefaultAction     *string                                       `json:"default-action,omitempty"`
	DomainController  *string                                       `json:"domain-controller,omitempty"`
	InspectionEntries *[]WebfilterProfileAntiphishInspectionEntries `json:"inspection-entries,omitempty"`
	Ldap              *string                                       `json:"ldap,omitempty"`
	MaxBodyLen        *float64                                      `json:"max-body-len,omitempty"`
	Status            *string                                       `json:"status,omitempty"`
}

type WebfilterProfileAntiphishCustomPatterns struct {
	Category *string `json:"category,omitempty"`
	Pattern  *string `json:"pattern,omitempty"`
	Type     *string `json:"type,omitempty"`
}

type WebfilterProfileAntiphishInspectionEntries struct {
	Action             *string `json:"action,omitempty"`
	FortiguardCategory *string `json:"fortiguard-category,omitempty"`
	Name               *string `json:"name,omitempty"`
}

type WebfilterProfileFileFilter struct {
	Entries             *[]WebfilterProfileFileFilterEntries `json:"entries,omitempty"`
	Log                 *string                              `json:"log,omitempty"`
	ScanArchiveContents *string                              `json:"scan-archive-contents,omitempty"`
	Status              *string                              `json:"status,omitempty"`
}

type WebfilterProfileFileFilterEntries struct {
	Action            *string                                      `json:"action,omitempty"`
	Comment           *string                                      `json:"comment,omitempty"`
	Direction         *string                                      `json:"direction,omitempty"`
	FileType          *[]WebfilterProfileFileFilterEntriesFileType `json:"file-type,omitempty"`
	Filter            *string                                      `json:"filter,omitempty"`
	PasswordProtected *string                                      `json:"password-protected,omitempty"`
	Protocol          *string                                      `json:"protocol,omitempty"`
}

type WebfilterProfileFileFilterEntriesFileType struct {
	Name *string `json:"name,omitempty"`
}

type WebfilterProfileFtgdWf struct {
	ExemptQuota        *string                          `json:"exempt-quota,omitempty"`
	Filters            *[]WebfilterProfileFtgdWfFilters `json:"filters,omitempty"`
	MaxQuotaTimeout    *float64                         `json:"max-quota-timeout,omitempty"`
	Options            *string                          `json:"options,omitempty"`
	Ovrd               *string                          `json:"ovrd,omitempty"`
	Quota              *[]WebfilterProfileFtgdWfQuota   `json:"quota,omitempty"`
	RateCrlUrls        *string                          `json:"rate-crl-urls,omitempty"`
	RateCssUrls        *string                          `json:"rate-css-urls,omitempty"`
	RateImageUrls      *string                          `json:"rate-image-urls,omitempty"`
	RateJavascriptUrls *string                          `json:"rate-javascript-urls,omitempty"`
}

type WebfilterProfileFtgdWfFilters struct {
	Action              *string                                    `json:"action,omitempty"`
	AuthUsrGrp          *[]WebfilterProfileFtgdWfFiltersAuthUsrGrp `json:"auth-usr-grp,omitempty"`
	Category            *float64                                   `json:"category,omitempty"`
	Id                  *float64                                   `json:"id,omitempty"`
	Log                 *string                                    `json:"log,omitempty"`
	OverrideReplacemsg  *string                                    `json:"override-replacemsg,omitempty"`
	WarnDuration        *string                                    `json:"warn-duration,omitempty"`
	WarningDurationType *string                                    `json:"warning-duration-type,omitempty"`
	WarningPrompt       *string                                    `json:"warning-prompt,omitempty"`
}

type WebfilterProfileFtgdWfFiltersAuthUsrGrp struct {
	Name *string `json:"name,omitempty"`
}

type WebfilterProfileFtgdWfQuota struct {
	Category           *string  `json:"category,omitempty"`
	Duration           *string  `json:"duration,omitempty"`
	Id                 *float64 `json:"id,omitempty"`
	OverrideReplacemsg *string  `json:"override-replacemsg,omitempty"`
	Type               *string  `json:"type,omitempty"`
	Unit               *string  `json:"unit,omitempty"`
	Value              *float64 `json:"value,omitempty"`
}

type WebfilterProfileOverride struct {
	OvrdCookie       *string                                  `json:"ovrd-cookie,omitempty"`
	OvrdDur          *string                                  `json:"ovrd-dur,omitempty"`
	OvrdDurMode      *string                                  `json:"ovrd-dur-mode,omitempty"`
	OvrdScope        *string                                  `json:"ovrd-scope,omitempty"`
	OvrdUserGroup    *[]WebfilterProfileOverrideOvrdUserGroup `json:"ovrd-user-group,omitempty"`
	Profile          *[]WebfilterProfileOverrideProfile       `json:"profile,omitempty"`
	ProfileAttribute *string                                  `json:"profile-attribute,omitempty"`
	ProfileType      *string                                  `json:"profile-type,omitempty"`
}

type WebfilterProfileOverrideOvrdUserGroup struct {
	Name *string `json:"name,omitempty"`
}

type WebfilterProfileOverrideProfile struct {
	Name *string `json:"name,omitempty"`
}

type WebfilterProfileWeb struct {
	Allowlist         *string                            `json:"allowlist,omitempty"`
	Blacklist         *string                            `json:"blacklist,omitempty"`
	Blocklist         *string                            `json:"blocklist,omitempty"`
	BwordTable        *float64                           `json:"bword-table,omitempty"`
	BwordThreshold    *float64                           `json:"bword-threshold,omitempty"`
	ContentHeaderList *float64                           `json:"content-header-list,omitempty"`
	KeywordMatch      *[]WebfilterProfileWebKeywordMatch `json:"keyword-match,omitempty"`
	LogSearch         *string                            `json:"log-search,omitempty"`
	SafeSearch        *string                            `json:"safe-search,omitempty"`
	UrlfilterTable    *float64                           `json:"urlfilter-table,omitempty"`
	VimeoRestrict     *string                            `json:"vimeo-restrict,omitempty"`
	Whitelist         *string                            `json:"whitelist,omitempty"`
	YoutubeRestrict   *string                            `json:"youtube-restrict,omitempty"`
}

type WebfilterProfileWebKeywordMatch struct {
	Pattern *string `json:"pattern,omitempty"`
}

type WebfilterProfileWispServers struct {
	Name *string `json:"name,omitempty"`
}

type WebfilterProfileYoutubeChannelFilter struct {
	ChannelId *string  `json:"channel-id,omitempty"`
	Comment   *string  `json:"comment,omitempty"`
	Id        *float64 `json:"id,omitempty"`
}
