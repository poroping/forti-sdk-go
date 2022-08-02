package models

const RouterIsisPath = "router/isis/"

type RouterIsis struct {
	AdjacencyCheck       *string                      `json:"adjacency-check,omitempty"`
	AdjacencyCheck6      *string                      `json:"adjacency-check6,omitempty"`
	AdvPassiveOnly       *string                      `json:"adv-passive-only,omitempty"`
	AdvPassiveOnly6      *string                      `json:"adv-passive-only6,omitempty"`
	AuthKeychainL1       *string                      `json:"auth-keychain-l1,omitempty"`
	AuthKeychainL2       *string                      `json:"auth-keychain-l2,omitempty"`
	AuthModeL1           *string                      `json:"auth-mode-l1,omitempty"`
	AuthModeL2           *string                      `json:"auth-mode-l2,omitempty"`
	AuthPasswordL1       *string                      `json:"auth-password-l1,omitempty"`
	AuthPasswordL2       *string                      `json:"auth-password-l2,omitempty"`
	AuthSendonlyL1       *string                      `json:"auth-sendonly-l1,omitempty"`
	AuthSendonlyL2       *string                      `json:"auth-sendonly-l2,omitempty"`
	DefaultOriginate     *string                      `json:"default-originate,omitempty"`
	DefaultOriginate6    *string                      `json:"default-originate6,omitempty"`
	DynamicHostname      *string                      `json:"dynamic-hostname,omitempty"`
	IgnoreLspErrors      *string                      `json:"ignore-lsp-errors,omitempty"`
	IsType               *string                      `json:"is-type,omitempty"`
	IsisInterface        *[]RouterIsisIsisInterface   `json:"isis-interface,omitempty"`
	IsisNet              *[]RouterIsisIsisNet         `json:"isis-net,omitempty"`
	LspGenIntervalL1     *int64                       `json:"lsp-gen-interval-l1,omitempty"`
	LspGenIntervalL2     *int64                       `json:"lsp-gen-interval-l2,omitempty"`
	LspRefreshInterval   *int64                       `json:"lsp-refresh-interval,omitempty"`
	MaxLspLifetime       *int64                       `json:"max-lsp-lifetime,omitempty"`
	MetricStyle          *string                      `json:"metric-style,omitempty"`
	OverloadBit          *string                      `json:"overload-bit,omitempty"`
	OverloadBitOnStartup *int64                       `json:"overload-bit-on-startup,omitempty"`
	OverloadBitSuppress  *string                      `json:"overload-bit-suppress,omitempty"`
	Redistribute         *[]RouterIsisRedistribute    `json:"redistribute,omitempty"`
	RedistributeL1       *string                      `json:"redistribute-l1,omitempty"`
	RedistributeL1List   *string                      `json:"redistribute-l1-list,omitempty"`
	RedistributeL2       *string                      `json:"redistribute-l2,omitempty"`
	RedistributeL2List   *string                      `json:"redistribute-l2-list,omitempty"`
	Redistribute6        *[]RouterIsisRedistribute6   `json:"redistribute6,omitempty"`
	Redistribute6L1      *string                      `json:"redistribute6-l1,omitempty"`
	Redistribute6L1List  *string                      `json:"redistribute6-l1-list,omitempty"`
	Redistribute6L2      *string                      `json:"redistribute6-l2,omitempty"`
	Redistribute6L2List  *string                      `json:"redistribute6-l2-list,omitempty"`
	SpfIntervalExpL1     *string                      `json:"spf-interval-exp-l1,omitempty"`
	SpfIntervalExpL2     *string                      `json:"spf-interval-exp-l2,omitempty"`
	SummaryAddress       *[]RouterIsisSummaryAddress  `json:"summary-address,omitempty"`
	SummaryAddress6      *[]RouterIsisSummaryAddress6 `json:"summary-address6,omitempty"`
}

const RouterIsisIsisInterfacePath = "router/isis/isis-interface/"

