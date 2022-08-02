package models

const SystemFederatedUpgradePath = "system/federated-upgrade/"

type SystemFederatedUpgrade struct {
	FailureDevice *string                           `json:"failure-device,omitempty"`
	FailureReason *string                           `json:"failure-reason,omitempty"`
	NextPathIndex *int64                            `json:"next-path-index,omitempty"`
	NodeList      *[]SystemFederatedUpgradeNodeList `json:"node-list,omitempty"`
	Status        *string                           `json:"status,omitempty"`
	UpgradeId     *int64                            `json:"upgrade-id,omitempty"`
}

const SystemFederatedUpgradeNodeListPath = "system/federated-upgrade/node-list/"

type SystemFederatedUpgradeNodeList struct {
	CoordinatingFortigate *string `json:"coordinating-fortigate,omitempty"`
	DeviceType            *string `json:"device-type,omitempty"`
	Serial                *string `json:"serial,omitempty"`
	SetupTime             *string `json:"setup-time,omitempty"`
	Time                  *string `json:"time,omitempty"`
	Timing                *string `json:"timing,omitempty"`
	UpgradePath           *string `json:"upgrade-path,omitempty"`
}
