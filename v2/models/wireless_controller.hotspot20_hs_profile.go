package models

const WirelessControllerHotspot20HsProfilePath = "wireless-controller.hotspot20/hs-profile/"

type WirelessControllerHotspot20HsProfile struct {
	The3gppPlmn           *string                                            `json:"3gpp-plmn,omitempty"`
	AccessNetworkAsra     *string                                            `json:"access-network-asra,omitempty"`
	AccessNetworkEsr      *string                                            `json:"access-network-esr,omitempty"`
	AccessNetworkInternet *string                                            `json:"access-network-internet,omitempty"`
	AccessNetworkType     *string                                            `json:"access-network-type,omitempty"`
	AccessNetworkUesa     *string                                            `json:"access-network-uesa,omitempty"`
	AdviceOfCharge        *string                                            `json:"advice-of-charge,omitempty"`
	AnqpDomainId          *int64                                             `json:"anqp-domain-id,omitempty"`
	BssTransition         *string                                            `json:"bss-transition,omitempty"`
	ConnCap               *string                                            `json:"conn-cap,omitempty"`
	DeauthRequestTimeout  *int64                                             `json:"deauth-request-timeout,omitempty"`
	Dgaf                  *string                                            `json:"dgaf,omitempty"`
	DomainName            *string                                            `json:"domain-name,omitempty"`
	GasComebackDelay      *int64                                             `json:"gas-comeback-delay,omitempty"`
	GasFragmentationLimit *int64                                             `json:"gas-fragmentation-limit,omitempty"`
	Hessid                *string                                            `json:"hessid,omitempty"`
	IpAddrType            *string                                            `json:"ip-addr-type,omitempty"`
	L2tif                 *string                                            `json:"l2tif,omitempty"`
	NaiRealm              *string                                            `json:"nai-realm,omitempty"`
	Name                  *string                                            `json:"name,omitempty"`
	NetworkAuth           *string                                            `json:"network-auth,omitempty"`
	OperFriendlyName      *string                                            `json:"oper-friendly-name,omitempty"`
	OperIcon              *string                                            `json:"oper-icon,omitempty"`
	OsuProvider           *[]WirelessControllerHotspot20HsProfileOsuProvider `json:"osu-provider,omitempty"`
	OsuProviderNai        *string                                            `json:"osu-provider-nai,omitempty"`
	OsuSsid               *string                                            `json:"osu-ssid,omitempty"`
	PameBi                *string                                            `json:"pame-bi,omitempty"`
	ProxyArp              *string                                            `json:"proxy-arp,omitempty"`
	QosMap                *string                                            `json:"qos-map,omitempty"`
	Release               *int64                                             `json:"release,omitempty"`
	RoamingConsortium     *string                                            `json:"roaming-consortium,omitempty"`
	TermsAndConditions    *string                                            `json:"terms-and-conditions,omitempty"`
	VenueGroup            *string                                            `json:"venue-group,omitempty"`
	VenueName             *string                                            `json:"venue-name,omitempty"`
	VenueType             *string                                            `json:"venue-type,omitempty"`
	VenueUrl              *string                                            `json:"venue-url,omitempty"`
	WanMetrics            *string                                            `json:"wan-metrics,omitempty"`
	WnmSleepMode          *string                                            `json:"wnm-sleep-mode,omitempty"`
}

type WirelessControllerHotspot20HsProfileOsuProvider struct {
	Name *string `json:"name,omitempty"`
}
