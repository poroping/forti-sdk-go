package models

const UserAdgrpPath = "user/adgrp/"

type UserAdgrp struct {
	ConnectorSource *string  `json:"connector-source,omitempty"`
	Fosid           *float64 `json:"fosid,omitempty"`
	Name            *string  `json:"name,omitempty"`
	ServerName      *string  `json:"server-name,omitempty"`
}
