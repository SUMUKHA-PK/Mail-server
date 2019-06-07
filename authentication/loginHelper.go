package authentication

import (
	"database/sql"

	"github.com/SUMUKHA-PK/Mail-Server/DataBase"
	"github.com/SUMUKHA-PK/Mail-Server/errorHandler"
)

func LoginHelper(dbPass string, username string, password string) int {
	pass := "root:" + dbPass + "@/credentials"
	db, err := sql.Open("mysql", pass)

	errorHandler.ErrorHandler(err)

	x := DataBase.AuthenticateLogin(db, username, password)

	return x
}
