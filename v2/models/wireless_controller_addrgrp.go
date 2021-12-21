package models

const WirelessControllerAddrgrpPath = "wireless-controller/addrgrp/"

type WirelessControllerAddrgrp struct {
	Addresses     *[]WirelessControllerAddrgrpAddresses `json:"addresses,omitempty"`
	DefaultPolicy *string                               `json:"default-policy,omitempty"`
	Fosid         *string                               `json:"fosid,omitempty"`
}

type WirelessControllerAddrgrpAddresses struct {
	Id *string `json:"id,omitempty"`
}
