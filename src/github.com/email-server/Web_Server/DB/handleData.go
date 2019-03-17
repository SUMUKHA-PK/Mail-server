package DB

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/email-server/Web_Server/authorisation"
	"github.com/email-server/Web_Server/errorHandler"
)

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

func UpdateDB(data [][]string) {
	dbPass := authorisation.ObtainPass()
	pass := "root:" + dbPass + "@/MailDB"
	db, err := sql.Open("mysql", pass)
	errorHandler.ErrorHandler(err)

	for i := 0; i < len(data[1]); i++ {
		email := "INSERT INTO " + data[1][i] + " VALUES(\"" + data[2][0] + "\",\"" + data[0][0] + "\",\"" + data[1][i] + "\",1,0)"
		fmt.Println(email)
		_, err = db.Exec(email)
	}

	sender := strings.Join(data[1], ";")
	email := "INSERT INTO " + data[0][0] + " VALUES(\"" + data[2][0] + "\",\"" + data[0][0] + "\",\"" + sender + "\",0,1)"
	fmt.Println(email)
	_, err = db.Exec(email)

}
