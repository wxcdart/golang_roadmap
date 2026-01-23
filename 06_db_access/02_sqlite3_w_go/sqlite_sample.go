// sqlite_sample.go
// Demonstrates basic usage of SQLite with Go using github.com/mattn/go-sqlite3.
//
// Dependency install (choose one):
//   go get github.com/mattn/go-sqlite3
//   OR (if using Go modules)
//   go mod tidy (after importing the package)

package main

//
// This example shows:
// - Opening a SQLite database
// - Creating a table
// - Inserting data
// - Querying data
// - Using the database/sql package with the go-sqlite3 driver
import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // SQLite driver (import for side effects)
)

func main() {
	// Open a new SQLite database file (creates it if it doesn't exist)
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		panic(err)
	}
	defer db.Close() // Always close the database when done

	// Create a table if it doesn't exist
	// AUTOINCREMENT makes id auto-increment for each new row
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	)`)
	if err != nil {
		panic(err)
	}

	// Insert some data using parameterized queries (prevents SQL injection)
	_, err = db.Exec(`INSERT INTO users (name, age) VALUES (?, ?)`, "Alice", 30)
	_, err2 := db.Exec(`INSERT INTO users (name, age) VALUES (?, ?)`, "Bob", 25)
	if err != nil || err2 != nil {
		panic(fmt.Errorf("insert error: %v %v", err, err2))
	}

	// Query the data
	rows, err := db.Query(`SELECT id, name, age FROM users`)
	if err != nil {
		panic(err)
	}
	defer rows.Close() // Always close rows when done

	fmt.Println("Users:")
	// Iterate over the result set
	for rows.Next() {
		var id, age int
		var name string
		// Scan copies the columns from the current row into the variables
		if err := rows.Scan(&id, &name, &age); err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
	// Always check for errors after iterating
	if err := rows.Err(); err != nil {
		panic(err)
	}
}
