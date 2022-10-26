package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

type User struct {
	ID       int64
	FName    string
	LName    string
	Password float32
}

const (
	host     = "my_db"
	port     = 3306
	user     = "db_user"
	password = "password4db"
	dbname   = "my_db"
)

func main() {
	fmt.Println("[INFO] Connecting to database")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer fmt.Println("[INFO] Connection closed")
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("[INFO] Connection established")

	userID, err := addUser(User{
		FName:    "John",
		LName:    "Doe",
		Password: 123,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID of added user: %v\n", userID)
}

func addUser(alb User) (int64, error) {
	result, err := db.Exec("INSERT INTO users (first_name, last_name, password) VALUES (?, ?, ?)", alb.FName, alb.LName, alb.Password)
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	return id, nil
}
