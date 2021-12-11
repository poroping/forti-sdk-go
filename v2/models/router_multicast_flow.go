package models

const RouterMulticastFlowPath = "router/multicast-flow/"

type RouterMulticastFlow struct {
	Comments *string                    `json:"comments,omitempty"`
	Flows    []RouterMulticastFlowFlows `json:"flows,omitempty"`
	Name     *string                    `json:"name,omitempty"`
}

type RouterMulticastFlowFlows struct {
	GroupAddr  *string `json:"group-addr,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	SourceAddr *string `json:"source-addr,omitempty"`
}
