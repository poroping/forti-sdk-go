package models

const WirelessControllerBleProfilePath = "wireless-controller/ble-profile/"

type WirelessControllerBleProfile struct {
	Advertising        *string `json:"advertising,omitempty"`
	BeaconInterval     *int64  `json:"beacon-interval,omitempty"`
	BleScanning        *string `json:"ble-scanning,omitempty"`
	Comment            *string `json:"comment,omitempty"`
	EddystoneInstance  *string `json:"eddystone-instance,omitempty"`
	EddystoneNamespace *string `json:"eddystone-namespace,omitempty"`
	EddystoneUrl       *string `json:"eddystone-url,omitempty"`
	IbeaconUuid        *string `json:"ibeacon-uuid,omitempty"`
	MajorId            *int64  `json:"major-id,omitempty"`
	MinorId            *int64  `json:"minor-id,omitempty"`
	Name               *string `json:"name,omitempty"`
	Txpower            *string `json:"txpower,omitempty"`
}
