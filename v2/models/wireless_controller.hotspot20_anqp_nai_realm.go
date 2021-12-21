package models

const WirelessControllerhotspot20AnqpNaiRealmPath = "wireless-controller.hotspot20/anqp-nai-realm/"

type WirelessControllerhotspot20AnqpNaiRealm struct {
	NaiList *[]WirelessControllerhotspot20AnqpNaiRealmNaiList `json:"nai-list,omitempty"`
	Name    *string                                           `json:"name,omitempty"`
}

type WirelessControllerhotspot20AnqpNaiRealmNaiList struct {
	EapMethod *[]WirelessControllerhotspot20AnqpNaiRealmNaiListEapMethod `json:"eap-method,omitempty"`
	Encoding  *string                                                    `json:"encoding,omitempty"`
	NaiRealm  *string                                                    `json:"nai-realm,omitempty"`
	Name      *string                                                    `json:"name,omitempty"`
}

type WirelessControllerhotspot20AnqpNaiRealmNaiListEapMethod struct {
	AuthParam *[]WirelessControllerhotspot20AnqpNaiRealmNaiListEapMethodAuthParam `json:"auth-param,omitempty"`
	Index     *int64                                                              `json:"index,omitempty"`
	Method    *string                                                             `json:"method,omitempty"`
}

type WirelessControllerhotspot20AnqpNaiRealmNaiListEapMethodAuthParam struct {
	Id    *string `json:"id,omitempty"`
	Index *int64  `json:"index,omitempty"`
	Val   *string `json:"val,omitempty"`
}
