package models

const SystemFipsCcPath = "system/fips-cc/"

type SystemFipsCc struct {
	EntropyToken          *string  `json:"entropy-token,omitempty"`
	KeyGenerationSelfTest *string  `json:"key-generation-self-test,omitempty"`
	SelfTestPeriod        *float64 `json:"self-test-period,omitempty"`
	Status                *string  `json:"status,omitempty"`
}
