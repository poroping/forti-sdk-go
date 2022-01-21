package models

const WirelessControllerLogPath = "wireless-controller/log/"

type WirelessControllerLog struct {
	AddrgrpLog    *string `json:"addrgrp-log,omitempty"`
	BleLog        *string `json:"ble-log,omitempty"`
	ClbLog        *string `json:"clb-log,omitempty"`
	DhcpStarvLog  *string `json:"dhcp-starv-log,omitempty"`
	LedSchedLog   *string `json:"led-sched-log,omitempty"`
	RadioEventLog *string `json:"radio-event-log,omitempty"`
	RogueEventLog *string `json:"rogue-event-log,omitempty"`
	StaEventLog   *string `json:"sta-event-log,omitempty"`
	StaLocateLog  *string `json:"sta-locate-log,omitempty"`
	Status        *string `json:"status,omitempty"`
	WidsLog       *string `json:"wids-log,omitempty"`
	WtpEventLog   *string `json:"wtp-event-log,omitempty"`
}
