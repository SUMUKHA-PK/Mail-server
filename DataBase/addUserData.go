package DataBase

import (
	"fmt"

	"database/sql"

	"github.com/SUMUKHA-PK/Mail-Server/authorisation"
	"github.com/SUMUKHA-PK/Mail-Server/errorHandler"
)

/*
*    AddUserData adds the data for every session to the DB after clearing any past sessions in the DB
 */

func AddUserData(userID string, loggedIn string, username string) {

	dbPass := authorisation.ObtainPass()
	pass := "root:" + dbPass + "@/MailDB"
	db, err := sql.Open("mysql", pass)
	errorHandler.ErrorHandler(err)

	// Remove only if the time is expired
	session := "delete from sessions where username = \"" + username + "\""

	_, err = db.Exec(session)
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

func AddDataToDB(roomName string, username string, data string) error {
	dbPass := authorisation.ObtainPass()
	pass := "root:" + dbPass + "@/credentials"
	db, err := sql.Open("mysql", pass)
	errorHandler.ErrorHandler(err)

	sessionData := "INSERT INTO " + roomName + " values(\"" + data + "\",\"" + username + "\")"

	fmt.Println(sessionData)
	_, err = db.Exec(sessionData)
	return err
}
