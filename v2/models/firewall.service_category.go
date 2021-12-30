package models

const FirewallserviceCategoryPath = "firewall/service/category/"

type FirewallserviceCategory struct {
	Comment      *string `json:"comment,omitempty"`
	FabricObject *string `json:"fabric-object,omitempty"`
	Name         *string `json:"name,omitempty"`
}
