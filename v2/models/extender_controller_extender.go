package models

const ExtenderControllerExtenderPath = "extender-controller/extender/"

type ExtenderControllerExtender struct {
	AaaSharedSecret             *string                                       `json:"aaa-shared-secret,omitempty"`
	AccessPointName             *string                                       `json:"access-point-name,omitempty"`
	Admin                       *string                                       `json:"admin,omitempty"`
	Allowaccess                 *string                                       `json:"allowaccess,omitempty"`
	AtDialScript                *string                                       `json:"at-dial-script,omitempty"`
	Authorized                  *string                                       `json:"authorized,omitempty"`
	BandwidthLimit              *float64                                      `json:"bandwidth-limit,omitempty"`
	BillingStartDay             *float64                                      `json:"billing-start-day,omitempty"`
	CdmaAaaSpi                  *string                                       `json:"cdma-aaa-spi,omitempty"`
	CdmaHaSpi                   *string                                       `json:"cdma-ha-spi,omitempty"`
	CdmaNai                     *string                                       `json:"cdma-nai,omitempty"`
	ConnStatus                  *float64                                      `json:"conn-status,omitempty"`
	ControllerReport            *[]ExtenderControllerExtenderControllerReport `json:"controller-report,omitempty"`
	Description                 *string                                       `json:"description,omitempty"`
	DeviceId                    *float64                                      `json:"device-id,omitempty"`
	DialMode                    *string                                       `json:"dial-mode,omitempty"`
	DialStatus                  *float64                                      `json:"dial-status,omitempty"`
	EnforceBandwidth            *string                                       `json:"enforce-bandwidth,omitempty"`
	ExtName                     *string                                       `json:"ext-name,omitempty"`
	ExtensionType               *string                                       `json:"extension-type,omitempty"`
	HaSharedSecret              *string                                       `json:"ha-shared-secret,omitempty"`
	Fosid                       *string                                       `json:"fosid,omitempty"`
	Ifname                      *string                                       `json:"ifname,omitempty"`
	InitiatedUpdate             *string                                       `json:"initiated-update,omitempty"`
	LoginPassword               *string                                       `json:"login-password,omitempty"`
	LoginPasswordChange         *string                                       `json:"login-password-change,omitempty"`
	Mode                        *string                                       `json:"mode,omitempty"`
	ModemPasswd                 *string                                       `json:"modem-passwd,omitempty"`
	ModemType                   *string                                       `json:"modem-type,omitempty"`
	Modem1                      *[]ExtenderControllerExtenderModem1           `json:"modem1,omitempty"`
	Modem2                      *[]ExtenderControllerExtenderModem2           `json:"modem2,omitempty"`
	MultiMode                   *string                                       `json:"multi-mode,omitempty"`
	Name                        *string                                       `json:"name,omitempty"`
	OverrideAllowaccess         *string                                       `json:"override-allowaccess,omitempty"`
	OverrideEnforceBandwidth    *string                                       `json:"override-enforce-bandwidth,omitempty"`
	OverrideLoginPasswordChange *string                                       `json:"override-login-password-change,omitempty"`
	PppAuthProtocol             *string                                       `json:"ppp-auth-protocol,omitempty"`
	PppEchoRequest              *string                                       `json:"ppp-echo-request,omitempty"`
	PppPassword                 *string                                       `json:"ppp-password,omitempty"`
	PppUsername                 *string                                       `json:"ppp-username,omitempty"`
	PrimaryHa                   *string                                       `json:"primary-ha,omitempty"`
	Profile                     *string                                       `json:"profile,omitempty"`
	QuotaLimitMb                *float64                                      `json:"quota-limit-mb,omitempty"`
	Redial                      *string                                       `json:"redial,omitempty"`
	RedundantIntf               *string                                       `json:"redundant-intf,omitempty"`
	Roaming                     *string                                       `json:"roaming,omitempty"`
	Role                        *string                                       `json:"role,omitempty"`
	SecondaryHa                 *string                                       `json:"secondary-ha,omitempty"`
	SimPin                      *string                                       `json:"sim-pin,omitempty"`
	Vdom                        *float64                                      `json:"vdom,omitempty"`
	WanExtension                *[]ExtenderControllerExtenderWanExtension     `json:"wan-extension,omitempty"`
	WimaxAuthProtocol           *string                                       `json:"wimax-auth-protocol,omitempty"`
	WimaxCarrier                *string                                       `json:"wimax-carrier,omitempty"`
	WimaxRealm                  *string                                       `json:"wimax-realm,omitempty"`
}

