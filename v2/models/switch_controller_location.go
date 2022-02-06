package models

const SwitchControllerLocationPath = "switch-controller/location/"

type SwitchControllerLocation struct {
	AddressCivic *SwitchControllerLocationAddressCivic `json:"address-civic,omitempty"`
	Coordinates  *SwitchControllerLocationCoordinates  `json:"coordinates,omitempty"`
	ElinNumber   *SwitchControllerLocationElinNumber   `json:"elin-number,omitempty"`
	Name         *string                               `json:"name,omitempty"`
}

type SwitchControllerLocationAddressCivic struct {
	Additional         *string `json:"additional,omitempty"`
	AdditionalCode     *string `json:"additional-code,omitempty"`
	Block              *string `json:"block,omitempty"`
	BranchRoad         *string `json:"branch-road,omitempty"`
	Building           *string `json:"building,omitempty"`
	City               *string `json:"city,omitempty"`
	CityDivision       *string `json:"city-division,omitempty"`
	Country            *string `json:"country,omitempty"`
	CountrySubdivision *string `json:"country-subdivision,omitempty"`
	County             *string `json:"county,omitempty"`
	Direction          *string `json:"direction,omitempty"`
	Floor              *string `json:"floor,omitempty"`
	Landmark           *string `json:"landmark,omitempty"`
	Language           *string `json:"language,omitempty"`
	Name               *string `json:"name,omitempty"`
	Number             *string `json:"number,omitempty"`
	NumberSuffix       *string `json:"number-suffix,omitempty"`
	ParentKey          *string `json:"parent-key,omitempty"`
	PlaceType          *string `json:"place-type,omitempty"`
	PostOfficeBox      *string `json:"post-office-box,omitempty"`
	PostalCommunity    *string `json:"postal-community,omitempty"`
	PrimaryRoad        *string `json:"primary-road,omitempty"`
	RoadSection        *string `json:"road-section,omitempty"`
	Room               *string `json:"room,omitempty"`
	Script             *string `json:"script,omitempty"`
	Seat               *string `json:"seat,omitempty"`
	Street             *string `json:"street,omitempty"`
	StreetNamePostMod  *string `json:"street-name-post-mod,omitempty"`
	StreetNamePreMod   *string `json:"street-name-pre-mod,omitempty"`
	StreetSuffix       *string `json:"street-suffix,omitempty"`
	SubBranchRoad      *string `json:"sub-branch-road,omitempty"`
	TrailingStrSuffix  *string `json:"trailing-str-suffix,omitempty"`
	Unit               *string `json:"unit,omitempty"`
	Zip                *string `json:"zip,omitempty"`
}

type SwitchControllerLocationCoordinates struct {
	Altitude     *string `json:"altitude,omitempty"`
	AltitudeUnit *string `json:"altitude-unit,omitempty"`
	Datum        *string `json:"datum,omitempty"`
	Latitude     *string `json:"latitude,omitempty"`
	Longitude    *string `json:"longitude,omitempty"`
	ParentKey    *string `json:"parent-key,omitempty"`
}

type SwitchControllerLocationElinNumber struct {
	ElinNum   *string `json:"elin-num,omitempty"`
	ParentKey *string `json:"parent-key,omitempty"`
}
