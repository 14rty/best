package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"log"
	_ "github.com/lib/pq"
)

var db *sql.DB

type User struct {
	ID       int
	FName    string
	LName    string
	Password int
}

const (
	host     = "my_db"
	port     = 5432
	user     = "db_user"
	password = "password4db"
	dbname   = "my_db"
)

func main() {
	fmt.Println("[INFO] Connecting to database")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
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

	user := User{
		FName:    "John",
		LName:    "Doe",
		Password: 123,
	}
	
	fmt.Println("[CHANGE] HELLOW")
	
	userID, err := addUser(user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID of added user: %v\n", userID)
	log.Fatal(http.ListenAndServe(":433", nil))
}

func addUser(alb User) (sql.Result, error) {
	fmt.Println(alb)
	result, err := db.Exec("INSERT INTO users (first_name, last_name, password) VALUES ($1, $2, $3) RETURNING Id", alb.FName, alb.LName, alb.Password)
	if err != nil {
		panic(err)
	}
	return result, nil
}

