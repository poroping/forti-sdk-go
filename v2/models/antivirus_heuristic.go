package models

const AntivirusHeuristicPath = "antivirus/heuristic/"

type AntivirusHeuristic struct {
	Mode *string `json:"mode,omitempty"`
}
