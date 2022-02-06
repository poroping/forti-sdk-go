package models

const ExtenderControllerExtenderProfilePath = "extender-controller/extender-profile/"

type ExtenderControllerExtenderProfile struct {
	Allowaccess         *string                                        `json:"allowaccess,omitempty"`
	BandwidthLimit      *int64                                         `json:"bandwidth-limit,omitempty"`
	Cellular            *ExtenderControllerExtenderProfileCellular     `json:"cellular,omitempty"`
	EnforceBandwidth    *string                                        `json:"enforce-bandwidth,omitempty"`
	Extension           *string                                        `json:"extension,omitempty"`
	Id                  *int64                                         `json:"id,omitempty"`
	LanExtension        *ExtenderControllerExtenderProfileLanExtension `json:"lan-extension,omitempty"`
	LoginPassword       *string                                        `json:"login-password,omitempty"`
	LoginPasswordChange *string                                        `json:"login-password-change,omitempty"`
	Model               *string                                        `json:"model,omitempty"`
	Name                *string                                        `json:"name,omitempty"`
}

type ExtenderControllerExtenderProfileCellular struct {
	ControllerReport *ExtenderControllerExtenderProfileCellularControllerReport `json:"controller-report,omitempty"`
	Dataplan         *[]ExtenderControllerExtenderProfileCellularDataplan       `json:"dataplan,omitempty"`
	Modem1           *ExtenderControllerExtenderProfileCellularModem1           `json:"modem1,omitempty"`
	Modem2           *ExtenderControllerExtenderProfileCellularModem2           `json:"modem2,omitempty"`
	SmsNotification  *ExtenderControllerExtenderProfileCellularSmsNotification  `json:"sms-notification,omitempty"`
}

type ExtenderControllerExtenderProfileCellularControllerReport struct {
	Interval        *int64  `json:"interval,omitempty"`
	SignalThreshold *int64  `json:"signal-threshold,omitempty"`
	Status          *string `json:"status,omitempty"`
}

type ExtenderControllerExtenderProfileCellularDataplan struct {
	Name *string `json:"name,omitempty"`
}

type ExtenderControllerExtenderProfileCellularModem1 struct {
	AutoSwitch       *ExtenderControllerExtenderProfileCellularModem1AutoSwitch `json:"auto-switch,omitempty"`
	ConnStatus       *int64                                                     `json:"conn-status,omitempty"`
	DefaultSim       *string                                                    `json:"default-sim,omitempty"`
	Gps              *string                                                    `json:"gps,omitempty"`
	PreferredCarrier *string                                                    `json:"preferred-carrier,omitempty"`
	RedundantIntf    *string                                                    `json:"redundant-intf,omitempty"`
	RedundantMode    *string                                                    `json:"redundant-mode,omitempty"`
	Sim1Pin          *string                                                    `json:"sim1-pin,omitempty"`
	Sim1PinCode      *string                                                    `json:"sim1-pin-code,omitempty"`
	Sim2Pin          *string                                                    `json:"sim2-pin,omitempty"`
	Sim2PinCode      *string                                                    `json:"sim2-pin-code,omitempty"`
}

type ExtenderControllerExtenderProfileCellularModem1AutoSwitch struct {
	Dataplan            *string `json:"dataplan,omitempty"`
	Disconnect          *string `json:"disconnect,omitempty"`
	DisconnectPeriod    *int64  `json:"disconnect-period,omitempty"`
	DisconnectThreshold *int64  `json:"disconnect-threshold,omitempty"`
	Signal              *string `json:"signal,omitempty"`
	SwitchBack          *string `json:"switch-back,omitempty"`
	SwitchBackTime      *string `json:"switch-back-time,omitempty"`
	SwitchBackTimer     *int64  `json:"switch-back-timer,omitempty"`
}

type ExtenderControllerExtenderProfileCellularModem2 struct {
	AutoSwitch       *ExtenderControllerExtenderProfileCellularModem2AutoSwitch `json:"auto-switch,omitempty"`
	ConnStatus       *int64                                                     `json:"conn-status,omitempty"`
	DefaultSim       *string                                                    `json:"default-sim,omitempty"`
	Gps              *string                                                    `json:"gps,omitempty"`
	PreferredCarrier *string                                                    `json:"preferred-carrier,omitempty"`
	RedundantIntf    *string                                                    `json:"redundant-intf,omitempty"`
	RedundantMode    *string                                                    `json:"redundant-mode,omitempty"`
	Sim1Pin          *string                                                    `json:"sim1-pin,omitempty"`
	Sim1PinCode      *string                                                    `json:"sim1-pin-code,omitempty"`
	Sim2Pin          *string                                                    `json:"sim2-pin,omitempty"`
	Sim2PinCode      *string                                                    `json:"sim2-pin-code,omitempty"`
}

type ExtenderControllerExtenderProfileCellularModem2AutoSwitch struct {
	Dataplan            *string `json:"dataplan,omitempty"`
	Disconnect          *string `json:"disconnect,omitempty"`
	DisconnectPeriod    *int64  `json:"disconnect-period,omitempty"`
	DisconnectThreshold *int64  `json:"disconnect-threshold,omitempty"`
	Signal              *string `json:"signal,omitempty"`
	SwitchBack          *string `json:"switch-back,omitempty"`
	SwitchBackTime      *string `json:"switch-back-time,omitempty"`
	SwitchBackTimer     *int64  `json:"switch-back-timer,omitempty"`
}

type ExtenderControllerExtenderProfileCellularSmsNotification struct {
	Alert    *ExtenderControllerExtenderProfileCellularSmsNotificationAlert      `json:"alert,omitempty"`
	Receiver *[]ExtenderControllerExtenderProfileCellularSmsNotificationReceiver `json:"receiver,omitempty"`
	Status   *string                                                             `json:"status,omitempty"`
}

type ExtenderControllerExtenderProfileCellularSmsNotificationAlert struct {
	DataExhausted       *string `json:"data-exhausted,omitempty"`
	FgtBackupModeSwitch *string `json:"fgt-backup-mode-switch,omitempty"`
	LowSignalStrength   *string `json:"low-signal-strength,omitempty"`
	ModeSwitch          *string `json:"mode-switch,omitempty"`
	OsImageFallback     *string `json:"os-image-fallback,omitempty"`
	SessionDisconnect   *string `json:"session-disconnect,omitempty"`
	SystemReboot        *string `json:"system-reboot,omitempty"`
}

type ExtenderControllerExtenderProfileCellularSmsNotificationReceiver struct {
	Alert       *string `json:"alert,omitempty"`
	Name        *string `json:"name,omitempty"`
	PhoneNumber *string `json:"phone-number,omitempty"`
	Status      *string `json:"status,omitempty"`
}

type ExtenderControllerExtenderProfileLanExtension struct {
	Backhaul          *[]ExtenderControllerExtenderProfileLanExtensionBackhaul `json:"backhaul,omitempty"`
	BackhaulInterface *string                                                  `json:"backhaul-interface,omitempty"`
	BackhaulIp        *string                                                  `json:"backhaul-ip,omitempty"`
	IpsecTunnel       *string                                                  `json:"ipsec-tunnel,omitempty"`
	LinkLoadbalance   *string                                                  `json:"link-loadbalance,omitempty"`
}

type ExtenderControllerExtenderProfileLanExtensionBackhaul struct {
	Name   *string `json:"name,omitempty"`
	Port   *string `json:"port,omitempty"`
	Role   *string `json:"role,omitempty"`
	Weight *int64  `json:"weight,omitempty"`
}
