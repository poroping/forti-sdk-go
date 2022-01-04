package models

const FirewallServiceCategoryPath = "firewall.service/category/"

type FirewallServiceCategory struct {
	Comment      *string `json:"comment,omitempty"`
	FabricObject *string `json:"fabric-object,omitempty"`
	Name         *string `json:"name,omitempty"`
}
