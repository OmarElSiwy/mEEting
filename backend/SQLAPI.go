package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"golang.org/x/crypto/bcrypt"
)

const (
	server   = "eleceng.database.windows.net" // Update with your server name
	port     = 1433                           // Default port for Azure SQL Database
	user     = "BabaOmar"
	password = "OmarSeesAll607"
	database = "EEStudent"
)

type TableType int

const (
	Forums TableType = iota
	Messages
)

var tableNames = map[TableType]string{
	Forums:   "forums",
	Messages: "messages",
}

type User struct {
	Username string
	Email    string
	Password string
}

type ForumPost struct {
	Title       string
	Description string
	CreatedBy   int
}

type Message struct {
	Content    string
	SenderID   int
	ReceiverID int
}

func main() {
	// print information that we are using to open the server
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	// Open the Server, and plan to close it at the end
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// Ping the Server to make sure it is active
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Print success if the database is connected
	fmt.Println("Successfully connected to the database")
}

func createUser(db *sql.DB, user User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, email, password_hash) VALUES (@p1, @p2, @p3)`
	_, err = db.Exec(query, user.Username, user.Email, string(hashedPassword))
	if err != nil {
		return err
	}

	fmt.Println("User created successfully")
	return nil
}

func authenticateUser(db *sql.DB, email, password string) (bool, error) {
	var hashedPassword string
	query := `SELECT password_hash FROM users WHERE email = @p1`
	row := db.QueryRow(query, email)
	err := row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // User not found
		}
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, nil // Incorrect password
	}

	return true, nil // Successful authentication
}

func deleteUser(db *sql.DB, userID int) error {
	query := `DELETE FROM users WHERE id = @p1`
	_, err := db.Exec(query, userID)
	if err != nil {
		return err
	}

	fmt.Println("User deleted successfully")
	return nil
}

func postToTable(db *sql.DB, tableType TableType, post interface{}) error {
	tableName, ok := tableNames[tableType]
	if !ok {
		return fmt.Errorf("invalid table type")
	}

	switch tableType {
	case Forums:
		forumPost, ok := post.(ForumPost)
		if !ok {
			return fmt.Errorf("invalid type for forum post")
		}
		query := `INSERT INTO forums (title, description, created_by) VALUES (@p1, @p2, @p3)`
		_, err := db.Exec(query, forumPost.Title, forumPost.Description, forumPost.CreatedBy)
		if err != nil {
			return err
		}
	case Messages:
		message, ok := post.(Message)
		if !ok {
			return fmt.Errorf("invalid type for message")
		}
		query := `INSERT INTO messages (content, sender_id, receiver_id) VALUES (@p1, @p2, @p3)`
		_, err := db.Exec(query, message.Content, message.SenderID, message.ReceiverID)
		if err != nil {
			return err
		}
	}

	fmt.Println("Post added successfully to", tableName)
	return nil
}
