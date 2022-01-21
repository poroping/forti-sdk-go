package models

const SystemLldpNetworkPolicyPath = "system.lldp/network-policy/"

type SystemLldpNetworkPolicy struct {
	Comment             *string                                       `json:"comment,omitempty"`
	Guest               *[]SystemLldpNetworkPolicyGuest               `json:"guest,omitempty"`
	GuestVoiceSignaling *[]SystemLldpNetworkPolicyGuestVoiceSignaling `json:"guest-voice-signaling,omitempty"`
	Name                *string                                       `json:"name,omitempty"`
	Softphone           *[]SystemLldpNetworkPolicySoftphone           `json:"softphone,omitempty"`
	StreamingVideo      *[]SystemLldpNetworkPolicyStreamingVideo      `json:"streaming-video,omitempty"`
	VideoConferencing   *[]SystemLldpNetworkPolicyVideoConferencing   `json:"video-conferencing,omitempty"`
	VideoSignaling      *[]SystemLldpNetworkPolicyVideoSignaling      `json:"video-signaling,omitempty"`
	Voice               *[]SystemLldpNetworkPolicyVoice               `json:"voice,omitempty"`
	VoiceSignaling      *[]SystemLldpNetworkPolicyVoiceSignaling      `json:"voice-signaling,omitempty"`
}

type SystemLldpNetworkPolicyGuest struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemLldpNetworkPolicyGuestVoiceSignaling struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemLldpNetworkPolicySoftphone struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemLldpNetworkPolicyStreamingVideo struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemLldpNetworkPolicyVideoConferencing struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemLldpNetworkPolicyVideoSignaling struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemLldpNetworkPolicyVoice struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemLldpNetworkPolicyVoiceSignaling struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}
