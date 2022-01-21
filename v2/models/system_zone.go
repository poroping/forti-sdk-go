package models

const SystemZonePath = "system/zone/"

type SystemZone struct {
	Description *string                `json:"description,omitempty"`
	Interface   *[]SystemZoneInterface `json:"interface,omitempty"`
	Intrazone   *string                `json:"intrazone,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Tagging     *[]SystemZoneTagging   `json:"tagging,omitempty"`
}

type SystemZoneInterface struct {
	InterfaceName *string `json:"interface-name,omitempty"`
}

type SystemZoneTagging struct {
	Category *string                  `json:"category,omitempty"`
	Name     *string                  `json:"name,omitempty"`
	Tags     *[]SystemZoneTaggingTags `json:"tags,omitempty"`
}

type SystemZoneTaggingTags struct {
	Name *string `json:"name,omitempty"`
}
