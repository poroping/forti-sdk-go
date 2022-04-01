package models

const SystemAcmePath = "system/acme/"

type SystemAcme struct {
	Accounts  *[]SystemAcmeAccounts  `json:"accounts,omitempty"`
	Interface *[]SystemAcmeInterface `json:"interface,omitempty"`
	SourceIp  *string                `json:"source-ip,omitempty"`
	SourceIp6 *string                `json:"source-ip6,omitempty"`
}

type SystemAcmeAccounts struct {
	Ca_url     *string `json:"ca_url,omitempty"`
	Email      *string `json:"email,omitempty"`
	Id         *string `json:"id,omitempty"`
	Privatekey *string `json:"privatekey,omitempty"`
	Status     *string `json:"status,omitempty"`
	Url        *string `json:"url,omitempty"`
}

type SystemAcmeInterface struct {
	InterfaceName *string `json:"interface-name,omitempty"`
}
