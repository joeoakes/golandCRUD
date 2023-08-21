package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// Define a struct to represent the data model
type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/database_name")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new user
	user := User{Name: "John Doe", Age: 25}
	err = createUser(db, user)
	if err != nil {
		log.Fatal(err)
	}

	// Read all users
	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Users:")
	for _, user := range users {
		fmt.Println(user.Name)
	}

	// Update a user
	err = updateUser(db, 1, "Jane Smith", 30)
	if err != nil {
		log.Fatal(err)
	}

	// Delete a user
	err = deleteUser(db, 1)
	if err != nil {
		log.Fatal(err)
	}
}

// Function to create a new user
func createUser(db *sql.DB, user User) error {
	_, err := db.Exec("INSERT INTO users(name, age) VALUES (?, ?)", user.Name, user.Age)
	return err
}

// Function to retrieve all users
func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Function to update a user
func updateUser(db *sql.DB, id int, name string, age int) error {
	_, err := db.Exec("UPDATE users SET name = ?, age = ? WHERE id = ?", name, age, id)
	return err
}

// Function to delete a user
func deleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
