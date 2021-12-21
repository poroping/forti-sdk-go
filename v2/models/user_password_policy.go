package models

const UserPasswordPolicyPath = "user/password-policy/"

type UserPasswordPolicy struct {
	ExpireDays             *float64 `json:"expire-days,omitempty"`
	ExpiredPasswordRenewal *string  `json:"expired-password-renewal,omitempty"`
	Name                   *string  `json:"name,omitempty"`
	WarnDays               *float64 `json:"warn-days,omitempty"`
}
