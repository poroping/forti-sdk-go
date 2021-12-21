package models

const WirelessControllerQosProfilePath = "wireless-controller/qos-profile/"

type WirelessControllerQosProfile struct {
	BandwidthAdmissionControl *string                                  `json:"bandwidth-admission-control,omitempty"`
	BandwidthCapacity         *float64                                 `json:"bandwidth-capacity,omitempty"`
	Burst                     *string                                  `json:"burst,omitempty"`
	CallAdmissionControl      *string                                  `json:"call-admission-control,omitempty"`
	CallCapacity              *float64                                 `json:"call-capacity,omitempty"`
	Comment                   *string                                  `json:"comment,omitempty"`
	Downlink                  *float64                                 `json:"downlink,omitempty"`
	DownlinkSta               *float64                                 `json:"downlink-sta,omitempty"`
	DscpWmmBe                 *[]WirelessControllerQosProfileDscpWmmBe `json:"dscp-wmm-be,omitempty"`
	DscpWmmBk                 *[]WirelessControllerQosProfileDscpWmmBk `json:"dscp-wmm-bk,omitempty"`
	DscpWmmMapping            *string                                  `json:"dscp-wmm-mapping,omitempty"`
	DscpWmmVi                 *[]WirelessControllerQosProfileDscpWmmVi `json:"dscp-wmm-vi,omitempty"`
	DscpWmmVo                 *[]WirelessControllerQosProfileDscpWmmVo `json:"dscp-wmm-vo,omitempty"`
	Name                      *string                                  `json:"name,omitempty"`
	Uplink                    *float64                                 `json:"uplink,omitempty"`
	UplinkSta                 *float64                                 `json:"uplink-sta,omitempty"`
	Wmm                       *string                                  `json:"wmm,omitempty"`
	WmmBeDscp                 *float64                                 `json:"wmm-be-dscp,omitempty"`
	WmmBkDscp                 *float64                                 `json:"wmm-bk-dscp,omitempty"`
	WmmDscpMarking            *string                                  `json:"wmm-dscp-marking,omitempty"`
	WmmUapsd                  *string                                  `json:"wmm-uapsd,omitempty"`
	WmmViDscp                 *float64                                 `json:"wmm-vi-dscp,omitempty"`
	WmmVoDscp                 *float64                                 `json:"wmm-vo-dscp,omitempty"`
}

type WirelessControllerQosProfileDscpWmmBe struct {
	Id *float64 `json:"id,omitempty"`
}

type WirelessControllerQosProfileDscpWmmBk struct {
	Id *float64 `json:"id,omitempty"`
}

type WirelessControllerQosProfileDscpWmmVi struct {
	Id *float64 `json:"id,omitempty"`
}

type WirelessControllerQosProfileDscpWmmVo struct {
	Id *float64 `json:"id,omitempty"`
}
