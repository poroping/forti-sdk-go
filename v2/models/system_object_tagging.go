package models

const SystemObjectTaggingPath = "system/object-tagging/"

type SystemObjectTagging struct {
	Address   *string                    `json:"address,omitempty"`
	Category  *string                    `json:"category,omitempty"`
	Color     *int64                     `json:"color,omitempty"`
	Device    *string                    `json:"device,omitempty"`
	Interface *string                    `json:"interface,omitempty"`
	Multiple  *string                    `json:"multiple,omitempty"`
	Tags      *[]SystemObjectTaggingTags `json:"tags,omitempty"`
}

type SystemObjectTaggingTags struct {
	Name *string `json:"name,omitempty"`
}
