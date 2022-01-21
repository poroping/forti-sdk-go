package models

const SctpFilterProfilePath = "sctp-filter/profile/"

type SctpFilterProfile struct {
	Comment     *string                         `json:"comment,omitempty"`
	Name        *string                         `json:"name,omitempty"`
	PpidFilters *[]SctpFilterProfilePpidFilters `json:"ppid-filters,omitempty"`
}

type SctpFilterProfilePpidFilters struct {
	Action  *string `json:"action,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	Ppid    *int64  `json:"ppid,omitempty"`
}
