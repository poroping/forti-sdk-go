package models

const ApplicationNamePath = "application/name/"

type ApplicationName struct {
	Behavior    *string                      `json:"behavior,omitempty"`
	Category    *float64                     `json:"category,omitempty"`
	Fosid       *float64                     `json:"fosid,omitempty"`
	Metadata    *[]ApplicationNameMetadata   `json:"metadata,omitempty"`
	Name        *string                      `json:"name,omitempty"`
	Parameter   *string                      `json:"parameter,omitempty"`
	Parameters  *[]ApplicationNameParameters `json:"parameters,omitempty"`
	Popularity  *float64                     `json:"popularity,omitempty"`
	Protocol    *string                      `json:"protocol,omitempty"`
	Risk        *float64                     `json:"risk,omitempty"`
	SubCategory *float64                     `json:"sub-category,omitempty"`
	Technology  *string                      `json:"technology,omitempty"`
	Vendor      *string                      `json:"vendor,omitempty"`
	Weight      *float64                     `json:"weight,omitempty"`
}

type ApplicationNameMetadata struct {
	Id      *float64 `json:"id,omitempty"`
	Metaid  *float64 `json:"metaid,omitempty"`
	Valueid *float64 `json:"valueid,omitempty"`
}

type ApplicationNameParameters struct {
	DefaultValue *string `json:"default value,omitempty"`
	Name         *string `json:"name,omitempty"`
}
