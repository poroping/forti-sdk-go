package models

const AntivirusProfilePath = "antivirus/profile/"

type AntivirusProfile struct {
	AnalyticsAcceptFiletype       *int64                               `json:"analytics-accept-filetype,omitempty"`
	AnalyticsBlFiletype           *int64                               `json:"analytics-bl-filetype,omitempty"`
	AnalyticsDb                   *string                              `json:"analytics-db,omitempty"`
	AnalyticsIgnoreFiletype       *int64                               `json:"analytics-ignore-filetype,omitempty"`
	AnalyticsMaxUpload            *int64                               `json:"analytics-max-upload,omitempty"`
	AnalyticsWlFiletype           *int64                               `json:"analytics-wl-filetype,omitempty"`
	AvBlockLog                    *string                              `json:"av-block-log,omitempty"`
	AvVirusLog                    *string                              `json:"av-virus-log,omitempty"`
	Cifs                          *AntivirusProfileCifs                `json:"cifs,omitempty"`
	Comment                       *string                              `json:"comment,omitempty"`
	ContentDisarm                 *AntivirusProfileContentDisarm       `json:"content-disarm,omitempty"`
	EmsThreatFeed                 *string                              `json:"ems-threat-feed,omitempty"`
	ExtendedLog                   *string                              `json:"extended-log,omitempty"`
	ExternalBlocklist             *[]AntivirusProfileExternalBlocklist `json:"external-blocklist,omitempty"`
	ExternalBlocklistArchiveScan  *string                              `json:"external-blocklist-archive-scan,omitempty"`
	ExternalBlocklistEnableAll    *string                              `json:"external-blocklist-enable-all,omitempty"`
	FeatureSet                    *string                              `json:"feature-set,omitempty"`
	FortiaiErrorAction            *string                              `json:"fortiai-error-action,omitempty"`
	FortiaiTimeoutAction          *string                              `json:"fortiai-timeout-action,omitempty"`
	FortindrErrorAction           *string                              `json:"fortindr-error-action,omitempty"`
	FortindrTimeoutAction         *string                              `json:"fortindr-timeout-action,omitempty"`
	FortisandboxErrorAction       *string                              `json:"fortisandbox-error-action,omitempty"`
	FortisandboxMaxUpload         *int64                               `json:"fortisandbox-max-upload,omitempty"`
	FortisandboxMode              *string                              `json:"fortisandbox-mode,omitempty"`
	FortisandboxTimeoutAction     *string                              `json:"fortisandbox-timeout-action,omitempty"`
	FtgdAnalytics                 *string                              `json:"ftgd-analytics,omitempty"`
	Ftp                           *AntivirusProfileFtp                 `json:"ftp,omitempty"`
	Http                          *AntivirusProfileHttp                `json:"http,omitempty"`
	Imap                          *AntivirusProfileImap                `json:"imap,omitempty"`
	Mapi                          *AntivirusProfileMapi                `json:"mapi,omitempty"`
	MobileMalwareDb               *string                              `json:"mobile-malware-db,omitempty"`
	NacQuar                       *AntivirusProfileNacQuar             `json:"nac-quar,omitempty"`
	Name                          *string                              `json:"name,omitempty"`
	Nntp                          *AntivirusProfileNntp                `json:"nntp,omitempty"`
	OutbreakPrevention            *AntivirusProfileOutbreakPrevention  `json:"outbreak-prevention,omitempty"`
	OutbreakPreventionArchiveScan *string                              `json:"outbreak-prevention-archive-scan,omitempty"`
	Pop3                          *AntivirusProfilePop3                `json:"pop3,omitempty"`
	ReplacemsgGroup               *string                              `json:"replacemsg-group,omitempty"`
	ScanMode                      *string                              `json:"scan-mode,omitempty"`
	Smtp                          *AntivirusProfileSmtp                `json:"smtp,omitempty"`
	Ssh                           *AntivirusProfileSsh                 `json:"ssh,omitempty"`
}

type AntivirusProfileCifs struct {
	ArchiveBlock       *string `json:"archive-block,omitempty"`
	ArchiveLog         *string `json:"archive-log,omitempty"`
	AvScan             *string `json:"av-scan,omitempty"`
	Emulator           *string `json:"emulator,omitempty"`
	ExternalBlocklist  *string `json:"external-blocklist,omitempty"`
	Fortiai            *string `json:"fortiai,omitempty"`
	Fortindr           *string `json:"fortindr,omitempty"`
	Fortisandbox       *string `json:"fortisandbox,omitempty"`
	Options            *string `json:"options,omitempty"`
	OutbreakPrevention *string `json:"outbreak-prevention,omitempty"`
	Quarantine         *string `json:"quarantine,omitempty"`
}

