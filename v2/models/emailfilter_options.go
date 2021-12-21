package models

const EmailfilterOptionsPath = "emailfilter/options/"

type EmailfilterOptions struct {
	DnsTimeout *float64 `json:"dns-timeout,omitempty"`
}
