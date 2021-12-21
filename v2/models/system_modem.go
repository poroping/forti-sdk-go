package models

const SystemModemPath = "system/modem/"

type SystemModem struct {
	Action          *string  `json:"action,omitempty"`
	Altmode         *string  `json:"altmode,omitempty"`
	Authtype1       *string  `json:"authtype1,omitempty"`
	Authtype2       *string  `json:"authtype2,omitempty"`
	Authtype3       *string  `json:"authtype3,omitempty"`
	AutoDial        *string  `json:"auto-dial,omitempty"`
	ConnectTimeout  *float64 `json:"connect-timeout,omitempty"`
	DialCmd1        *string  `json:"dial-cmd1,omitempty"`
	DialCmd2        *string  `json:"dial-cmd2,omitempty"`
	DialCmd3        *string  `json:"dial-cmd3,omitempty"`
	DialOnDemand    *string  `json:"dial-on-demand,omitempty"`
	Distance        *float64 `json:"distance,omitempty"`
	DontSendCR1     *string  `json:"dont-send-CR1,omitempty"`
	DontSendCR2     *string  `json:"dont-send-CR2,omitempty"`
	DontSendCR3     *string  `json:"dont-send-CR3,omitempty"`
	ExtraInit1      *string  `json:"extra-init1,omitempty"`
	ExtraInit2      *string  `json:"extra-init2,omitempty"`
	ExtraInit3      *string  `json:"extra-init3,omitempty"`
	HolddownTimer   *float64 `json:"holddown-timer,omitempty"`
	IdleTimer       *float64 `json:"idle-timer,omitempty"`
	Interface       *string  `json:"interface,omitempty"`
	LockdownLac     *string  `json:"lockdown-lac,omitempty"`
	Mode            *string  `json:"mode,omitempty"`
	NetworkInit     *string  `json:"network-init,omitempty"`
	Passwd1         *string  `json:"passwd1,omitempty"`
	Passwd2         *string  `json:"passwd2,omitempty"`
	Passwd3         *string  `json:"passwd3,omitempty"`
	PeerModem1      *string  `json:"peer-modem1,omitempty"`
	PeerModem2      *string  `json:"peer-modem2,omitempty"`
	PeerModem3      *string  `json:"peer-modem3,omitempty"`
	Phone1          *string  `json:"phone1,omitempty"`
	Phone2          *string  `json:"phone2,omitempty"`
	Phone3          *string  `json:"phone3,omitempty"`
	PinInit         *string  `json:"pin-init,omitempty"`
	PppEchoRequest1 *string  `json:"ppp-echo-request1,omitempty"`
	PppEchoRequest2 *string  `json:"ppp-echo-request2,omitempty"`
	PppEchoRequest3 *string  `json:"ppp-echo-request3,omitempty"`
	Priority        *float64 `json:"priority,omitempty"`
	Redial          *string  `json:"redial,omitempty"`
	Reset           *float64 `json:"reset,omitempty"`
	Status          *string  `json:"status,omitempty"`
	TrafficCheck    *string  `json:"traffic-check,omitempty"`
	Username1       *string  `json:"username1,omitempty"`
	Username2       *string  `json:"username2,omitempty"`
	Username3       *string  `json:"username3,omitempty"`
	WirelessPort    *float64 `json:"wireless-port,omitempty"`
}
