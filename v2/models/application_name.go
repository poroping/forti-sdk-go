package models

const ApplicationNamePath = "application/name/"

type ApplicationName struct {
	Behavior    *string                      `json:"behavior,omitempty"`
	Category    *int64                       `json:"category,omitempty"`
	Fosid       *int64                       `json:"fosid,omitempty"`
	Metadata    *[]ApplicationNameMetadata   `json:"metadata,omitempty"`
	Name        *string                      `json:"name,omitempty"`
	Parameter   *string                      `json:"parameter,omitempty"`
	Parameters  *[]ApplicationNameParameters `json:"parameters,omitempty"`
	Popularity  *int64                       `json:"popularity,omitempty"`
	Protocol    *string                      `json:"protocol,omitempty"`
	Risk        *int64                       `json:"risk,omitempty"`
	SubCategory *int64                       `json:"sub-category,omitempty"`
	Technology  *string                      `json:"technology,omitempty"`
	Vendor      *string                      `json:"vendor,omitempty"`
	Weight      *int64                       `json:"weight,omitempty"`
}

type ApplicationNameMetadata struct {
	Id      *int64 `json:"id,omitempty"`
	Metaid  *int64 `json:"metaid,omitempty"`
	Valueid *int64 `json:"valueid,omitempty"`
}

type ApplicationNameParameters struct {
	DefaultValue *string `json:"default value,omitempty"`
	Name         *string `json:"name,omitempty"`
}
