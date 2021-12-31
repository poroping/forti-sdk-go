package models

const UserAdgrpPath = "user/adgrp/"

type UserAdgrp struct {
	ConnectorSource *string `json:"connector-source,omitempty"`
	Id              *int64  `json:"id,omitempty"`
	Name            *string `json:"name,omitempty"`
	ServerName      *string `json:"server-name,omitempty"`
}
