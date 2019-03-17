package authentication

import (
	"database/sql"

	"github.com/email-server/Web_Server/DB"
	"github.com/email-server/Web_Server/errorHandler"
)

func LoginHelper(dbPass string, username string, password string) int {
	pass := "root:" + dbPass + "@/credentials"
	db, err := sql.Open("mysql", pass)

	errorHandler.ErrorHandler(err)

	x := DB.AuthenticateLogin(db, username, password)

	return x
}
