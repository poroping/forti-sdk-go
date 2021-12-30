package models

const SystemPasswordPolicyPath = "system/password-policy/"

type SystemPasswordPolicy struct {
	ApplyTo             *string `json:"apply-to,omitempty"`
	Change4Characters   *string `json:"change-4-characters,omitempty"`
	ExpireDay           *int64  `json:"expire-day,omitempty"`
	ExpireStatus        *string `json:"expire-status,omitempty"`
	MinChangeCharacters *int64  `json:"min-change-characters,omitempty"`
	MinLowerCaseLetter  *int64  `json:"min-lower-case-letter,omitempty"`
	MinNonAlphanumeric  *int64  `json:"min-non-alphanumeric,omitempty"`
	MinNumber           *int64  `json:"min-number,omitempty"`
	MinUpperCaseLetter  *int64  `json:"min-upper-case-letter,omitempty"`
	MinimumLength       *int64  `json:"minimum-length,omitempty"`
	ReusePassword       *string `json:"reuse-password,omitempty"`
	Status              *string `json:"status,omitempty"`
}
