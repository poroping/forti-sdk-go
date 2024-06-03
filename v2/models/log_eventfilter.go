package models

const LogEventfilterPath = "log/eventfilter/"

type LogEventfilter struct {
	Cifs             *string `json:"cifs,omitempty"`
	Connector        *string `json:"connector,omitempty"`
	Endpoint         *string `json:"endpoint,omitempty"`
	Event            *string `json:"event,omitempty"`
	Fortiextender    *string `json:"fortiextender,omitempty"`
	Ha               *string `json:"ha,omitempty"`
	RestApi          *string `json:"rest-api,omitempty"`
	Router           *string `json:"router,omitempty"`
	Sdwan            *string `json:"sdwan,omitempty"`
	SecurityRating   *string `json:"security-rating,omitempty"`
	SwitchController *string `json:"switch-controller,omitempty"`
	System           *string `json:"system,omitempty"`
	User             *string `json:"user,omitempty"`
	Vpn              *string `json:"vpn,omitempty"`
	WanOpt           *string `json:"wan-opt,omitempty"`
	Webproxy         *string `json:"webproxy,omitempty"`
	WirelessActivity *string `json:"wireless-activity,omitempty"`
}
