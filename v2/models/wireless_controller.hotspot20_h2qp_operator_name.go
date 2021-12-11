package models

const WirelessControllerhotspot20H2qpOperatorNamePath = "wireless-controller.hotspot20/h2qp-operator-name/"

type WirelessControllerhotspot20H2qpOperatorName struct {
	Name      *string                                                `json:"name,omitempty"`
	ValueList []WirelessControllerhotspot20H2qpOperatorNameValueList `json:"value-list,omitempty"`
}

type WirelessControllerhotspot20H2qpOperatorNameValueList struct {
	Index *int64  `json:"index,omitempty"`
	Lang  *string `json:"lang,omitempty"`
	Value *string `json:"value,omitempty"`
}
