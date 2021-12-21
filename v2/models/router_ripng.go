package models

const RouterRipngPath = "router/ripng/"

type RouterRipng struct {
	AggregateAddress            *[]RouterRipngAggregateAddress `json:"aggregate-address,omitempty"`
	DefaultInformationOriginate *string                        `json:"default-information-originate,omitempty"`
	DefaultMetric               *float64                       `json:"default-metric,omitempty"`
	Distance                    *[]RouterRipngDistance         `json:"distance,omitempty"`
	DistributeList              *[]RouterRipngDistributeList   `json:"distribute-list,omitempty"`
	GarbageTimer                *float64                       `json:"garbage-timer,omitempty"`
	Interface                   *[]RouterRipngInterface        `json:"interface,omitempty"`
	MaxOutMetric                *float64                       `json:"max-out-metric,omitempty"`
	Neighbor                    *[]RouterRipngNeighbor         `json:"neighbor,omitempty"`
	Network                     *[]RouterRipngNetwork          `json:"network,omitempty"`
	OffsetList                  *[]RouterRipngOffsetList       `json:"offset-list,omitempty"`
	PassiveInterface            *[]RouterRipngPassiveInterface `json:"passive-interface,omitempty"`
	Redistribute                *[]RouterRipngRedistribute     `json:"redistribute,omitempty"`
	TimeoutTimer                *float64                       `json:"timeout-timer,omitempty"`
	UpdateTimer                 *float64                       `json:"update-timer,omitempty"`
}

type RouterRipngAggregateAddress struct {
	Id      *float64 `json:"id,omitempty"`
	Prefix6 *string  `json:"prefix6,omitempty"`
}

type RouterRipngDistance struct {
	AccessList6 *string  `json:"access-list6,omitempty"`
	Distance    *float64 `json:"distance,omitempty"`
	Id          *float64 `json:"id,omitempty"`
	Prefix6     *string  `json:"prefix6,omitempty"`
}

type RouterRipngDistributeList struct {
	Direction *string  `json:"direction,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	Interface *string  `json:"interface,omitempty"`
	Listname  *string  `json:"listname,omitempty"`
	Status    *string  `json:"status,omitempty"`
}

type RouterRipngInterface struct {
	Flags              *float64 `json:"flags,omitempty"`
	Name               *string  `json:"name,omitempty"`
	SplitHorizon       *string  `json:"split-horizon,omitempty"`
	SplitHorizonStatus *string  `json:"split-horizon-status,omitempty"`
}

type RouterRipngNeighbor struct {
	Id        *float64 `json:"id,omitempty"`
	Interface *string  `json:"interface,omitempty"`
	Ip6       *string  `json:"ip6,omitempty"`
}

type RouterRipngNetwork struct {
	Id     *float64 `json:"id,omitempty"`
	Prefix *string  `json:"prefix,omitempty"`
}

type RouterRipngOffsetList struct {
	AccessList6 *string  `json:"access-list6,omitempty"`
	Direction   *string  `json:"direction,omitempty"`
	Id          *float64 `json:"id,omitempty"`
	Interface   *string  `json:"interface,omitempty"`
	Offset      *float64 `json:"offset,omitempty"`
	Status      *string  `json:"status,omitempty"`
}

type RouterRipngPassiveInterface struct {
	Name *string `json:"name,omitempty"`
}

type RouterRipngRedistribute struct {
	Metric   *float64 `json:"metric,omitempty"`
	Name     *string  `json:"name,omitempty"`
	Routemap *string  `json:"routemap,omitempty"`
	Status   *string  `json:"status,omitempty"`
}
