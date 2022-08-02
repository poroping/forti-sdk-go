package models

const ApplicationNamePath = "application/name/"

type ApplicationName struct {
	Behavior   *string                      `json:"behavior,omitempty"`
	Category   *int64                       `json:"category,omitempty"`
	Id         *int64                       `json:"id,omitempty"`
	Metadata   *[]ApplicationNameMetadata   `json:"metadata,omitempty"`
	Name       *string                      `json:"name,omitempty"`
	Parameters *[]ApplicationNameParameters `json:"parameters,omitempty"`
	Popularity *int64                       `json:"popularity,omitempty"`
	Protocol   *string                      `json:"protocol,omitempty"`
	Risk       *int64                       `json:"risk,omitempty"`
	Technology *string                      `json:"technology,omitempty"`
	Vendor     *string                      `json:"vendor,omitempty"`
	Weight     *int64                       `json:"weight,omitempty"`
}

const ApplicationNameMetadataPath = "application/name/metadata/"

type ApplicationNameMetadata struct {
	Id      *int64 `json:"id,omitempty"`
	Metaid  *int64 `json:"metaid,omitempty"`
	Valueid *int64 `json:"valueid,omitempty"`
}

// Set ApplicationNameMetadata values to defaults
func (def *ApplicationNameMetadata) Defaults() {
	def.Id = intPtr(0)
	def.Metaid = intPtr(0)
	def.Valueid = intPtr(0)
}

const ApplicationNameParametersPath = "application/name/parameters/"

type ApplicationNameParameters struct {
	DefaultValue *string `json:"default value,omitempty"`
	Name         *string `json:"name,omitempty"`
}

// Set ApplicationNameParameters values to defaults
func (def *ApplicationNameParameters) Defaults() {
	def.DefaultValue = stringPtr("")
	def.Name = stringPtr("")
}

// Set ApplicationName values to defaults
func (def *ApplicationName) Defaults() {
	def.Behavior = stringPtr("")
	def.Category = intPtr(0)
	def.Id = intPtr(0)
	def.Name = stringPtr("")
	def.Popularity = intPtr(0)
	def.Protocol = stringPtr("")
	def.Risk = intPtr(0)
	def.Technology = stringPtr("")
	def.Vendor = stringPtr("")
	def.Weight = intPtr(0)
}
