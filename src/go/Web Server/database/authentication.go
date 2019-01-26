package Database

import (
	"database/sql"
	"log"
)

func AuthenticateSignUp(db *sql.DB, username string, password string) int {

	cred := "INSERT INTO Credentials values (\"" + username + "\",\"" + password + "\")"

	_, err := db.Exec(cred)

	if err != nil {
		return 0
	}
	return 1
}

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
