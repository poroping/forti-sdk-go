package models

const UserFortitokenPath = "user/fortitoken/"

type UserFortitoken struct {
	ActivationCode   *string `json:"activation-code,omitempty"`
	ActivationExpire *int64  `json:"activation-expire,omitempty"`
	Comments         *string `json:"comments,omitempty"`
	License          *string `json:"license,omitempty"`
	OsVer            *string `json:"os-ver,omitempty"`
	RegId            *string `json:"reg-id,omitempty"`
	SerialNumber     *string `json:"serial-number,omitempty"`
	Status           *string `json:"status,omitempty"`
}
