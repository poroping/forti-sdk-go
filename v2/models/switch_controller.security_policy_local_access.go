package models

const SwitchControllerSecurityPolicyLocalAccessPath = "switch-controller.security-policy/local-access/"

type SwitchControllerSecurityPolicyLocalAccess struct {
	InternalAllowaccess *string `json:"internal-allowaccess,omitempty"`
	MgmtAllowaccess     *string `json:"mgmt-allowaccess,omitempty"`
	Name                *string `json:"name,omitempty"`
}
