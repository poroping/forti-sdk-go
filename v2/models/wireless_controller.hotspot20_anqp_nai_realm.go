package models

const WirelessControllerHotspot20AnqpNaiRealmPath = "wireless-controller.hotspot20/anqp-nai-realm/"

type WirelessControllerHotspot20AnqpNaiRealm struct {
	NaiList *[]WirelessControllerHotspot20AnqpNaiRealmNaiList `json:"nai-list,omitempty"`
	Name    *string                                           `json:"name,omitempty"`
}

type WirelessControllerHotspot20AnqpNaiRealmNaiList struct {
	EapMethod *[]WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod `json:"eap-method,omitempty"`
	Encoding  *string                                                    `json:"encoding,omitempty"`
	NaiRealm  *string                                                    `json:"nai-realm,omitempty"`
	Name      *string                                                    `json:"name,omitempty"`
}

type WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod struct {
	AuthParam *[]WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam `json:"auth-param,omitempty"`
	Index     *int64                                                              `json:"index,omitempty"`
	Method    *string                                                             `json:"method,omitempty"`
}

type WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam struct {
	Id    *string `json:"id,omitempty"`
	Index *int64  `json:"index,omitempty"`
	Val   *string `json:"val,omitempty"`
}