type AntivirusProfileContentDisarm struct {
	CoverPage               *string `json:"cover-page,omitempty"`
	DetectOnly              *string `json:"detect-only,omitempty"`
	ErrorAction             *string `json:"error-action,omitempty"`
	OfficeAction            *string `json:"office-action,omitempty"`
	OfficeDde               *string `json:"office-dde,omitempty"`
	OfficeEmbed             *string `json:"office-embed,omitempty"`
	OfficeHylink            *string `json:"office-hylink,omitempty"`
	OfficeLinked            *string `json:"office-linked,omitempty"`
	OfficeMacro             *string `json:"office-macro,omitempty"`
	OriginalFileDestination *string `json:"original-file-destination,omitempty"`
	PdfActForm              *string `json:"pdf-act-form,omitempty"`
	PdfActGotor             *string `json:"pdf-act-gotor,omitempty"`
	PdfActJava              *string `json:"pdf-act-java,omitempty"`
	PdfActLaunch            *string `json:"pdf-act-launch,omitempty"`
	PdfActMovie             *string `json:"pdf-act-movie,omitempty"`
	PdfActSound             *string `json:"pdf-act-sound,omitempty"`
	PdfEmbedfile            *string `json:"pdf-embedfile,omitempty"`
	PdfHyperlink            *string `json:"pdf-hyperlink,omitempty"`
	PdfJavacode             *string `json:"pdf-javacode,omitempty"`
}

type AntivirusProfileExternalBlocklist struct {
	Name *string `json:"name,omitempty"`
}

type AntivirusProfileFtp struct {
	ArchiveBlock       *string `json:"archive-block,omitempty"`
	ArchiveLog         *string `json:"archive-log,omitempty"`
	AvScan             *string `json:"av-scan,omitempty"`
	Emulator           *string `json:"emulator,omitempty"`
	ExternalBlocklist  *string `json:"external-blocklist,omitempty"`
	Fortiai            *string `json:"fortiai,omitempty"`
	Fortindr           *string `json:"fortindr,omitempty"`
	Fortisandbox       *string `json:"fortisandbox,omitempty"`
	Options            *string `json:"options,omitempty"`
	OutbreakPrevention *string `json:"outbreak-prevention,omitempty"`
	Quarantine         *string `json:"quarantine,omitempty"`
}

type AntivirusProfileHttp struct {
	ArchiveBlock           *string `json:"archive-block,omitempty"`
	ArchiveLog             *string `json:"archive-log,omitempty"`
	AvScan                 *string `json:"av-scan,omitempty"`
	ContentDisarm          *string `json:"content-disarm,omitempty"`
	Emulator               *string `json:"emulator,omitempty"`
	ExternalBlocklist      *string `json:"external-blocklist,omitempty"`
	Fortiai                *string `json:"fortiai,omitempty"`
	Fortindr               *string `json:"fortindr,omitempty"`
	Fortisandbox           *string `json:"fortisandbox,omitempty"`
	Options                *string `json:"options,omitempty"`
	OutbreakPrevention     *string `json:"outbreak-prevention,omitempty"`
	Quarantine             *string `json:"quarantine,omitempty"`
	UnknownContentEncoding *string `json:"unknown-content-encoding,omitempty"`
}

type AntivirusProfileImap struct {
	ArchiveBlock       *string `json:"archive-block,omitempty"`
	ArchiveLog         *string `json:"archive-log,omitempty"`
	AvScan             *string `json:"av-scan,omitempty"`
	ContentDisarm      *string `json:"content-disarm,omitempty"`
	Emulator           *string `json:"emulator,omitempty"`
	Executables        *string `json:"executables,omitempty"`
	ExternalBlocklist  *string `json:"external-blocklist,omitempty"`
	Fortiai            *string `json:"fortiai,omitempty"`
	Fortindr           *string `json:"fortindr,omitempty"`
	Fortisandbox       *string `json:"fortisandbox,omitempty"`
	Options            *string `json:"options,omitempty"`
	OutbreakPrevention *string `json:"outbreak-prevention,omitempty"`
	Quarantine         *string `json:"quarantine,omitempty"`
}

