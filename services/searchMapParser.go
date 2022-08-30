package services

import "strings"

func SearchMapParser(data map[string]string) string {
	searchMap := make(map[string]string)
	searchConditions := ""
	if sId, ok := data["Id"]; ok {
		searchMap["id"] = sId
	}

	if sName, ok := data["Name"]; ok {
		searchMap["name"] = sName
	}

	if sCategories, ok := data["Categories"]; ok {
		searchMap["categories"] = sCategories
	}

	if sPrice, ok := data["Price"]; ok {
		searchMap["price"] = sPrice
	}

	if sDiscount, ok := data["Discount"]; ok {
		searchMap["discount"] = sDiscount
	}

	if sDescription, ok := data["Description"]; ok {
		searchMap["description"] = sDescription
	}

	if sImages, ok := data["Images"]; ok {
		searchMap["images"] = sImages
	}

	if len(searchMap) > 0 {
		searchConditions += "WHERE "
		for k, v := range searchMap {
			searchConditions += k + " = '" + v + "' AND "
		}
	}
	searchConditions = strings.TrimSuffix(searchConditions, "AND ")
	return searchConditions
}
