package models

const SystemGeoipCountryPath = "system/geoip-country/"

type SystemGeoipCountry struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Set SystemGeoipCountry values to defaults
func (def *SystemGeoipCountry) Defaults() {
	def.Id = stringPtr("")
	def.Name = stringPtr("")
}