type AntivirusProfileMapi struct {
	ArchiveBlock       *string `json:"archive-block,omitempty"`
	ArchiveLog         *string `json:"archive-log,omitempty"`
	AvScan             *string `json:"av-scan,omitempty"`
	Emulator           *string `json:"emulator,omitempty"`
	Executables        *string `json:"executables,omitempty"`
	ExternalBlocklist  *string `json:"external-blocklist,omitempty"`
	Fortiai            *string `json:"fortiai,omitempty"`
	Fortindr           *string `json:"fortindr,omitempty"`
	Fortisandbox       *string `json:"fortisandbox,omitempty"`
	Options            *string `json:"options,omitempty"`
	OutbreakPrevention *string `json:"outbreak-prevention,omitempty"`
	Quarantine         *string `json:"quarantine,omitempty"`
}

type AntivirusProfileNacQuar struct {
	Expiry   *string `json:"expiry,omitempty"`
	Infected *string `json:"infected,omitempty"`
	Log      *string `json:"log,omitempty"`
}

type AntivirusProfileNntp struct {
	ArchiveBlock       *string `json:"archive-block,omitempty"`
	ArchiveLog         *string `json:"archive-log,omitempty"`
	AvScan             *string `json:"av-scan,omitempty"`
	Emulator           *string `json:"emulator,omitempty"`
	ExternalBlocklist  *string `json:"external-blocklist,omitempty"`
	Fortiai            *string `json:"fortiai,omitempty"`
	Fortindr           *string `json:"fortindr,omitempty"`
	Fortisandbox       *string `json:"fortisandbox,omitempty"`
	Options            *string `json:"options,omitempty"`
	OutbreakPrevention *string `json:"outbreak-prevention,omitempty"`
	Quarantine         *string `json:"quarantine,omitempty"`
}

type AntivirusProfileOutbreakPrevention struct {
	ExternalBlocklist *string `json:"external-blocklist,omitempty"`
	FtgdService       *string `json:"ftgd-service,omitempty"`
}

type AntivirusProfilePop3 struct {
	ArchiveBlock       *string `json:"archive-block,omitempty"`
	ArchiveLog         *string `json:"archive-log,omitempty"`
	AvScan             *string `json:"av-scan,omitempty"`
	ContentDisarm      *string `json:"content-disarm,omitempty"`
	Emulator           *string `json:"emulator,omitempty"`
	Executables        *string `json:"executables,omitempty"`
	ExternalBlocklist  *string `json:"external-blocklist,omitempty"`
	Fortiai            *string `json:"fortiai,omitempty"`
	Fortindr           *string `json:"fortindr,omitempty"`
	Fortisandbox       *string `json:"fortisandbox,omitempty"`
	Options            *string `json:"options,omitempty"`
	OutbreakPrevention *string `json:"outbreak-prevention,omitempty"`
	Quarantine         *string `json:"quarantine,omitempty"`
}

type AntivirusProfileSmtp struct {
	ArchiveBlock       *string `json:"archive-block,omitempty"`
	ArchiveLog         *string `json:"archive-log,omitempty"`
	AvScan             *string `json:"av-scan,omitempty"`
	ContentDisarm      *string `json:"content-disarm,omitempty"`
	Emulator           *string `json:"emulator,omitempty"`
	Executables        *string `json:"executables,omitempty"`
	ExternalBlocklist  *string `json:"external-blocklist,omitempty"`
	Fortiai            *string `json:"fortiai,omitempty"`
	Fortindr           *string `json:"fortindr,omitempty"`
	Fortisandbox       *string `json:"fortisandbox,omitempty"`
	Options            *string `json:"options,omitempty"`
	OutbreakPrevention *string `json:"outbreak-prevention,omitempty"`
	Quarantine         *string `json:"quarantine,omitempty"`
}

type AntivirusProfileSsh struct {
	ArchiveBlock       *string `json:"archive-block,omitempty"`
	ArchiveLog         *string `json:"archive-log,omitempty"`
	AvScan             *string `json:"av-scan,omitempty"`
	Emulator           *string `json:"emulator,omitempty"`
	ExternalBlocklist  *string `json:"external-blocklist,omitempty"`
	Fortiai            *string `json:"fortiai,omitempty"`
	Fortindr           *string `json:"fortindr,omitempty"`
	Fortisandbox       *string `json:"fortisandbox,omitempty"`
	Options            *string `json:"options,omitempty"`
	OutbreakPrevention *string `json:"outbreak-prevention,omitempty"`
	Quarantine         *string `json:"quarantine,omitempty"`
}
