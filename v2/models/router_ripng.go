package models

const RouterRipngPath = "router/ripng/"

type RouterRipng struct {
	AggregateAddress            *[]RouterRipngAggregateAddress `json:"aggregate-address,omitempty"`
	DefaultInformationOriginate *string                        `json:"default-information-originate,omitempty"`
	DefaultMetric               *int64                         `json:"default-metric,omitempty"`
	Distance                    *[]RouterRipngDistance         `json:"distance,omitempty"`
	DistributeList              *[]RouterRipngDistributeList   `json:"distribute-list,omitempty"`
	GarbageTimer                *int64                         `json:"garbage-timer,omitempty"`
	Interface                   *[]RouterRipngInterface        `json:"interface,omitempty"`
	MaxOutMetric                *int64                         `json:"max-out-metric,omitempty"`
	Neighbor                    *[]RouterRipngNeighbor         `json:"neighbor,omitempty"`
	Network                     *[]RouterRipngNetwork          `json:"network,omitempty"`
	OffsetList                  *[]RouterRipngOffsetList       `json:"offset-list,omitempty"`
	PassiveInterface            *[]RouterRipngPassiveInterface `json:"passive-interface,omitempty"`
	Redistribute                *[]RouterRipngRedistribute     `json:"redistribute,omitempty"`
	TimeoutTimer                *int64                         `json:"timeout-timer,omitempty"`
	UpdateTimer                 *int64                         `json:"update-timer,omitempty"`
}

type RouterRipngAggregateAddress struct {
	Id      *int64  `json:"id,omitempty"`
	Prefix6 *string `json:"prefix6,omitempty"`
}

type RouterRipngDistance struct {
	AccessList6 *string `json:"access-list6,omitempty"`
	Distance    *int64  `json:"distance,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	Prefix6     *string `json:"prefix6,omitempty"`
}

type RouterRipngDistributeList struct {
	Direction *string `json:"direction,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	Interface *string `json:"interface,omitempty"`
	Listname  *string `json:"listname,omitempty"`
	Status    *string `json:"status,omitempty"`
}

type RouterRipngInterface struct {
	Flags              *int64  `json:"flags,omitempty"`
	Name               *string `json:"name,omitempty"`
	SplitHorizon       *string `json:"split-horizon,omitempty"`
	SplitHorizonStatus *string `json:"split-horizon-status,omitempty"`
}

type RouterRipngNeighbor struct {
	Id        *int64  `json:"id,omitempty"`
	Interface *string `json:"interface,omitempty"`
	Ip6       *string `json:"ip6,omitempty"`
}

type RouterRipngNetwork struct {
	Id     *int64  `json:"id,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
}

type RouterRipngOffsetList struct {
	AccessList6 *string `json:"access-list6,omitempty"`
	Direction   *string `json:"direction,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	Interface   *string `json:"interface,omitempty"`
	Offset      *int64  `json:"offset,omitempty"`
	Status      *string `json:"status,omitempty"`
}

type RouterRipngPassiveInterface struct {
	Name *string `json:"name,omitempty"`
}

type RouterRipngRedistribute struct {
	Metric   *int64  `json:"metric,omitempty"`
	Name     *string `json:"name,omitempty"`
	Routemap *string `json:"routemap,omitempty"`
	Status   *string `json:"status,omitempty"`
}

//defaultfuncs
func (def *RouterRipngRedistribute) defaults() {
	def.Metric = "0"
	def.Name = ""
	def.Routemap = ""
	def.Status = "disable"
}
