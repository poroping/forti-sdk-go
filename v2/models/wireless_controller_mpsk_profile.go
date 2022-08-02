package models

const WirelessControllerMpskProfilePath = "wireless-controller/mpsk-profile/"

type WirelessControllerMpskProfile struct {
	MpskConcurrentClients *int64                                    `json:"mpsk-concurrent-clients,omitempty"`
	MpskGroup             *[]WirelessControllerMpskProfileMpskGroup `json:"mpsk-group,omitempty"`
	Name                  *string                                   `json:"name,omitempty"`
}

const WirelessControllerMpskProfileMpskGroupPath = "wireless-controller/mpsk-profile/mpsk-group/"

type WirelessControllerMpskProfileMpskGroup struct {
	MpskKey  *[]WirelessControllerMpskProfileMpskGroupMpskKey `json:"mpsk-key,omitempty"`
	Name     *string                                          `json:"name,omitempty"`
	VlanId   *int64                                           `json:"vlan-id,omitempty"`
	VlanType *string                                          `json:"vlan-type,omitempty"`
}

const WirelessControllerMpskProfileMpskGroupMpskKeyPath = "wireless-controller/mpsk-profile/mpsk-group/mpsk-key/"

type WirelessControllerMpskProfileMpskGroupMpskKey struct {
	Comment                   *string                                                       `json:"comment,omitempty"`
	ConcurrentClientLimitType *string                                                       `json:"concurrent-client-limit-type,omitempty"`
	ConcurrentClients         *int64                                                        `json:"concurrent-clients,omitempty"`
	Mac                       *string                                                       `json:"mac,omitempty"`
	MpskSchedules             *[]WirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules `json:"mpsk-schedules,omitempty"`
	Name                      *string                                                       `json:"name,omitempty"`
	Passphrase                *string                                                       `json:"passphrase,omitempty"`
}

const WirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedulesPath = "wireless-controller/mpsk-profile/mpsk-group/mpsk-key/mpsk-schedules/"

type WirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules struct {
	Name *string `json:"name,omitempty"`
}
