package models

const RouterRipPath = "router/rip/"

type RouterRip struct {
	DefaultInformationOriginate *string                      `json:"default-information-originate,omitempty"`
	DefaultMetric               *int64                       `json:"default-metric,omitempty"`
	Distance                    *[]RouterRipDistance         `json:"distance,omitempty"`
	DistributeList              *[]RouterRipDistributeList   `json:"distribute-list,omitempty"`
	GarbageTimer                *int64                       `json:"garbage-timer,omitempty"`
	Interface                   *[]RouterRipInterface        `json:"interface,omitempty"`
	MaxOutMetric                *int64                       `json:"max-out-metric,omitempty"`
	Neighbor                    *[]RouterRipNeighbor         `json:"neighbor,omitempty"`
	Network                     *[]RouterRipNetwork          `json:"network,omitempty"`
	OffsetList                  *[]RouterRipOffsetList       `json:"offset-list,omitempty"`
	PassiveInterface            *[]RouterRipPassiveInterface `json:"passive-interface,omitempty"`
	RecvBufferSize              *int64                       `json:"recv-buffer-size,omitempty"`
	Redistribute                *[]RouterRipRedistribute     `json:"redistribute,omitempty"`
	TimeoutTimer                *int64                       `json:"timeout-timer,omitempty"`
	UpdateTimer                 *int64                       `json:"update-timer,omitempty"`
	Version                     *string                      `json:"version,omitempty"`
}

type RouterRipDistance struct {
	AccessList *string `json:"access-list,omitempty"`
	Distance   *int64  `json:"distance,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	Prefix     *string `json:"prefix,omitempty"`
}

type RouterRipDistributeList struct {
	Direction *string `json:"direction,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	Interface *string `json:"interface,omitempty"`
	Listname  *string `json:"listname,omitempty"`
	Status    *string `json:"status,omitempty"`
}

type RouterRipInterface struct {
	AuthKeychain          *string `json:"auth-keychain,omitempty"`
	AuthMode              *string `json:"auth-mode,omitempty"`
	AuthString            *string `json:"auth-string,omitempty"`
	Flags                 *int64  `json:"flags,omitempty"`
	Name                  *string `json:"name,omitempty"`
	ReceiveVersion        *string `json:"receive-version,omitempty"`
	SendVersion           *string `json:"send-version,omitempty"`
	SendVersion2Broadcast *string `json:"send-version2-broadcast,omitempty"`
	SplitHorizon          *string `json:"split-horizon,omitempty"`
	SplitHorizonStatus    *string `json:"split-horizon-status,omitempty"`
}

type RouterRipNeighbor struct {
	Id *int64  `json:"id,omitempty"`
	Ip *string `json:"ip,omitempty"`
}

type RouterRipNetwork struct {
	Id     *int64  `json:"id,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
}

type RouterRipOffsetList struct {
	AccessList *string `json:"access-list,omitempty"`
	Direction  *string `json:"direction,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	Interface  *string `json:"interface,omitempty"`
	Offset     *int64  `json:"offset,omitempty"`
	Status     *string `json:"status,omitempty"`
}

type RouterRipPassiveInterface struct {
	Name *string `json:"name,omitempty"`
}

type RouterRipRedistribute struct {
	Metric   *int64  `json:"metric,omitempty"`
	Name     *string `json:"name,omitempty"`
	Routemap *string `json:"routemap,omitempty"`
	Status   *string `json:"status,omitempty"`
}

//defaultfuncs
func (def *RouterRipRedistribute) defaults() {
	def.Metric = "0"
	def.Name = ""
	def.Routemap = ""
	def.Status = "disable"
}
