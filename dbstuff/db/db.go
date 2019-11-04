package db

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DbConnect() database {
	db, err := sql.Open("mysql", "root:2HousesOn1PathFollowing9*!@/localhost:3306")

	if err != nil {
		fmt.Println("Problem with connection")
	} else {
		fmt.Println("We good.")
	}

	return db
}