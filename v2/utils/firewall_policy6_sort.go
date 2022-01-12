package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

type utilsFirewallPolicy6 struct {
	Comments string `json:"comments,omitempty"`
	Name     string `json:"name,omitempty"`
	PolicyID int64  `json:"policyid,omitempty"`
}

const utilsFirewallPolicy6Path = "firewall/policy6/"

func (c *Client) utilsGetFirewallPolicy6List(params *models.CmdbRequestParams) (policyList *[]utilsFirewallPolicy6, err error) {
	req := &models.CmdbRequest{}

	format := []string{"policyid", "name", "comments"}

	req.HTTPMethod = "GET"
	req.Payload = nil
	req.Path = models.CmdbBasePath + utilsFirewallPolicy6Path
	req.Params = *params
	req.Params.Format = &format

	res, err := request.Read(c.config, req)
	if err != nil {
		return nil, err
	}

	if tmp, ok := res.Results.([]interface{}); ok {
		jsontmp, err := json.Marshal(tmp)
		if err != nil {
			return nil, err
		}
		v := []utilsFirewallPolicy6{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}

	return nil, err
}

func (c *Client) utilsSortFirewallPolicy6List(policyList []utilsFirewallPolicy6, sortBy, sortDirection string) (err error) {
	log.Println("[DEBUG] sorting firewall policy list")
	if len(policyList) == 0 {
		log.Println("[INFO] list is empty")
		return nil
	}
	if sortDirection == "descending" {
		switch sortBy := strings.ToLower(sortBy); sortBy {
		case "name":
			sort.SliceStable(policyList, func(i, j int) bool {
				v1 := policyList[i].Name
				v2 := policyList[j].Name
				return v1 < v2
			})
		case "policyid":
			sort.SliceStable(policyList, func(i, j int) bool {
				v1 := policyList[i].PolicyID
				v2 := policyList[j].PolicyID
				return v1 < v2
			})
		case "comments":
			sort.SliceStable(policyList, func(i, j int) bool {
				v1 := policyList[i].Comments
				v2 := policyList[j].Comments
				return v1 < v2
			})
		}
		return nil
	} else if sortDirection == "ascending" {
		switch sortBy := strings.ToLower(sortBy); sortBy {
		case "name":
			sort.SliceStable(policyList, func(i, j int) bool {
				v1 := policyList[i].Name
				v2 := policyList[j].Name
				return v1 > v2
			})
		case "policyid":
			sort.SliceStable(policyList, func(i, j int) bool {
				v1 := policyList[i].PolicyID
				v2 := policyList[j].PolicyID
				return v1 > v2
			})
		case "comments":
			sort.SliceStable(policyList, func(i, j int) bool {
				v1 := policyList[i].Comments
				v2 := policyList[j].Comments
				return v1 > v2
			})
		}
		return nil
	}
	return fmt.Errorf("sorting failed")
}

func (c *Client) utilsFirewallPolicy6ListIsSorted(policyList []utilsFirewallPolicy6, sortBy, sortDirection string) (sorted bool, err error) {
	log.Println("[DEBUG] checking sort status")
	if len(policyList) == 0 {
		log.Println("[INFO] list is empty")
		return true, nil
	}
	if sortDirection == "descending" {
		switch sortBy := strings.ToLower(sortBy); sortBy {
		case "name":
			sorted := sort.SliceIsSorted(policyList, func(i, j int) bool {
				v1 := policyList[i].Name
				v2 := policyList[j].Name
				return v1 < v2
			})
			return sorted, nil
		case "policyid":
			return sort.SliceIsSorted(policyList, func(i, j int) bool {
				v1 := policyList[i].PolicyID
				v2 := policyList[j].PolicyID
				return v1 < v2
			}), nil
		case "comments":
			return sort.SliceIsSorted(policyList, func(i, j int) bool {
				v1 := policyList[i].Comments
				v2 := policyList[j].Comments
				return v1 < v2
			}), nil
		}
		return false, fmt.Errorf("unsupported sortBy: %s", sortBy)
	} else if sortDirection == "ascending" {
		switch sortBy := strings.ToLower(sortBy); sortBy {
		case "name":
			sorted := sort.SliceIsSorted(policyList, func(i, j int) bool {
				v1 := policyList[i].Name
				v2 := policyList[j].Name
				return v1 > v2
			})
			return sorted, nil
		case "policyid":
			return sort.SliceIsSorted(policyList, func(i, j int) bool {
				v1 := policyList[i].PolicyID
				v2 := policyList[j].PolicyID
				return v1 > v2
			}), nil
		case "comments":
			return sort.SliceIsSorted(policyList, func(i, j int) bool {
				v1 := policyList[i].Comments
				v2 := policyList[j].Comments
				return v1 > v2
			}), nil
		}
		return false, fmt.Errorf("unsupported sortBy: %s", sortBy)
	}
	return false, fmt.Errorf("sorting check failed")
}

// Move the policyMoved policy AFTER the policyTarget policy
func (c *Client) utilsMoveAfterFirewallPolicy6(policyMoved, policyTarget int) (err error) {
	if policyMoved == policyTarget {
		log.Println("[INFO] policyid matches, ignoring.")
		return nil
	}
	idMoved := strconv.Itoa(policyMoved)
	idTarget := strconv.Itoa(policyTarget)

	req := &models.CmdbRequest{}
	params := &models.CmdbRequestParams{}
	params.Action = "move"
	params.After = idTarget

	req.HTTPMethod = "PUT"
	req.Payload = nil
	req.Path = models.CmdbBasePath + utilsFirewallPolicy6Path + idMoved
	req.Params = *params

	_, err = request.CreateUpdate(c.config, req)
	if err != nil {
		return err
	}
	return nil
}

// Move the policyMoved policy BEFORE the policyTarget policy
func (c *Client) utilsMoveBeforeFirewallPolicy6(policyMoved, policyTarget int) (err error) {
	if policyMoved == policyTarget {
		log.Println("[INFO] policyid matches, ignoring.")
		return nil
	}
	idMoved := strconv.Itoa(policyMoved)
	idTarget := strconv.Itoa(policyTarget)

	req := &models.CmdbRequest{}
	params := &models.CmdbRequestParams{}
	params.Action = "move"
	params.Before = idTarget

	req.HTTPMethod = "PUT"
	req.Payload = nil
	req.Path = models.CmdbBasePath + utilsFirewallPolicy6Path + idMoved
	req.Params = *params

	_, err = request.CreateUpdate(c.config, req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SortFirewallPolicy6(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error) {
	policyList, err := c.utilsGetFirewallPolicy6List(params)
	if err != nil {
		return err
	}

	sorted, err := c.utilsFirewallPolicy6ListIsSorted(*policyList, sortBy, sortDirection)
	if err != nil {
		return err
	}

	if sorted {
		return nil
	}

	currentLast := (*policyList)[len(*policyList)-1].PolicyID
	c.utilsSortFirewallPolicy6List(*policyList, sortBy, sortDirection)

	for i := range *policyList {
		i1 := len(*policyList) - i - 1
		p1 := (*policyList)[i1]
		if i == 0 {
			err := c.utilsMoveAfterFirewallPolicy6(int(p1.PolicyID), int(currentLast))
			if err != nil {
				return err
			}
		} else {
			i2 := len(*policyList) - i
			p2 := (*policyList)[i2]
			err := c.utilsMoveBeforeFirewallPolicy6(int(p1.PolicyID), int(p2.PolicyID))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Client) FirewallPolicy6ListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error) {
	policyList, err := c.utilsGetFirewallPolicy6List(params)
	if err != nil {
		return false, err
	}

	sorted, err = c.utilsFirewallPolicy6ListIsSorted(*policyList, sortBy, sortDirection)
	if err != nil {
		return false, err
	}

	if sorted {
		return true, nil
	}

	return false, nil
}
