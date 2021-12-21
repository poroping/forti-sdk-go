package models

const IpsCustomPath = "ips/custom/"

type IpsCustom struct {
	Action      *string  `json:"action,omitempty"`
	Application *string  `json:"application,omitempty"`
	Comment     *string  `json:"comment,omitempty"`
	Location    *string  `json:"location,omitempty"`
	Log         *string  `json:"log,omitempty"`
	LogPacket   *string  `json:"log-packet,omitempty"`
	Os          *string  `json:"os,omitempty"`
	Protocol    *string  `json:"protocol,omitempty"`
	RuleId      *float64 `json:"rule-id,omitempty"`
	Severity    *string  `json:"severity,omitempty"`
	Signature   *string  `json:"signature,omitempty"`
	Status      *string  `json:"status,omitempty"`
	Tag         *string  `json:"tag,omitempty"`
}
