package authentication

import (
	"database/sql"

	"../database"
	"../errorHandler"
)

func LoginHelper(dbPass string, username string, password string) int {
	pass := "root:" + dbPass + "@/credentials"
	db, err := sql.Open("mysql", pass)

	errorHandler.ErrorHandler(err)

	x := Database.AuthenticateLogin(db, username, password)

	return x
}
