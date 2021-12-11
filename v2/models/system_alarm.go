package models

const SystemAlarmPath = "system/alarm/"

type SystemAlarm struct {
	Audible *string             `json:"audible,omitempty"`
	Groups  []SystemAlarmGroups `json:"groups,omitempty"`
	Status  *string             `json:"status,omitempty"`
}

type SystemAlarmGroups struct {
	AdminAuthFailureThreshold  *int64                                `json:"admin-auth-failure-threshold,omitempty"`
	AdminAuthLockoutThreshold  *int64                                `json:"admin-auth-lockout-threshold,omitempty"`
	DecryptionFailureThreshold *int64                                `json:"decryption-failure-threshold,omitempty"`
	EncryptionFailureThreshold *int64                                `json:"encryption-failure-threshold,omitempty"`
	FwPolicyId                 *int64                                `json:"fw-policy-id,omitempty"`
	FwPolicyIdThreshold        *int64                                `json:"fw-policy-id-threshold,omitempty"`
	FwPolicyViolations         []SystemAlarmGroupsFwPolicyViolations `json:"fw-policy-violations,omitempty"`
	Id                         *int64                                `json:"id,omitempty"`
	LogFullWarningThreshold    *int64                                `json:"log-full-warning-threshold,omitempty"`
	Period                     *int64                                `json:"period,omitempty"`
	ReplayAttemptThreshold     *int64                                `json:"replay-attempt-threshold,omitempty"`
	SelfTestFailureThreshold   *int64                                `json:"self-test-failure-threshold,omitempty"`
	UserAuthFailureThreshold   *int64                                `json:"user-auth-failure-threshold,omitempty"`
	UserAuthLockoutThreshold   *int64                                `json:"user-auth-lockout-threshold,omitempty"`
}

type SystemAlarmGroupsFwPolicyViolations struct {
	DstIp     *string `json:"dst-ip,omitempty"`
	DstPort   *int64  `json:"dst-port,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	SrcIp     *string `json:"src-ip,omitempty"`
	SrcPort   *int64  `json:"src-port,omitempty"`
	Threshold *int64  `json:"threshold,omitempty"`
}
