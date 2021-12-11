package models

const SystemlldpNetworkPolicyPath = "system.lldp/network-policy/"

type SystemlldpNetworkPolicy struct {
	Comment             *string                                      `json:"comment,omitempty"`
	Guest               []SystemlldpNetworkPolicyGuest               `json:"guest,omitempty"`
	GuestVoiceSignaling []SystemlldpNetworkPolicyGuestVoiceSignaling `json:"guest-voice-signaling,omitempty"`
	Name                *string                                      `json:"name,omitempty"`
	Softphone           []SystemlldpNetworkPolicySoftphone           `json:"softphone,omitempty"`
	StreamingVideo      []SystemlldpNetworkPolicyStreamingVideo      `json:"streaming-video,omitempty"`
	VideoConferencing   []SystemlldpNetworkPolicyVideoConferencing   `json:"video-conferencing,omitempty"`
	VideoSignaling      []SystemlldpNetworkPolicyVideoSignaling      `json:"video-signaling,omitempty"`
	Voice               []SystemlldpNetworkPolicyVoice               `json:"voice,omitempty"`
	VoiceSignaling      []SystemlldpNetworkPolicyVoiceSignaling      `json:"voice-signaling,omitempty"`
}

type SystemlldpNetworkPolicyGuest struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicyGuestVoiceSignaling struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicySoftphone struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicyStreamingVideo struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicyVideoConferencing struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicyVideoSignaling struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicyVoice struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

type SystemlldpNetworkPolicyVoiceSignaling struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}
