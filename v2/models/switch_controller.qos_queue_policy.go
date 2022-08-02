package models

const SwitchControllerQosQueuePolicyPath = "switch-controller.qos/queue-policy/"

type SwitchControllerQosQueuePolicy struct {
	CosQueue *[]SwitchControllerQosQueuePolicyCosQueue `json:"cos-queue,omitempty"`
	Name     *string                                   `json:"name,omitempty"`
	RateBy   *string                                   `json:"rate-by,omitempty"`
	Schedule *string                                   `json:"schedule,omitempty"`
}

const SwitchControllerQosQueuePolicyCosQueuePath = "switch-controller.qos/queue-policy/cos-queue/"

type SwitchControllerQosQueuePolicyCosQueue struct {
	Description    *string `json:"description,omitempty"`
	DropPolicy     *string `json:"drop-policy,omitempty"`
	Ecn            *string `json:"ecn,omitempty"`
	MaxRate        *int64  `json:"max-rate,omitempty"`
	MaxRatePercent *int64  `json:"max-rate-percent,omitempty"`
	MinRate        *int64  `json:"min-rate,omitempty"`
	MinRatePercent *int64  `json:"min-rate-percent,omitempty"`
	Name           *string `json:"name,omitempty"`
	Weight         *int64  `json:"weight,omitempty"`
}
