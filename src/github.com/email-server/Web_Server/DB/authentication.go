package DB

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/email-server/Web_Server/authorisation"
	"github.com/email-server/Web_Server/errorHandler"
	"github.com/email-server/Web_Server/util"
)

func AuthenticateSignUp(db *sql.DB, username string, password string) int {

	cred := "INSERT INTO Credentials values (\"" + username + "\",\"" + password + "\")"

	_, err := db.Exec(cred)

	if err != nil {
		fmt.Print(err)
		return 0
	}
	return 1
}

// AddTable adds a table for every user to store their data
func AddTable(db *sql.DB, username string) int {

	user := "CREATE TABLE " + username + " (body LONGTEXT ,from_addr  VARCHAR(256) ,to_addr VARCHAR(256) ,inbox INT(2) ,sent INT(2) )"

	fmt.Print(user)

	_, err := db.Exec(user)

	if err != nil {
		fmt.Print(err)
		return 0
	}
	return 1
}

// AuthenticateLogin returns 0 or 1 based on existance of account in the DB(1) or not(0)
func AuthenticateLogin(db *sql.DB, username string, password string) int {

	cred := "SELECT username FROM Credentials WHERE username = \"" + username + "\" AND password = \"" + password + "\""

	rows, err := db.Query(cred)

	defer rows.Close()

	names := make([]string, 0)

	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			log.Fatal(err)
		}
		names = append(names, username)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if err != nil || len(names) == 0 {
		return 0
	}

	return 1
}

// Checks the existance of a session entry in the DB for the username
// Returns empty array and / or error on invalid queries

func CheckActiveSession(username string) ([]util.UserData, error) {

	dbPass := authorisation.ObtainPass()
	pass := "root:" + dbPass + "@/MailDB"
	db, err := sql.Open("mysql", pass)
	errorHandler.ErrorHandler(err)

	query := "SELECT * FROM sessions WHERE username=\"" + username + "\""
	
	fmt.Println(query)
	rows,err := db.Query(query)

	defer rows.Close()

	var user []util.UserData

	var userID string
	var loggedIn string
	var Username string

	for rows.Next() {
		if err := rows.Scan(&userID, &loggedIn, &Username); err != nil {
			log.Fatal(err)
		}
		user = append(user, util.UserData{ID: userID, LoggedIn: loggedIn, UserName: Username})
	}

	err = rows.Err()
	
	return user,err
	
}