type ExtenderControllerExtenderControllerReport struct {
	Interval        *float64 `json:"interval,omitempty"`
	SignalThreshold *float64 `json:"signal-threshold,omitempty"`
	Status          *string  `json:"status,omitempty"`
}

type ExtenderControllerExtenderModem1 struct {
	AutoSwitch       *[]ExtenderControllerExtenderModem1AutoSwitch `json:"auto-switch,omitempty"`
	ConnStatus       *float64                                      `json:"conn-status,omitempty"`
	DefaultSim       *string                                       `json:"default-sim,omitempty"`
	Gps              *string                                       `json:"gps,omitempty"`
	Ifname           *string                                       `json:"ifname,omitempty"`
	PreferredCarrier *string                                       `json:"preferred-carrier,omitempty"`
	RedundantIntf    *string                                       `json:"redundant-intf,omitempty"`
	RedundantMode    *string                                       `json:"redundant-mode,omitempty"`
	Sim1Pin          *string                                       `json:"sim1-pin,omitempty"`
	Sim1PinCode      *string                                       `json:"sim1-pin-code,omitempty"`
	Sim2Pin          *string                                       `json:"sim2-pin,omitempty"`
	Sim2PinCode      *string                                       `json:"sim2-pin-code,omitempty"`
}

type ExtenderControllerExtenderModem1AutoSwitch struct {
	Dataplan            *string  `json:"dataplan,omitempty"`
	Disconnect          *string  `json:"disconnect,omitempty"`
	DisconnectPeriod    *float64 `json:"disconnect-period,omitempty"`
	DisconnectThreshold *float64 `json:"disconnect-threshold,omitempty"`
	Signal              *string  `json:"signal,omitempty"`
	SwitchBack          *string  `json:"switch-back,omitempty"`
	SwitchBackTime      *string  `json:"switch-back-time,omitempty"`
	SwitchBackTimer     *float64 `json:"switch-back-timer,omitempty"`
}

type ExtenderControllerExtenderModem2 struct {
	AutoSwitch       *[]ExtenderControllerExtenderModem2AutoSwitch `json:"auto-switch,omitempty"`
	ConnStatus       *float64                                      `json:"conn-status,omitempty"`
	DefaultSim       *string                                       `json:"default-sim,omitempty"`
	Gps              *string                                       `json:"gps,omitempty"`
	Ifname           *string                                       `json:"ifname,omitempty"`
	PreferredCarrier *string                                       `json:"preferred-carrier,omitempty"`
	RedundantIntf    *string                                       `json:"redundant-intf,omitempty"`
	RedundantMode    *string                                       `json:"redundant-mode,omitempty"`
	Sim1Pin          *string                                       `json:"sim1-pin,omitempty"`
	Sim1PinCode      *string                                       `json:"sim1-pin-code,omitempty"`
	Sim2Pin          *string                                       `json:"sim2-pin,omitempty"`
	Sim2PinCode      *string                                       `json:"sim2-pin-code,omitempty"`
}

type ExtenderControllerExtenderModem2AutoSwitch struct {
	Dataplan            *string  `json:"dataplan,omitempty"`
	Disconnect          *string  `json:"disconnect,omitempty"`
	DisconnectPeriod    *float64 `json:"disconnect-period,omitempty"`
	DisconnectThreshold *float64 `json:"disconnect-threshold,omitempty"`
	Signal              *string  `json:"signal,omitempty"`
	SwitchBack          *string  `json:"switch-back,omitempty"`
	SwitchBackTime      *string  `json:"switch-back-time,omitempty"`
	SwitchBackTimer     *float64 `json:"switch-back-timer,omitempty"`
}

type ExtenderControllerExtenderWanExtension struct {
	Modem1Extension *string `json:"modem1-extension,omitempty"`
	Modem2Extension *string `json:"modem2-extension,omitempty"`
}
