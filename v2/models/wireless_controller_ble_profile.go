package models

const WirelessControllerBleProfilePath = "wireless-controller/ble-profile/"

type WirelessControllerBleProfile struct {
	Advertising        *string  `json:"advertising,omitempty"`
	BeaconInterval     *float64 `json:"beacon-interval,omitempty"`
	BleScanning        *string  `json:"ble-scanning,omitempty"`
	Comment            *string  `json:"comment,omitempty"`
	EddystoneInstance  *string  `json:"eddystone-instance,omitempty"`
	EddystoneNamespace *string  `json:"eddystone-namespace,omitempty"`
	EddystoneUrl       *string  `json:"eddystone-url,omitempty"`
	IbeaconUuid        *string  `json:"ibeacon-uuid,omitempty"`
	MajorId            *float64 `json:"major-id,omitempty"`
	MinorId            *float64 `json:"minor-id,omitempty"`
	Name               *string  `json:"name,omitempty"`
	Txpower            *string  `json:"txpower,omitempty"`
}
