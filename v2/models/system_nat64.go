package models

const SystemNat64Path = "system/nat64/"

type SystemNat64 struct {
	AlwaysSynthesizeAaaaRecord     *string                       `json:"always-synthesize-aaaa-record,omitempty"`
	GenerateIpv6FragmentHeader     *string                       `json:"generate-ipv6-fragment-header,omitempty"`
	Nat46ForceIpv4PacketForwarding *string                       `json:"nat46-force-ipv4-packet-forwarding,omitempty"`
	Nat64Prefix                    *string                       `json:"nat64-prefix,omitempty"`
	SecondaryPrefix                *[]SystemNat64SecondaryPrefix `json:"secondary-prefix,omitempty"`
	SecondaryPrefixStatus          *string                       `json:"secondary-prefix-status,omitempty"`
	Status                         *string                       `json:"status,omitempty"`
}

const SystemNat64SecondaryPrefixPath = "system/nat64/secondary-prefix/"

type SystemNat64SecondaryPrefix struct {
	Name        *string `json:"name,omitempty"`
	Nat64Prefix *string `json:"nat64-prefix,omitempty"`
}
