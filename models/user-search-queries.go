package models

import (
	"fmt"
	"regexp"
	"strings"
)

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

func UserSearchQueriesParser(sq UserSearchQueries) (searchConditionsPrepare string, searchConditionsValues []string) {

	// destruct ID
	if sq.ID != 0 {
		searchConditionsPrepare += "id = ?"
		searchConditionsValues = append(searchConditionsValues, fmt.Sprint(sq.ID))
		return
	}

	// destruct Username
	if sq.Username != "" {
		searchConditionsPrepare += "username = ?"
		searchConditionsValues = append(searchConditionsValues, sq.Username)
		return
	}

	// destruct Email
	if sq.Email != "" {
		searchConditionsPrepare += "email = ?"
		searchConditionsValues = append(searchConditionsValues, sq.Email)
		return
	}

	// destruct Name
	if sq.Name != "" {
		searchConditionsPrepare += "AND name LIKE ?"
		searchConditionsValues = append(searchConditionsValues, "%"+sq.Name+"%")
	}

	// destruct Phone
	if sq.Phone != "" {
		extractCompareMark(sq.Phone)
		searchConditionsPrepare += "AND phone LIKE ?"
		searchConditionsValues = append(searchConditionsValues, "%"+sq.Phone+"%")
	}

	if strings.Split(searchConditionsPrepare, " ")[0] == "AND" {
		searchConditionsPrepare = searchConditionsPrepare[3:]
	}
	return
}

func extractCompareMark(input string) {
	pattern, compileErr := regexp.Compile("[A-z]ork")
	correctAnswer := "Yes, I love new york city"
	question := "Do you love new york city?"
	wrongAnswer := "I love dogs"
	if compileErr == nil {
		fmt.Println("Question:", pattern.MatchString(question))
		fmt.Println("Correct Answer:", pattern.MatchString(correctAnswer))
		fmt.Println("Wrong Answer:", pattern.MatchString(wrongAnswer))
	} else {
		fmt.Println("Error:", compileErr)
	}
}