type RouterIsisIsisInterface struct {
	AuthKeychainL1        *string `json:"auth-keychain-l1,omitempty"`
	AuthKeychainL2        *string `json:"auth-keychain-l2,omitempty"`
	AuthModeL1            *string `json:"auth-mode-l1,omitempty"`
	AuthModeL2            *string `json:"auth-mode-l2,omitempty"`
	AuthPasswordL1        *string `json:"auth-password-l1,omitempty"`
	AuthPasswordL2        *string `json:"auth-password-l2,omitempty"`
	AuthSendOnlyL1        *string `json:"auth-send-only-l1,omitempty"`
	AuthSendOnlyL2        *string `json:"auth-send-only-l2,omitempty"`
	CircuitType           *string `json:"circuit-type,omitempty"`
	CsnpIntervalL1        *int64  `json:"csnp-interval-l1,omitempty"`
	CsnpIntervalL2        *int64  `json:"csnp-interval-l2,omitempty"`
	HelloIntervalL1       *int64  `json:"hello-interval-l1,omitempty"`
	HelloIntervalL2       *int64  `json:"hello-interval-l2,omitempty"`
	HelloMultiplierL1     *int64  `json:"hello-multiplier-l1,omitempty"`
	HelloMultiplierL2     *int64  `json:"hello-multiplier-l2,omitempty"`
	HelloPadding          *string `json:"hello-padding,omitempty"`
	LspInterval           *int64  `json:"lsp-interval,omitempty"`
	LspRetransmitInterval *int64  `json:"lsp-retransmit-interval,omitempty"`
	MeshGroup             *string `json:"mesh-group,omitempty"`
	MeshGroupId           *int64  `json:"mesh-group-id,omitempty"`
	MetricL1              *int64  `json:"metric-l1,omitempty"`
	MetricL2              *int64  `json:"metric-l2,omitempty"`
	Name                  *string `json:"name,omitempty"`
	NetworkType           *string `json:"network-type,omitempty"`
	PriorityL1            *int64  `json:"priority-l1,omitempty"`
	PriorityL2            *int64  `json:"priority-l2,omitempty"`
	Status                *string `json:"status,omitempty"`
	Status6               *string `json:"status6,omitempty"`
	WideMetricL1          *int64  `json:"wide-metric-l1,omitempty"`
	WideMetricL2          *int64  `json:"wide-metric-l2,omitempty"`
}

const RouterIsisIsisNetPath = "router/isis/isis-net/"

type RouterIsisIsisNet struct {
	Id  *int64  `json:"id,omitempty"`
	Net *string `json:"net,omitempty"`
}

const RouterIsisRedistributePath = "router/isis/redistribute/"

type RouterIsisRedistribute struct {
	Level      *string `json:"level,omitempty"`
	Metric     *int64  `json:"metric,omitempty"`
	MetricType *string `json:"metric-type,omitempty"`
	Protocol   *string `json:"protocol,omitempty"`
	Routemap   *string `json:"routemap,omitempty"`
	Status     *string `json:"status,omitempty"`
}

// Set RouterIsisRedistribute values to defaults
func (def *RouterIsisRedistribute) Defaults() {
	def.Level = stringPtr("level-2")
	def.Metric = intPtr(0)
	def.MetricType = stringPtr("internal")
	def.Protocol = stringPtr("")
	def.Routemap = stringPtr("")
	def.Status = stringPtr("disable")
}

const RouterIsisRedistribute6Path = "router/isis/redistribute6/"

type RouterIsisRedistribute6 struct {
	Level      *string `json:"level,omitempty"`
	Metric     *int64  `json:"metric,omitempty"`
	MetricType *string `json:"metric-type,omitempty"`
	Protocol   *string `json:"protocol,omitempty"`
	Routemap   *string `json:"routemap,omitempty"`
	Status     *string `json:"status,omitempty"`
}

// Set RouterIsisRedistribute6 values to defaults
func (def *RouterIsisRedistribute6) Defaults() {
	def.Level = stringPtr("level-2")
	def.Metric = intPtr(0)
	def.MetricType = stringPtr("internal")
	def.Protocol = stringPtr("")
	def.Routemap = stringPtr("")
	def.Status = stringPtr("disable")
}

const RouterIsisSummaryAddressPath = "router/isis/summary-address/"

type RouterIsisSummaryAddress struct {
	Id     *int64  `json:"id,omitempty"`
	Level  *string `json:"level,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
}

const RouterIsisSummaryAddress6Path = "router/isis/summary-address6/"

type RouterIsisSummaryAddress6 struct {
	Id      *int64  `json:"id,omitempty"`
	Level   *string `json:"level,omitempty"`
	Prefix6 *string `json:"prefix6,omitempty"`
}
