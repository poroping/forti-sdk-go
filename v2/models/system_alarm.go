package models

const SystemAlarmPath = "system/alarm/"

type SystemAlarm struct {
	Audible *string              `json:"audible,omitempty"`
	Groups  *[]SystemAlarmGroups `json:"groups,omitempty"`
	Status  *string              `json:"status,omitempty"`
}

type SystemAlarmGroups struct {
	AdminAuthFailureThreshold  *float64                               `json:"admin-auth-failure-threshold,omitempty"`
	AdminAuthLockoutThreshold  *float64                               `json:"admin-auth-lockout-threshold,omitempty"`
	DecryptionFailureThreshold *float64                               `json:"decryption-failure-threshold,omitempty"`
	EncryptionFailureThreshold *float64                               `json:"encryption-failure-threshold,omitempty"`
	FwPolicyId                 *float64                               `json:"fw-policy-id,omitempty"`
	FwPolicyIdThreshold        *float64                               `json:"fw-policy-id-threshold,omitempty"`
	FwPolicyViolations         *[]SystemAlarmGroupsFwPolicyViolations `json:"fw-policy-violations,omitempty"`
	Id                         *float64                               `json:"id,omitempty"`
	LogFullWarningThreshold    *float64                               `json:"log-full-warning-threshold,omitempty"`
	Period                     *float64                               `json:"period,omitempty"`
	ReplayAttemptThreshold     *float64                               `json:"replay-attempt-threshold,omitempty"`
	SelfTestFailureThreshold   *float64                               `json:"self-test-failure-threshold,omitempty"`
	UserAuthFailureThreshold   *float64                               `json:"user-auth-failure-threshold,omitempty"`
	UserAuthLockoutThreshold   *float64                               `json:"user-auth-lockout-threshold,omitempty"`
}

type SystemAlarmGroupsFwPolicyViolations struct {
	DstIp     *string  `json:"dst-ip,omitempty"`
	DstPort   *float64 `json:"dst-port,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	SrcIp     *string  `json:"src-ip,omitempty"`
	SrcPort   *float64 `json:"src-port,omitempty"`
	Threshold *float64 `json:"threshold,omitempty"`
}
