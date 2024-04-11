package conn

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnFunc() *sql.DB {
	// host := "localhost"
	// port := 5432
	// user := "postgres"
	// password := "12345"
	// dbname := "mobileapp"
	// connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	db, err := sql.Open("postgres", "postgres://yonyswmk:WlrJTTKdk43TpeyXgYQZRVnBpVA-YQaM@bubble.db.elephantsql.com/yonyswmk")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the PostgreSQL database")
	return db
}
