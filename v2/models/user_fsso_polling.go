package models

const UserFssoPollingPath = "user/fsso-polling/"

type UserFssoPolling struct {
	Adgrp            *[]UserFssoPollingAdgrp `json:"adgrp,omitempty"`
	DefaultDomain    *string                 `json:"default-domain,omitempty"`
	Fosid            *float64                `json:"fosid,omitempty"`
	LdapServer       *string                 `json:"ldap-server,omitempty"`
	LogonHistory     *float64                `json:"logon-history,omitempty"`
	Password         *string                 `json:"password,omitempty"`
	PollingFrequency *float64                `json:"polling-frequency,omitempty"`
	Port             *float64                `json:"port,omitempty"`
	Server           *string                 `json:"server,omitempty"`
	SmbNtlmv1Auth    *string                 `json:"smb-ntlmv1-auth,omitempty"`
	Smbv1            *string                 `json:"smbv1,omitempty"`
	Status           *string                 `json:"status,omitempty"`
	User             *string                 `json:"user,omitempty"`
}

type UserFssoPollingAdgrp struct {
	Name *string `json:"name,omitempty"`
}
