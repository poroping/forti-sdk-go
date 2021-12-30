package models

const WanoptContentDeliveryNetworkRulePath = "wanopt/content-delivery-network-rule/"

type WanoptContentDeliveryNetworkRule struct {
	Category             *string                                                 `json:"category,omitempty"`
	Comment              *string                                                 `json:"comment,omitempty"`
	HostDomainNameSuffix *[]WanoptContentDeliveryNetworkRuleHostDomainNameSuffix `json:"host-domain-name-suffix,omitempty"`
	Name                 *string                                                 `json:"name,omitempty"`
	RequestCacheControl  *string                                                 `json:"request-cache-control,omitempty"`
	ResponseCacheControl *string                                                 `json:"response-cache-control,omitempty"`
	ResponseExpires      *string                                                 `json:"response-expires,omitempty"`
	Rules                *[]WanoptContentDeliveryNetworkRuleRules                `json:"rules,omitempty"`
	Status               *string                                                 `json:"status,omitempty"`
	Updateserver         *string                                                 `json:"updateserver,omitempty"`
}

type WanoptContentDeliveryNetworkRuleHostDomainNameSuffix struct {
	Name *string `json:"name,omitempty"`
}

type WanoptContentDeliveryNetworkRuleRules struct {
	ContentId    *[]WanoptContentDeliveryNetworkRuleRulesContentId    `json:"content-id,omitempty"`
	MatchEntries *[]WanoptContentDeliveryNetworkRuleRulesMatchEntries `json:"match-entries,omitempty"`
	MatchMode    *string                                              `json:"match-mode,omitempty"`
	Name         *string                                              `json:"name,omitempty"`
	SkipEntries  *[]WanoptContentDeliveryNetworkRuleRulesSkipEntries  `json:"skip-entries,omitempty"`
	SkipRuleMode *string                                              `json:"skip-rule-mode,omitempty"`
}

type WanoptContentDeliveryNetworkRuleRulesContentId struct {
	EndDirection   *string `json:"end-direction,omitempty"`
	EndSkip        *int64  `json:"end-skip,omitempty"`
	EndStr         *string `json:"end-str,omitempty"`
	RangeStr       *string `json:"range-str,omitempty"`
	StartDirection *string `json:"start-direction,omitempty"`
	StartSkip      *int64  `json:"start-skip,omitempty"`
	StartStr       *string `json:"start-str,omitempty"`
	Target         *string `json:"target,omitempty"`
}

type WanoptContentDeliveryNetworkRuleRulesMatchEntries struct {
	Id      *int64                                                      `json:"id,omitempty"`
	Pattern *[]WanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern `json:"pattern,omitempty"`
	Target  *string                                                     `json:"target,omitempty"`
}

type WanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern struct {
	String *string `json:"string,omitempty"`
}

type WanoptContentDeliveryNetworkRuleRulesSkipEntries struct {
	Id      *int64                                                     `json:"id,omitempty"`
	Pattern *[]WanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern `json:"pattern,omitempty"`
	Target  *string                                                    `json:"target,omitempty"`
}

type WanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern struct {
	String *string `json:"string,omitempty"`
}
