package models

const HardwareNicPath = "hardware/nic/"

type HardwareNic struct {
	Nic *string `json:"<nic>,omitempty"`
}
