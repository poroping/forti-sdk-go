package models

const SystemautoupdateTunnelingPath = "system.autoupdate/tunneling/"

type SystemautoupdateTunneling struct {
	Address  *string  `json:"address,omitempty"`
	Password *string  `json:"password,omitempty"`
	Port     *float64 `json:"port,omitempty"`
	Status   *string  `json:"status,omitempty"`
	Username *string  `json:"username,omitempty"`
}
