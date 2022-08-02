package models

const AntivirusQuarantinePath = "antivirus/quarantine/"

type AntivirusQuarantine struct {
	Agelimit             *int64  `json:"agelimit,omitempty"`
	Destination          *string `json:"destination,omitempty"`
	DropBlocked          *string `json:"drop-blocked,omitempty"`
	DropHeuristic        *string `json:"drop-heuristic,omitempty"`
	DropInfected         *string `json:"drop-infected,omitempty"`
	DropMachineLearning  *string `json:"drop-machine-learning,omitempty"`
	Lowspace             *string `json:"lowspace,omitempty"`
	Maxfilesize          *int64  `json:"maxfilesize,omitempty"`
	QuarantineQuota      *int64  `json:"quarantine-quota,omitempty"`
	StoreBlocked         *string `json:"store-blocked,omitempty"`
	StoreHeuristic       *string `json:"store-heuristic,omitempty"`
	StoreInfected        *string `json:"store-infected,omitempty"`
	StoreMachineLearning *string `json:"store-machine-learning,omitempty"`
}
