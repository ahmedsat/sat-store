package models

import (
	"fmt"
)

type searchConditions struct {
	Prepare string
	Values  []string
}

//	TODO :
//		[ ] search by CreatedAt
//		[ ] search by UpdatedAt
//		[ ] search by DeletedAt

type UserSearchQueries struct {
	ID         uint   `form:"id"`
	Username   string `form:"username" `
	Email      string `form:"email" `
	Password   string `form:"password"`
	Name       string `form:"name"`
	Phone      string `form:"Phone"`
	Address    string `form:"address"`
	Privileges string `form:"privileges"`

	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}

var prefixList = map[string]bool{
	"=": true,
	">": true,
	"<": true,
	"!": true,
	"~": true,
}

// to see what happen if return a string instated of pointer of string
// see this : "https://go.dev/play/p/cf_Cg_vtXpP"
func UserSearchQueriesParser(sq UserSearchQueries) (*string, []string) {

	searchConditionsObject := searchConditions{}

	defer func(searchConditions *searchConditions) {

		// if searchConditionsPrepare not empty
		// remove " And" from the beginning
		if searchConditions.Prepare != "" {
			searchConditions.Prepare = searchConditions.Prepare[4:]
		}
		// move the value of searchConditionsPrepare to new function scoped variable,
		// // and reset it's value .
		// // ! this step is necessary because global variables will be remembered to next request
		// searchPrepare = searchConditionsPrepare
		// searchConditionsPrepare = ""

		// // we have to do in searchConditionsValues the same as searchConditionsPrepare
		// searchValues = searchConditionsValues
		// searchConditionsValues = []string{}
	}(&searchConditionsObject)

	// destruct ID
	if sq.ID != 0 {
		destructSearchField("="+fmt.Sprint(sq.ID), "id", &searchConditionsObject)
		return &searchConditionsObject.Prepare, searchConditionsObject.Values
	}

	// destruct Username
	if sq.Username != "" {

		destructSearchField("="+sq.Username, "username", &searchConditionsObject)
		return &searchConditionsObject.Prepare, searchConditionsObject.Values

	}

	// destruct Email
	if sq.Email != "" {

		destructSearchField("="+sq.Email, "email", &searchConditionsObject)
		return &searchConditionsObject.Prepare, searchConditionsObject.Values

	}

	// destruct Phone
	if sq.Phone != "" {

		destructSearchField(sq.Phone, "phone", &searchConditionsObject)

	}

	// destruct Name
	if sq.Name != "" {

		destructSearchField(sq.Name, "name", &searchConditionsObject)

	}

	// destruct Address
	if sq.Address != "" {

		destructSearchField(sq.Address, "address", &searchConditionsObject)

	}

	// destruct Privileges
	if sq.Privileges != "" {

		destructSearchField("="+sq.Privileges, "privileges", &searchConditionsObject)

	}
	return &searchConditionsObject.Prepare, searchConditionsObject.Values
}

func extractCompareMark(input string) (compareMark, searchKey string) {

	if !prefixList[string(input[0])] {
		input = "=" + input
	}
	compareMark, searchKey = getPrefix("", input)
	return
}

func getPrefix(prefix, input string) (newPrefix, newInput string) {
	newPrefix = prefix
	newInput = input

	if prefixList[string(input[0])] {
		newPrefix = prefix + string(input[0])
		newInput = input[1:]
		return getPrefix(newPrefix, newInput)

	}

	return
}

func destructSearchField(search, tableName string, searchConditions *searchConditions) {
	compareMark, searchKey := extractCompareMark(search) // parsing requested search

	// map available Compare Marks in search to SQL compare marks
	availableCompareMarks := map[string]string{
		">":  ">",
		"<":  "<",
		"=":  "=",
		"==": "=",
		"!":  "<>",
		"!=": "<>",
		"<>": "<>",
		"<=": "<=",
		">=": ">=",
		"~":  "LIKE",
		"~=": "LIKE",
	}

	// if requested search prefix mach with map do sql request
	if _, ok := availableCompareMarks[compareMark]; ok {
		searchConditions.Prepare += fmt.Sprintf(" AND %s %s ?", tableName, availableCompareMarks[compareMark])
		// if search request asks for same equal add wild card
		if availableCompareMarks[compareMark] == "LIKE" {
			searchConditions.Values = append(searchConditions.Values, "%"+searchKey+"%")
		} else {
			searchConditions.Values = append(searchConditions.Values, searchKey)
		}
	}

}
