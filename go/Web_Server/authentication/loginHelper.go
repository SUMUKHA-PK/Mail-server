package authentication

import (
	"database/sql"

	"../DB"
	"../errorHandler"
)

func LoginHelper(dbPass string, username string, password string) int {
	pass := "root:" + dbPass + "@/credentials"
	db, err := sql.Open("mysql", pass)

	errorHandler.ErrorHandler(err)

	x := DB.AuthenticateLogin(db, username, password)

	return x
}
