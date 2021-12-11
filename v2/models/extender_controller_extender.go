package models

const ExtenderControllerExtenderPath = "extender-controller/extender/"

type ExtenderControllerExtender struct {
	Allowaccess                 *string                                  `json:"allowaccess,omitempty"`
	Authorized                  *string                                  `json:"authorized,omitempty"`
	BandwidthLimit              *int64                                   `json:"bandwidth-limit,omitempty"`
	Description                 *string                                  `json:"description,omitempty"`
	DeviceId                    *int64                                   `json:"device-id,omitempty"`
	EnforceBandwidth            *string                                  `json:"enforce-bandwidth,omitempty"`
	ExtName                     *string                                  `json:"ext-name,omitempty"`
	ExtensionType               *string                                  `json:"extension-type,omitempty"`
	Fosid                       *string                                  `json:"fosid,omitempty"`
	LoginPassword               *string                                  `json:"login-password,omitempty"`
	LoginPasswordChange         *string                                  `json:"login-password-change,omitempty"`
	Name                        *string                                  `json:"name,omitempty"`
	OverrideAllowaccess         *string                                  `json:"override-allowaccess,omitempty"`
	OverrideEnforceBandwidth    *string                                  `json:"override-enforce-bandwidth,omitempty"`
	OverrideLoginPasswordChange *string                                  `json:"override-login-password-change,omitempty"`
	Profile                     *string                                  `json:"profile,omitempty"`
	Vdom                        *int64                                   `json:"vdom,omitempty"`
	WanExtension                []ExtenderControllerExtenderWanExtension `json:"wan-extension,omitempty"`
}

type ExtenderControllerExtenderWanExtension struct {
	Modem1Extension *string `json:"modem1-extension,omitempty"`
	Modem2Extension *string `json:"modem2-extension,omitempty"`
}
