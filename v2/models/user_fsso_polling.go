package models

const UserFssoPollingPath = "user/fsso-polling/"

type UserFssoPolling struct {
	Adgrp            *[]UserFssoPollingAdgrp `json:"adgrp,omitempty"`
	DefaultDomain    *string                 `json:"default-domain,omitempty"`
	Id               *int64                  `json:"id,omitempty"`
	LdapServer       *string                 `json:"ldap-server,omitempty"`
	LogonHistory     *int64                  `json:"logon-history,omitempty"`
	Password         *string                 `json:"password,omitempty"`
	PollingFrequency *int64                  `json:"polling-frequency,omitempty"`
	Port             *int64                  `json:"port,omitempty"`
	Server           *string                 `json:"server,omitempty"`
	SmbNtlmv1Auth    *string                 `json:"smb-ntlmv1-auth,omitempty"`
	Smbv1            *string                 `json:"smbv1,omitempty"`
	Status           *string                 `json:"status,omitempty"`
	User             *string                 `json:"user,omitempty"`
}

const UserFssoPollingAdgrpPath = "user/fsso-polling/adgrp/"

type UserFssoPollingAdgrp struct {
	Name *string `json:"name,omitempty"`
}
