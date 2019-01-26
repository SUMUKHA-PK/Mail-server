package Database

import (
	"database/sql"
	"fmt"

	"../errorHandler"
)

func AuthenticateFromDb(db *sql.DB, username string, password string) int {
	cred := "INSERT INTO Credentials values (\"" + username + "\",\"" + password + "\")"

	_, err := db.Exec(cred)

	errorHandler.ErrorHandler(err)

	fmt.Print(cred)
	return 0
}
