package authentication

import (
	"database/sql"
	"fmt"

	"../database"
	"../errorHandler"
)

func SignUpHelper(dbPass string, username string, password string, otp string) int {

	// Match OTP with the USER

	pass := "root:" + dbPass + "@/credentials"
	db, err := sql.Open("mysql", pass)

	errorHandler.ErrorHandler(err)

	x := Database.AuthenticateSignUp(db, username, password)

	pass = "root:" + dbPass + "@/MailDB"
	db, err = sql.Open("mysql", pass)

	y := Database.AddTable(db, username)

	fmt.Print(x)
	fmt.Print(y)
	if x == y && x == 1 {
		return 1
	} else {
		return 0
	}
}
