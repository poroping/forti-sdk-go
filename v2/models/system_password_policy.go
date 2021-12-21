package models

const SystemPasswordPolicyPath = "system/password-policy/"

type SystemPasswordPolicy struct {
	ApplyTo             *string  `json:"apply-to,omitempty"`
	Change4Characters   *string  `json:"change-4-characters,omitempty"`
	ExpireDay           *float64 `json:"expire-day,omitempty"`
	ExpireStatus        *string  `json:"expire-status,omitempty"`
	MinChangeCharacters *float64 `json:"min-change-characters,omitempty"`
	MinLowerCaseLetter  *float64 `json:"min-lower-case-letter,omitempty"`
	MinNonAlphanumeric  *float64 `json:"min-non-alphanumeric,omitempty"`
	MinNumber           *float64 `json:"min-number,omitempty"`
	MinUpperCaseLetter  *float64 `json:"min-upper-case-letter,omitempty"`
	MinimumLength       *float64 `json:"minimum-length,omitempty"`
	ReusePassword       *string  `json:"reuse-password,omitempty"`
	Status              *string  `json:"status,omitempty"`
}
