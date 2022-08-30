package main

import (
	"database/sql"
	"log"

	"github.com/ahmedsat/sat-store/controllers"
	"github.com/ahmedsat/sat-store/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

//	type user struct {
//		id        int
//		name      string
//		username  string
//		email     string
//		password  string
//		privilege string
//	}






func main() {

	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	controllers.DB = db

	r := gin.Default()

	routes.GetRoutes(r)

	r.Run("localhost:8080")

}



// func allUsers() ([]user, error) {
// 	var users []user

// 	rows, err := db.Query(`SELECT * FROM users`)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var u user
// 		if err = rows.Scan(&u.id, &u.name, &u.username, &u.email, &u.password, &u.privilege); err != nil {
// 			return nil, err
// 		}
// 		users = append(users, u)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }
