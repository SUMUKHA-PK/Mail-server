package DB

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/email-server/Web_Server/authorisation"
	"github.com/email-server/Web_Server/errorHandler"
)

// GetEmails gets all the emails related to the username.
// inbox arg : 0 -> SentBox, 1-> Inbox
func GetEmails(username string, inbox string) *sql.Rows {

	dbPass := authorisation.ObtainPass()
	pass := "root:" + dbPass + "@/MailDB"
	db, err := sql.Open("mysql", pass)
	errorHandler.ErrorHandler(err)

	email := "SELECT * FROM " + username + " WHERE (inbox = " + inbox + ")"
	fmt.Println(email)
	rows, err := db.Query(email)
	errorHandler.ErrorHandler(err)

	return rows
}

// UpdateDB is used to handle sent and received emails from and to the users
func UpdateDB(data [][]string, username string) {
	dbPass := authorisation.ObtainPass()
	pass := "root:" + dbPass + "@/MailDB"
	db, err := sql.Open("mysql", pass)
	errorHandler.ErrorHandler(err)

	var user []string
	user = append(user, username)
	data = append([][]string{user}, data...)

	for i := 0; i < len(data[1]); i++ {
		email := "INSERT INTO " + data[1][i] + " VALUES(\"" + data[2][0] + "\",\"" + data[0][0] + "\",\"" + data[1][i] + "\",1,0)"
		_, err = db.Exec(email)
	}

	sender := strings.Join(data[1], ";")
	email := "INSERT INTO " + data[0][0] + " VALUES(\"" + data[2][0] + "\",\"" + data[0][0] + "\",\"" + sender + "\",0,1)"
	_, err = db.Exec(email)

}

// GetUserData gets the users session Data from the DB
func GetUserData(username string) *sql.Rows {

	dbPass := authorisation.ObtainPass()
	pass := "root:" + dbPass + "@/MailDB"
	db, err := sql.Open("mysql", pass)
	errorHandler.ErrorHandler(err)

	data := "SELECT * FROM sessions WHERE username = \"" + username + "\""

	rows, err := db.Query(data)
	errorHandler.ErrorHandler(err)

	return rows
}

// GetRoomData gives all the emails in the room for the given roomName
func GetRoomData(roomName string) *sql.Rows {
	dbPass := authorisation.ObtainPass()
	pass := "root:" + dbPass + "@/credentials"
	db, err := sql.Open("mysql", pass)
	errorHandler.ErrorHandler(err)

	data := "SELECT * FROM " + roomName

	rows, err := db.Query(data)
	errorHandler.ErrorHandler(err)

	return rows
}

// GetRoomsUser gives the rooms associated with the user
func GetRoomsUser(username string) *sql.Rows {
	dbPass := authorisation.ObtainPass()
	pass := "root:" + dbPass + "@/roomDB"
	db, err := sql.Open("mysql", pass)
	errorHandler.ErrorHandler(err)

	data := "SELECT * FROM rooms WHERE userName = \"" + username + "\""

	rows, err := db.Query(data)
	errorHandler.ErrorHandler(err)

	return rows
}

// GetUserRoomData gives the roomData associated with the roomName
func GetUserRoomData(roomName string) (*sql.Rows, error) {
	dbPass := authorisation.ObtainPass()
	pass := "root:" + dbPass + "@/roomDB"
	db, err := sql.Open("mysql", pass)
	errorHandler.ErrorHandler(err)

	data := "SELECT * FROM rooms WHERE roomName = \"" + roomName + "\""

	rows, err := db.Query(data)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
