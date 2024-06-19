package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

/*
|--------------------------------------------------------------------------
| D. Database: Postgres Connection
|--------------------------------------------------------------------------
|
| Initialize the Postgres database connection. This will be used to
| interact with the database and perform queries to retrieve or store
| data in the database.
|
*/

func SetupPostgres() {
	db_user := os.Getenv(("DB_USER"))
	db_pass := os.Getenv(("DB_PASS"))
	db_name := os.Getenv(("DB_NAME"))
	db_host := os.Getenv(("DB_HOST"))
	db_port := os.Getenv(("DB_PORT"))

	connStr := "postgres://" + db_user + ":" + db_pass + "@" + db_host + ":" + db_port + "/" + db_name + "?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	// Execute a query
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
