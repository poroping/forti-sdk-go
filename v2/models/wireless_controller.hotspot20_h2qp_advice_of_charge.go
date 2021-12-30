package models

const WirelessControllerhotspot20H2qpAdviceOfChargePath = "wireless-controller/hotspot20/h2qp-advice-of-charge/"

type WirelessControllerhotspot20H2qpAdviceOfCharge struct {
	AocList *[]WirelessControllerhotspot20H2qpAdviceOfChargeAocList `json:"aoc-list,omitempty"`
	Name    *string                                                 `json:"name,omitempty"`
}

type WirelessControllerhotspot20H2qpAdviceOfChargeAocList struct {
	NaiRealm         *string                                                         `json:"nai-realm,omitempty"`
	NaiRealmEncoding *string                                                         `json:"nai-realm-encoding,omitempty"`
	Name             *string                                                         `json:"name,omitempty"`
	PlanInfo         *[]WirelessControllerhotspot20H2qpAdviceOfChargeAocListPlanInfo `json:"plan-info,omitempty"`
	Type             *string                                                         `json:"type,omitempty"`
}

type WirelessControllerhotspot20H2qpAdviceOfChargeAocListPlanInfo struct {
	Currency *string `json:"currency,omitempty"`
	InfoFile *string `json:"info-file,omitempty"`
	Lang     *string `json:"lang,omitempty"`
	Name     *string `json:"name,omitempty"`
}
