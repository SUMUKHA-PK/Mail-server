package authentication

import (
	"database/sql"

	"../Database"
	"../errorHandler"
)

func SignUpHelper(dbPass string, username string, password string) int {

	pass := "root:" + dbPass + "@/credentials"
	db, err := sql.Open("mysql", pass)

	errorHandler.ErrorHandler(err)

	x := Database.AuthenticateSignUp(db, username, password)

	return x
}
