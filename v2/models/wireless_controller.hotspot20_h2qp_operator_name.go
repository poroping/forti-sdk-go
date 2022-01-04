package models

const WirelessControllerHotspot20H2qpOperatorNamePath = "wireless-controller.hotspot20/h2qp-operator-name/"

type WirelessControllerHotspot20H2qpOperatorName struct {
	Name      *string                                                 `json:"name,omitempty"`
	ValueList *[]WirelessControllerHotspot20H2qpOperatorNameValueList `json:"value-list,omitempty"`
}

type WirelessControllerHotspot20H2qpOperatorNameValueList struct {
	Index *int64  `json:"index,omitempty"`
	Lang  *string `json:"lang,omitempty"`
	Value *string `json:"value,omitempty"`
}
