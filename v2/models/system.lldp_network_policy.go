package models

const SystemLldpNetworkPolicyPath = "system.lldp/network-policy/"

type SystemLldpNetworkPolicy struct {
	Comment             *string                                     `json:"comment,omitempty"`
	Guest               *SystemLldpNetworkPolicyGuest               `json:"guest,omitempty"`
	GuestVoiceSignaling *SystemLldpNetworkPolicyGuestVoiceSignaling `json:"guest-voice-signaling,omitempty"`
	Name                *string                                     `json:"name,omitempty"`
	Softphone           *SystemLldpNetworkPolicySoftphone           `json:"softphone,omitempty"`
	StreamingVideo      *SystemLldpNetworkPolicyStreamingVideo      `json:"streaming-video,omitempty"`
	VideoConferencing   *SystemLldpNetworkPolicyVideoConferencing   `json:"video-conferencing,omitempty"`
	VideoSignaling      *SystemLldpNetworkPolicyVideoSignaling      `json:"video-signaling,omitempty"`
	Voice               *SystemLldpNetworkPolicyVoice               `json:"voice,omitempty"`
	VoiceSignaling      *SystemLldpNetworkPolicyVoiceSignaling      `json:"voice-signaling,omitempty"`
}

const SystemLldpNetworkPolicyGuestPath = "system.lldp/network-policy/guest/"

type SystemLldpNetworkPolicyGuest struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

const SystemLldpNetworkPolicyGuestVoiceSignalingPath = "system.lldp/network-policy/guest-voice-signaling/"

type SystemLldpNetworkPolicyGuestVoiceSignaling struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

const SystemLldpNetworkPolicySoftphonePath = "system.lldp/network-policy/softphone/"

type SystemLldpNetworkPolicySoftphone struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

const SystemLldpNetworkPolicyStreamingVideoPath = "system.lldp/network-policy/streaming-video/"

type SystemLldpNetworkPolicyStreamingVideo struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

const SystemLldpNetworkPolicyVideoConferencingPath = "system.lldp/network-policy/video-conferencing/"

type SystemLldpNetworkPolicyVideoConferencing struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

const SystemLldpNetworkPolicyVideoSignalingPath = "system.lldp/network-policy/video-signaling/"

type SystemLldpNetworkPolicyVideoSignaling struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

const SystemLldpNetworkPolicyVoicePath = "system.lldp/network-policy/voice/"

type SystemLldpNetworkPolicyVoice struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}

const SystemLldpNetworkPolicyVoiceSignalingPath = "system.lldp/network-policy/voice-signaling/"

type SystemLldpNetworkPolicyVoiceSignaling struct {
	Dscp     *int64  `json:"dscp,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
	Tag      *string `json:"tag,omitempty"`
	Vlan     *int64  `json:"vlan,omitempty"`
}
