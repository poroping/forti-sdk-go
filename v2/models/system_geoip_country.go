package models

const SystemGeoipCountryPath = "system/geoip-country/"

type SystemGeoipCountry struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

//defaultfuncs
func (def *SystemGeoipCountry) defaults() {
	def.Id = ""
	def.Name = ""
}
