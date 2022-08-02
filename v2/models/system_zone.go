package models

const SystemZonePath = "system/zone/"

type SystemZone struct {
	Description *string                `json:"description,omitempty"`
	Interface   *[]SystemZoneInterface `json:"interface,omitempty"`
	Intrazone   *string                `json:"intrazone,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Tagging     *[]SystemZoneTagging   `json:"tagging,omitempty"`
}

const SystemZoneInterfacePath = "system/zone/interface/"

type SystemZoneInterface struct {
	InterfaceName *string `json:"interface-name,omitempty"`
}

const SystemZoneTaggingPath = "system/zone/tagging/"

type SystemZoneTagging struct {
	Category *string                  `json:"category,omitempty"`
	Name     *string                  `json:"name,omitempty"`
	Tags     *[]SystemZoneTaggingTags `json:"tags,omitempty"`
}

const SystemZoneTaggingTagsPath = "system/zone/tagging/tags/"

type SystemZoneTaggingTags struct {
	Name *string `json:"name,omitempty"`
}
