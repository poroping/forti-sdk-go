package models

const ApplicationGroupPath = "application/group/"

type ApplicationGroup struct {
	Application []ApplicationGroupApplication `json:"application,omitempty"`
	Behavior    *string                       `json:"behavior,omitempty"`
	Category    []ApplicationGroupCategory    `json:"category,omitempty"`
	Comment     *string                       `json:"comment,omitempty"`
	Name        *string                       `json:"name,omitempty"`
	Popularity  *string                       `json:"popularity,omitempty"`
	Protocols   *string                       `json:"protocols,omitempty"`
	Risk        []ApplicationGroupRisk        `json:"risk,omitempty"`
	Technology  *string                       `json:"technology,omitempty"`
	Type        *string                       `json:"type,omitempty"`
	Vendor      *string                       `json:"vendor,omitempty"`
}

type ApplicationGroupApplication struct {
	Id *int64 `json:"id,omitempty"`
}

type ApplicationGroupCategory struct {
	Id *int64 `json:"id,omitempty"`
}

type ApplicationGroupRisk struct {
	Level *int64 `json:"level,omitempty"`
}
