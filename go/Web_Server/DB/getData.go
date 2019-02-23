package DB

import (
	"database/sql"

	"../authorisation"
	"../errorHandler"
)

func GetEmails(username string, inbox string) *sql.Rows {

	dbPass := authorisation.ObtainPass()
	pass := "root:" + dbPass + "@/MailDB"
	db, err := sql.Open("mysql", pass)

	errorHandler.ErrorHandler(err)
	email := "SELECT * FROM " + username + " WHERE (inbox = " + inbox + ")"

	rows, err := db.Query(email)
	errorHandler.ErrorHandler(err)

	return rows
}