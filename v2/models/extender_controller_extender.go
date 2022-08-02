package models

const ExtenderControllerExtenderPath = "extender-controller/extender/"

type ExtenderControllerExtender struct {
	Allowaccess                 *string                                     `json:"allowaccess,omitempty"`
	Authorized                  *string                                     `json:"authorized,omitempty"`
	BandwidthLimit              *int64                                      `json:"bandwidth-limit,omitempty"`
	ControllerReport            *ExtenderControllerExtenderControllerReport `json:"controller-report,omitempty"`
	Description                 *string                                     `json:"description,omitempty"`
	DeviceId                    *int64                                      `json:"device-id,omitempty"`
	EnforceBandwidth            *string                                     `json:"enforce-bandwidth,omitempty"`
	ExtName                     *string                                     `json:"ext-name,omitempty"`
	ExtensionType               *string                                     `json:"extension-type,omitempty"`
	Id                          *string                                     `json:"id,omitempty"`
	LoginPassword               *string                                     `json:"login-password,omitempty"`
	LoginPasswordChange         *string                                     `json:"login-password-change,omitempty"`
	Modem1                      *ExtenderControllerExtenderModem1           `json:"modem1,omitempty"`
	Modem2                      *ExtenderControllerExtenderModem2           `json:"modem2,omitempty"`
	Name                        *string                                     `json:"name,omitempty"`
	OverrideAllowaccess         *string                                     `json:"override-allowaccess,omitempty"`
	OverrideEnforceBandwidth    *string                                     `json:"override-enforce-bandwidth,omitempty"`
	OverrideLoginPasswordChange *string                                     `json:"override-login-password-change,omitempty"`
	Profile                     *string                                     `json:"profile,omitempty"`
	Vdom                        *int64                                      `json:"vdom,omitempty"`
	WanExtension                *ExtenderControllerExtenderWanExtension     `json:"wan-extension,omitempty"`
}

const ExtenderControllerExtenderControllerReportPath = "extender-controller/extender/controller-report/"

type ExtenderControllerExtenderControllerReport struct {
	Interval        *int64  `json:"interval,omitempty"`
	SignalThreshold *int64  `json:"signal-threshold,omitempty"`
	Status          *string `json:"status,omitempty"`
}

const ExtenderControllerExtenderModem1Path = "extender-controller/extender/modem1/"

type ExtenderControllerExtenderModem1 struct {
	AutoSwitch       *ExtenderControllerExtenderModem1AutoSwitch `json:"auto-switch,omitempty"`
	ConnStatus       *int64                                      `json:"conn-status,omitempty"`
	DefaultSim       *string                                     `json:"default-sim,omitempty"`
	Gps              *string                                     `json:"gps,omitempty"`
	Ifname           *string                                     `json:"ifname,omitempty"`
	PreferredCarrier *string                                     `json:"preferred-carrier,omitempty"`
	RedundantIntf    *string                                     `json:"redundant-intf,omitempty"`
	RedundantMode    *string                                     `json:"redundant-mode,omitempty"`
	Sim1Pin          *string                                     `json:"sim1-pin,omitempty"`
	Sim1PinCode      *string                                     `json:"sim1-pin-code,omitempty"`
	Sim2Pin          *string                                     `json:"sim2-pin,omitempty"`
	Sim2PinCode      *string                                     `json:"sim2-pin-code,omitempty"`
}

const ExtenderControllerExtenderModem1AutoSwitchPath = "extender-controller/extender/modem1/auto-switch/"

type ExtenderControllerExtenderModem1AutoSwitch struct {
	Dataplan            *string `json:"dataplan,omitempty"`
	Disconnect          *string `json:"disconnect,omitempty"`
	DisconnectPeriod    *int64  `json:"disconnect-period,omitempty"`
	DisconnectThreshold *int64  `json:"disconnect-threshold,omitempty"`
	Signal              *string `json:"signal,omitempty"`
	SwitchBack          *string `json:"switch-back,omitempty"`
	SwitchBackTime      *string `json:"switch-back-time,omitempty"`
	SwitchBackTimer     *int64  `json:"switch-back-timer,omitempty"`
}

const ExtenderControllerExtenderModem2Path = "extender-controller/extender/modem2/"

type ExtenderControllerExtenderModem2 struct {
	AutoSwitch       *ExtenderControllerExtenderModem2AutoSwitch `json:"auto-switch,omitempty"`
	ConnStatus       *int64                                      `json:"conn-status,omitempty"`
	DefaultSim       *string                                     `json:"default-sim,omitempty"`
	Gps              *string                                     `json:"gps,omitempty"`
	Ifname           *string                                     `json:"ifname,omitempty"`
	PreferredCarrier *string                                     `json:"preferred-carrier,omitempty"`
	RedundantIntf    *string                                     `json:"redundant-intf,omitempty"`
	RedundantMode    *string                                     `json:"redundant-mode,omitempty"`
	Sim1Pin          *string                                     `json:"sim1-pin,omitempty"`
	Sim1PinCode      *string                                     `json:"sim1-pin-code,omitempty"`
	Sim2Pin          *string                                     `json:"sim2-pin,omitempty"`
	Sim2PinCode      *string                                     `json:"sim2-pin-code,omitempty"`
}

const ExtenderControllerExtenderModem2AutoSwitchPath = "extender-controller/extender/modem2/auto-switch/"

type ExtenderControllerExtenderModem2AutoSwitch struct {
	Dataplan            *string `json:"dataplan,omitempty"`
	Disconnect          *string `json:"disconnect,omitempty"`
	DisconnectPeriod    *int64  `json:"disconnect-period,omitempty"`
	DisconnectThreshold *int64  `json:"disconnect-threshold,omitempty"`
	Signal              *string `json:"signal,omitempty"`
	SwitchBack          *string `json:"switch-back,omitempty"`
	SwitchBackTime      *string `json:"switch-back-time,omitempty"`
	SwitchBackTimer     *int64  `json:"switch-back-timer,omitempty"`
}

const ExtenderControllerExtenderWanExtensionPath = "extender-controller/extender/wan-extension/"

type ExtenderControllerExtenderWanExtension struct {
	Modem1Extension *string `json:"modem1-extension,omitempty"`
	Modem2Extension *string `json:"modem2-extension,omitempty"`
}
