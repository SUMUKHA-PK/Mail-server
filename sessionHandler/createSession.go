/*
File to create a new session.
*/

package sessionHandler

import (
	"log"
	"net/http"
	"time"

	"github.com/SUMUKHA-PK/Mail-Server/DataBase"
	"github.com/SUMUKHA-PK/Mail-Server/util"
)

/*
CreateSession returns a new session pointer by creating a new session or getting an existing one.
Username and password are the credentials of the User
*/

func CreateSession(w http.ResponseWriter, username string) util.UserData {

	var user util.UserData

	user.ID = util.GenerateRandomString(32)
	user.UserName = username
	user.LoggedIn = "true"
	expiration := time.Now().Add(time.Hour)
	cookie := http.Cookie{Name: "session-id", Value: user.ID, Expires: expiration}
	http.SetCookie(w, &cookie)

	DataBase.AddUserData(user.ID, user.LoggedIn, user.UserName)

	return user

}

/*
CheckActiveSession queries the DB for the current users data, checks the cookie
for corresponding data and returns true if the data is as expected, else
returns false.
*/
func CheckActiveSession(r *http.Request) ([]util.UserData, bool) {

	cookie, err := r.Cookie("session-id")
	if err != nil {
		return []util.UserData{}, false
	} else {
		user, err := DataBase.CheckActiveSession(cookie.Value)
		if err != nil {
			return []util.UserData{}, false
		}
		return user, true
	}
}

/*
Returns the active session after CheckActiveSession confirms that the
data existing in the DB is consistent user data.
*/
func GetActiveSession(username string) util.UserData {

	var user util.UserData

	rows := DataBase.GetUserData(username)

	var userID string
	var loggedIn string
	var Username string

	for rows.Next() {
		err := rows.Scan(&userID, &loggedIn, &Username)
		if err != nil {
			log.Println(err)
		}

		user.ID = userID
		user.UserName = Username
		user.LoggedIn = loggedIn
	}

	return user
}

/*
DestroySession destroys an existing session before logging the user out
*/
func DestroySession(w http.ResponseWriter, user []util.UserData) {

	expiration := time.Now().Add(0)
	cookie := http.Cookie{Name: user[0].ID, Value: user[0].UserName, Expires: expiration}
	http.SetCookie(w, &cookie)

	DataBase.RemoveUserData(user[0].ID, user[0].LoggedIn, user[0].UserName)
}
