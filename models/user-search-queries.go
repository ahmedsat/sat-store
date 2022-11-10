package models

import (
	"fmt"
)

var searchConditionsPrepare string
var searchConditionsValues []string

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

func UserSearchQueriesParser(sq UserSearchQueries) (searchPrepare string, searchValues []string) {

	defer func(SearchConditionsPrepare *string, SearchConditionsValues []string) {

		// if searchConditionsPrepare not empty
		// remove " And" from the beginning
		if searchConditionsPrepare != "" {
			searchConditionsPrepare = searchConditionsPrepare[4:]
		}
		// move the value of searchConditionsPrepare to new function scoped variable,
		// and reset it's value .
		// ! this step is necessary because global variables will be remembered to next request
		searchPrepare = searchConditionsPrepare
		searchConditionsPrepare = ""

		// we have to do in searchConditionsValues the same as searchConditionsPrepare
		searchValues = searchConditionsValues
		searchConditionsValues = []string{}
	}(&searchPrepare, searchValues)

	// destruct ID
	if sq.ID != 0 {
		destructSearchField("="+fmt.Sprint(sq.ID), "id")
		return
	}

	// destruct Username
	if sq.Username != "" {

		destructSearchField("="+sq.Username, "username")
		return

	}

	// destruct Email
	if sq.Email != "" {

		destructSearchField("="+sq.Email, "email")
		return

	}

	// destruct Phone
	if sq.Phone != "" {

		destructSearchField(sq.Phone, "phone")

	}

	// destruct Name
	if sq.Name != "" {

		destructSearchField(sq.Name, "name")

	}

	// destruct Address
	if sq.Address != "" {

		destructSearchField(sq.Address, "address")

	}

	// destruct Privileges
	if sq.Privileges != "" {

		destructSearchField("="+sq.Privileges, "privileges")

	}
	return
}

func extractCompareMark(input string) (compareMark, searchKey string) {

	compareMark, searchKey = getPrefix("", input)
	return
}

func getPrefix(prefix, input string) (newPrefix, newInput string) {
	newPrefix = prefix
	newInput = input
	prefixList := map[string]bool{
		"=": true,
		">": true,
		"<": true,
		"!": true,
		"~": true,
	}

	if prefixList[string(input[0])] {
		newPrefix = prefix + string(input[0])
		newInput = input[1:]
		return getPrefix(newPrefix, newInput)

	}

	return
}

func destructSearchField(search, tableName string) {
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
		searchConditionsPrepare += fmt.Sprintf(" AND %s %s ?", tableName, availableCompareMarks[compareMark])
		// if search request asks for same equal add wild card
		if availableCompareMarks[compareMark] == "LIKE" {
			searchConditionsValues = append(searchConditionsValues, "%"+searchKey+"%")
		} else {
			searchConditionsValues = append(searchConditionsValues, searchKey)
		}
	}

}
