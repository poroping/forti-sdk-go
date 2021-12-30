package models

const EmailfilterOptionsPath = "emailfilter/options/"

type EmailfilterOptions struct {
	DnsTimeout *int64 `json:"dns-timeout,omitempty"`
}
