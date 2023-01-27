package DbConnect

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	connectionString := "user=postgres dbname=product_catalog password=akash host=localhost sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("hello") //check here
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
}

func ConnectToDB() *sql.DB {
	return db
}
