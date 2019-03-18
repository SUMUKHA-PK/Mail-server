package DB

import (
	"fmt"

	"database/sql"
	"github.com/email-server/Web_Server/authorisation"
	"github.com/email-server/Web_Server/errorHandler"
)

/*
*    AddUserData adds the data for every session to the DB
*/

func AddUserData(userID string, loggedIn string, username string){

	dbPass := authorisation.ObtainPass()
	pass := "root:" + dbPass + "@/MailDB"
	db, err := sql.Open("mysql", pass)
	errorHandler.ErrorHandler(err)

	sessionData := "INSERT INTO sessions values(\"" + userID + "\",\"" + loggedIn + "\",\"" + username + "\")"

	fmt.Println(sessionData)
	_, err = db.Exec(sessionData)
	errorHandler.ErrorHandler(err)
}


func RemoveUserData(userID string, loggedIn string, username string) {

	dbPass := authorisation.ObtainPass()
	pass := "root:" + dbPass + "@/MailDB"
	db, err := sql.Open("mysql", pass)
	errorHandler.ErrorHandler(err)

	sessionData := "DELETE FROM sessions WHERE userID=\"" + userID + "\" AND loggedIn= \"" + loggedIn + "\" AND username=\"" + username + "\""

	fmt.Println(sessionData)
	_, err = db.Exec(sessionData)
	errorHandler.ErrorHandler(err)	
}