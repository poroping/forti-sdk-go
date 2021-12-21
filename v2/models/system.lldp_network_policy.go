package models

const SystemlldpNetworkPolicyPath = "system.lldp/network-policy/"

type SystemlldpNetworkPolicy struct {
	Comment             *string                                       `json:"comment,omitempty"`
	Guest               *[]SystemlldpNetworkPolicyGuest               `json:"guest,omitempty"`
	GuestVoiceSignaling *[]SystemlldpNetworkPolicyGuestVoiceSignaling `json:"guest-voice-signaling,omitempty"`
	Name                *string                                       `json:"name,omitempty"`
	Softphone           *[]SystemlldpNetworkPolicySoftphone           `json:"softphone,omitempty"`
	StreamingVideo      *[]SystemlldpNetworkPolicyStreamingVideo      `json:"streaming-video,omitempty"`
	VideoConferencing   *[]SystemlldpNetworkPolicyVideoConferencing   `json:"video-conferencing,omitempty"`
	VideoSignaling      *[]SystemlldpNetworkPolicyVideoSignaling      `json:"video-signaling,omitempty"`
	Voice               *[]SystemlldpNetworkPolicyVoice               `json:"voice,omitempty"`
	VoiceSignaling      *[]SystemlldpNetworkPolicyVoiceSignaling      `json:"voice-signaling,omitempty"`
}

type SystemlldpNetworkPolicyGuest struct {
	Dscp     *float64 `json:"dscp,omitempty"`
	Priority *float64 `json:"priority,omitempty"`
	Status   *string  `json:"status,omitempty"`
	Tag      *string  `json:"tag,omitempty"`
	Vlan     *float64 `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicyGuestVoiceSignaling struct {
	Dscp     *float64 `json:"dscp,omitempty"`
	Priority *float64 `json:"priority,omitempty"`
	Status   *string  `json:"status,omitempty"`
	Tag      *string  `json:"tag,omitempty"`
	Vlan     *float64 `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicySoftphone struct {
	Dscp     *float64 `json:"dscp,omitempty"`
	Priority *float64 `json:"priority,omitempty"`
	Status   *string  `json:"status,omitempty"`
	Tag      *string  `json:"tag,omitempty"`
	Vlan     *float64 `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicyStreamingVideo struct {
	Dscp     *float64 `json:"dscp,omitempty"`
	Priority *float64 `json:"priority,omitempty"`
	Status   *string  `json:"status,omitempty"`
	Tag      *string  `json:"tag,omitempty"`
	Vlan     *float64 `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicyVideoConferencing struct {
	Dscp     *float64 `json:"dscp,omitempty"`
	Priority *float64 `json:"priority,omitempty"`
	Status   *string  `json:"status,omitempty"`
	Tag      *string  `json:"tag,omitempty"`
	Vlan     *float64 `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicyVideoSignaling struct {
	Dscp     *float64 `json:"dscp,omitempty"`
	Priority *float64 `json:"priority,omitempty"`
	Status   *string  `json:"status,omitempty"`
	Tag      *string  `json:"tag,omitempty"`
	Vlan     *float64 `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicyVoice struct {
	Dscp     *float64 `json:"dscp,omitempty"`
	Priority *float64 `json:"priority,omitempty"`
	Status   *string  `json:"status,omitempty"`
	Tag      *string  `json:"tag,omitempty"`
	Vlan     *float64 `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicyVoiceSignaling struct {
	Dscp     *float64 `json:"dscp,omitempty"`
	Priority *float64 `json:"priority,omitempty"`
	Status   *string  `json:"status,omitempty"`
	Tag      *string  `json:"tag,omitempty"`
	Vlan     *float64 `json:"vlan,omitempty"`
}
