package authentication

import (
	"database/sql"

	"../database"
	"../errorHandler"
)

func SignUpHelper(dbPass string, username string, password string, otp string) int {

	// Match OTP with the USER

	pass := "root:" + dbPass + "@/credentials"
	db, err := sql.Open("mysql", pass)

	errorHandler.ErrorHandler(err)

	x := Database.AuthenticateSignUp(db, username, password)

	return x
}
