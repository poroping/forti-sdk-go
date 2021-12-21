package models

const CifsProfilePath = "cifs/profile/"

type CifsProfile struct {
	DomainController     *string                    `json:"domain-controller,omitempty"`
	FileFilter           *[]CifsProfileFileFilter   `json:"file-filter,omitempty"`
	Name                 *string                    `json:"name,omitempty"`
	ServerCredentialType *string                    `json:"server-credential-type,omitempty"`
	ServerKeytab         *[]CifsProfileServerKeytab `json:"server-keytab,omitempty"`
}

type CifsProfileFileFilter struct {
	Entries *[]CifsProfileFileFilterEntries `json:"entries,omitempty"`
	Log     *string                         `json:"log,omitempty"`
	Status  *string                         `json:"status,omitempty"`
}

type CifsProfileFileFilterEntries struct {
	Action    *string                                 `json:"action,omitempty"`
	Comment   *string                                 `json:"comment,omitempty"`
	Direction *string                                 `json:"direction,omitempty"`
	FileType  *[]CifsProfileFileFilterEntriesFileType `json:"file-type,omitempty"`
	Filter    *string                                 `json:"filter,omitempty"`
}

type CifsProfileFileFilterEntriesFileType struct {
	Name *string `json:"name,omitempty"`
}

type CifsProfileServerKeytab struct {
	Keytab    *string `json:"keytab,omitempty"`
	Principal *string `json:"principal,omitempty"`
}
