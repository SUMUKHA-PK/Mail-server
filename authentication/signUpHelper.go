package authentication

import (
	"database/sql"
	"fmt"

	"github.com/SUMUKHA-PK/Mail-Server/DataBase"
	"github.com/SUMUKHA-PK/Mail-Server/errorHandler"
)

func SignUpHelper(dbPass string, username string, password string) int {

	// Match OTP with the USER

	pass := "root:" + dbPass + "@/credentials"
	db, err := sql.Open("mysql", pass)

	errorHandler.ErrorHandler(err)

	x := DataBase.AuthenticateSignUp(db, username, password)

	pass = "root:" + dbPass + "@/MailDB"
	db, err = sql.Open("mysql", pass)

	y := DataBase.AddTable(db, username)

	fmt.Print(x)
	fmt.Print(y)
	if x == y && x == 1 {
		return 1
	} else {
		return 0
	}
}
