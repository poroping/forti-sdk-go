package models

const ExtenderControllerDataplanPath = "extender-controller/dataplan/"

type ExtenderControllerDataplan struct {
	APN             *string `json:"APN,omitempty"`
	PDN             *string `json:"PDN,omitempty"`
	Apn             *string `json:"apn,omitempty"`
	AuthType        *string `json:"auth-type,omitempty"`
	BillingDate     *int64  `json:"billing-date,omitempty"`
	Capacity        *int64  `json:"capacity,omitempty"`
	Carrier         *string `json:"carrier,omitempty"`
	Iccid           *string `json:"iccid,omitempty"`
	ModemId         *string `json:"modem-id,omitempty"`
	MonthlyFee      *int64  `json:"monthly-fee,omitempty"`
	Name            *string `json:"name,omitempty"`
	Overage         *string `json:"overage,omitempty"`
	Password        *string `json:"password,omitempty"`
	Pdn             *string `json:"pdn,omitempty"`
	PreferredSubnet *int64  `json:"preferred-subnet,omitempty"`
	PrivateNetwork  *string `json:"private-network,omitempty"`
	SignalPeriod    *int64  `json:"signal-period,omitempty"`
	SignalThreshold *int64  `json:"signal-threshold,omitempty"`
	Slot            *string `json:"slot,omitempty"`
	Type            *string `json:"type,omitempty"`
	Username        *string `json:"username,omitempty"`
}
