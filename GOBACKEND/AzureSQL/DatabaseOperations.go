package AzureSQL

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	server   = ""
	user     = ""
	password = ""
	port     = ""
	database = ""
)

type User struct {
	UserID   int
	Email    string
	Username string
}

var db *sql.DB

func init() {
	var err error
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		os.Getenv("DB_SERVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), 1433, os.Getenv("DB_NAME"))
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
}

func insertUser(email, password, username string) error {
	query := "INSERT INTO Users (Email, Password, Username) VALUES (?, ?, ?)"
	_, err := db.Exec(query, email, password, username)
	return err
}

func insertPost(title, description, imageURL string) error {
	query := "INSERT INTO Posts (Title, Description, ImageURL) VALUES (?, ?, ?)"
	_, err := db.Exec(query, title, description, imageURL)
	return err
}

func fetchUsers() ([]User, error) {
	users := []User{}
	rows, err := db.Query("SELECT UserID, Email, Username FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.UserID, &u.Email, &u.Username); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func searchUserByUsername(username string) ([]User, error) {
	users := []User{}
	query := "SELECT UserID, Email, Username FROM Users WHERE Username = @username"
	rows, err := db.Query(query, sql.Named("username", username))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.UserID, &u.Email, &u.Username); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
