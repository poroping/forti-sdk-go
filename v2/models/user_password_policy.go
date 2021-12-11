package models

const UserPasswordPolicyPath = "user/password-policy/"

type UserPasswordPolicy struct {
	ExpireDays             *int64  `json:"expire-days,omitempty"`
	ExpiredPasswordRenewal *string `json:"expired-password-renewal,omitempty"`
	Name                   *string `json:"name,omitempty"`
	WarnDays               *int64  `json:"warn-days,omitempty"`
}
