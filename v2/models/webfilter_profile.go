package models

const WebfilterProfilePath = "webfilter/profile/"

type WebfilterProfile struct {
	Antiphish                 *WebfilterProfileAntiphish              `json:"antiphish,omitempty"`
	Comment                   *string                                 `json:"comment,omitempty"`
	ExtendedLog               *string                                 `json:"extended-log,omitempty"`
	FeatureSet                *string                                 `json:"feature-set,omitempty"`
	FtgdWf                    *WebfilterProfileFtgdWf                 `json:"ftgd-wf,omitempty"`
	HttpsReplacemsg           *string                                 `json:"https-replacemsg,omitempty"`
	LogAllUrl                 *string                                 `json:"log-all-url,omitempty"`
	Name                      *string                                 `json:"name,omitempty"`
	Options                   *string                                 `json:"options,omitempty"`
	Override                  *WebfilterProfileOverride               `json:"override,omitempty"`
	OvrdPerm                  *string                                 `json:"ovrd-perm,omitempty"`
	PostAction                *string                                 `json:"post-action,omitempty"`
	ReplacemsgGroup           *string                                 `json:"replacemsg-group,omitempty"`
	Web                       *WebfilterProfileWeb                    `json:"web,omitempty"`
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

const WebfilterProfileAntiphishPath = "webfilter/profile/antiphish/"

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
	MaxBodyLen        *int64                                        `json:"max-body-len,omitempty"`
	Status            *string                                       `json:"status,omitempty"`
}

const WebfilterProfileAntiphishCustomPatternsPath = "webfilter/profile/antiphish/custom-patterns/"

type WebfilterProfileAntiphishCustomPatterns struct {
	Category *string `json:"category,omitempty"`
	Pattern  *string `json:"pattern,omitempty"`
	Type     *string `json:"type,omitempty"`
}

const WebfilterProfileAntiphishInspectionEntriesPath = "webfilter/profile/antiphish/inspection-entries/"

type WebfilterProfileAntiphishInspectionEntries struct {
	Action             *string `json:"action,omitempty"`
	FortiguardCategory *string `json:"fortiguard-category,omitempty"`
	Name               *string `json:"name,omitempty"`
}

const WebfilterProfileFtgdWfPath = "webfilter/profile/ftgd-wf/"

type WebfilterProfileFtgdWf struct {
	ExemptQuota        *string                          `json:"exempt-quota,omitempty"`
	Filters            *[]WebfilterProfileFtgdWfFilters `json:"filters,omitempty"`
	MaxQuotaTimeout    *int64                           `json:"max-quota-timeout,omitempty"`
	Options            *string                          `json:"options,omitempty"`
	Ovrd               *string                          `json:"ovrd,omitempty"`
	Quota              *[]WebfilterProfileFtgdWfQuota   `json:"quota,omitempty"`
	RateCrlUrls        *string                          `json:"rate-crl-urls,omitempty"`
	RateCssUrls        *string                          `json:"rate-css-urls,omitempty"`
	RateJavascriptUrls *string                          `json:"rate-javascript-urls,omitempty"`
}

const WebfilterProfileFtgdWfFiltersPath = "webfilter/profile/ftgd-wf/filters/"

type WebfilterProfileFtgdWfFilters struct {
	Action              *string                                    `json:"action,omitempty"`
	AuthUsrGrp          *[]WebfilterProfileFtgdWfFiltersAuthUsrGrp `json:"auth-usr-grp,omitempty"`
	Category            *int64                                     `json:"category,omitempty"`
	Id                  *int64                                     `json:"id,omitempty"`
	Log                 *string                                    `json:"log,omitempty"`
	OverrideReplacemsg  *string                                    `json:"override-replacemsg,omitempty"`
	WarnDuration        *string                                    `json:"warn-duration,omitempty"`
	WarningDurationType *string                                    `json:"warning-duration-type,omitempty"`
	WarningPrompt       *string                                    `json:"warning-prompt,omitempty"`
}

const WebfilterProfileFtgdWfFiltersAuthUsrGrpPath = "webfilter/profile/ftgd-wf/filters/auth-usr-grp/"

type WebfilterProfileFtgdWfFiltersAuthUsrGrp struct {
	Name *string `json:"name,omitempty"`
}

const WebfilterProfileFtgdWfQuotaPath = "webfilter/profile/ftgd-wf/quota/"

type WebfilterProfileFtgdWfQuota struct {
	Category           *string `json:"category,omitempty"`
	Duration           *string `json:"duration,omitempty"`
	Id                 *int64  `json:"id,omitempty"`
	OverrideReplacemsg *string `json:"override-replacemsg,omitempty"`
	Type               *string `json:"type,omitempty"`
	Unit               *string `json:"unit,omitempty"`
	Value              *int64  `json:"value,omitempty"`
}

const WebfilterProfileOverridePath = "webfilter/profile/override/"

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

const WebfilterProfileOverrideOvrdUserGroupPath = "webfilter/profile/override/ovrd-user-group/"

type WebfilterProfileOverrideOvrdUserGroup struct {
	Name *string `json:"name,omitempty"`
}

const WebfilterProfileOverrideProfilePath = "webfilter/profile/override/profile/"

type WebfilterProfileOverrideProfile struct {
	Name *string `json:"name,omitempty"`
}

const WebfilterProfileWebPath = "webfilter/profile/web/"

type WebfilterProfileWeb struct {
	Allowlist         *string                            `json:"allowlist,omitempty"`
	Blacklist         *string                            `json:"blacklist,omitempty"`
	Blocklist         *string                            `json:"blocklist,omitempty"`
	BwordTable        *int64                             `json:"bword-table,omitempty"`
	BwordThreshold    *int64                             `json:"bword-threshold,omitempty"`
	ContentHeaderList *int64                             `json:"content-header-list,omitempty"`
	KeywordMatch      *[]WebfilterProfileWebKeywordMatch `json:"keyword-match,omitempty"`
	LogSearch         *string                            `json:"log-search,omitempty"`
	SafeSearch        *string                            `json:"safe-search,omitempty"`
	UrlfilterTable    *int64                             `json:"urlfilter-table,omitempty"`
	VimeoRestrict     *string                            `json:"vimeo-restrict,omitempty"`
	Whitelist         *string                            `json:"whitelist,omitempty"`
	YoutubeRestrict   *string                            `json:"youtube-restrict,omitempty"`
}

const WebfilterProfileWebKeywordMatchPath = "webfilter/profile/web/keyword-match/"

type WebfilterProfileWebKeywordMatch struct {
	Pattern *string `json:"pattern,omitempty"`
}

const WebfilterProfileWispServersPath = "webfilter/profile/wisp-servers/"

type WebfilterProfileWispServers struct {
	Name *string `json:"name,omitempty"`
}

const WebfilterProfileYoutubeChannelFilterPath = "webfilter/profile/youtube-channel-filter/"

type WebfilterProfileYoutubeChannelFilter struct {
	ChannelId *string `json:"channel-id,omitempty"`
	Comment   *string `json:"comment,omitempty"`
	Id        *int64  `json:"id,omitempty"`
}
