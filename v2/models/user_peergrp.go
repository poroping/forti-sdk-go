package models

const UserPeergrpPath = "user/peergrp/"

type UserPeergrp struct {
	Member *[]UserPeergrpMember `json:"member,omitempty"`
	Name   *string              `json:"name,omitempty"`
}

const UserPeergrpMemberPath = "user/peergrp/member/"

type UserPeergrpMember struct {
	Name *string `json:"name,omitempty"`
}
