package models

const WirelessControllerQosProfilePath = "wireless-controller/qos-profile/"

type WirelessControllerQosProfile struct {
	BandwidthAdmissionControl *string                                  `json:"bandwidth-admission-control,omitempty"`
	BandwidthCapacity         *int64                                   `json:"bandwidth-capacity,omitempty"`
	Burst                     *string                                  `json:"burst,omitempty"`
	CallAdmissionControl      *string                                  `json:"call-admission-control,omitempty"`
	CallCapacity              *int64                                   `json:"call-capacity,omitempty"`
	Comment                   *string                                  `json:"comment,omitempty"`
	Downlink                  *int64                                   `json:"downlink,omitempty"`
	DownlinkSta               *int64                                   `json:"downlink-sta,omitempty"`
	DscpWmmBe                 *[]WirelessControllerQosProfileDscpWmmBe `json:"dscp-wmm-be,omitempty"`
	DscpWmmBk                 *[]WirelessControllerQosProfileDscpWmmBk `json:"dscp-wmm-bk,omitempty"`
	DscpWmmMapping            *string                                  `json:"dscp-wmm-mapping,omitempty"`
	DscpWmmVi                 *[]WirelessControllerQosProfileDscpWmmVi `json:"dscp-wmm-vi,omitempty"`
	DscpWmmVo                 *[]WirelessControllerQosProfileDscpWmmVo `json:"dscp-wmm-vo,omitempty"`
	Name                      *string                                  `json:"name,omitempty"`
	Uplink                    *int64                                   `json:"uplink,omitempty"`
	UplinkSta                 *int64                                   `json:"uplink-sta,omitempty"`
	Wmm                       *string                                  `json:"wmm,omitempty"`
	WmmBeDscp                 *int64                                   `json:"wmm-be-dscp,omitempty"`
	WmmBkDscp                 *int64                                   `json:"wmm-bk-dscp,omitempty"`
	WmmDscpMarking            *string                                  `json:"wmm-dscp-marking,omitempty"`
	WmmUapsd                  *string                                  `json:"wmm-uapsd,omitempty"`
	WmmViDscp                 *int64                                   `json:"wmm-vi-dscp,omitempty"`
	WmmVoDscp                 *int64                                   `json:"wmm-vo-dscp,omitempty"`
}

const WirelessControllerQosProfileDscpWmmBePath = "wireless-controller/qos-profile/dscp-wmm-be/"

type WirelessControllerQosProfileDscpWmmBe struct {
	Id *int64 `json:"id,omitempty"`
}

const WirelessControllerQosProfileDscpWmmBkPath = "wireless-controller/qos-profile/dscp-wmm-bk/"

type WirelessControllerQosProfileDscpWmmBk struct {
	Id *int64 `json:"id,omitempty"`
}

const WirelessControllerQosProfileDscpWmmViPath = "wireless-controller/qos-profile/dscp-wmm-vi/"

type WirelessControllerQosProfileDscpWmmVi struct {
	Id *int64 `json:"id,omitempty"`
}

const WirelessControllerQosProfileDscpWmmVoPath = "wireless-controller/qos-profile/dscp-wmm-vo/"

type WirelessControllerQosProfileDscpWmmVo struct {
	Id *int64 `json:"id,omitempty"`
}
