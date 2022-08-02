package models

const SystemAcmePath = "system/acme/"

type SystemAcme struct {
	Accounts  *[]SystemAcmeAccounts  `json:"accounts,omitempty"`
	Interface *[]SystemAcmeInterface `json:"interface,omitempty"`
}

const SystemAcmeAccountsPath = "system/acme/accounts/"

type SystemAcmeAccounts struct {
	Ca_url     *string `json:"ca_url,omitempty"`
	Email      *string `json:"email,omitempty"`
	Id         *string `json:"id,omitempty"`
	Privatekey *string `json:"privatekey,omitempty"`
	Status     *string `json:"status,omitempty"`
	Url        *string `json:"url,omitempty"`
}

const SystemAcmeInterfacePath = "system/acme/interface/"

type SystemAcmeInterface struct {
	InterfaceName *string `json:"interface-name,omitempty"`
}
