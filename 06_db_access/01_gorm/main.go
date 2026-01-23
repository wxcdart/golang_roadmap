package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User is a simple GORM model for the example
type User struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	// opens an in-memory sqlite database; requires gorm and the sqlite driver
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// automigrate and basic create/query
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	db.Create(&User{Name: "Alice", Age: 30})

	var u User
	db.First(&u)
	fmt.Printf("found user: %+v\n", u)
}
