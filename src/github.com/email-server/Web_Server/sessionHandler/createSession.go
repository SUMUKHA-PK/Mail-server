/*
File to create a new session.
*/

package sessionHandler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/email-server/Web_Server/util"
)

/*
CreateSession returns a new session pointer by creating a new session or getting an existing one.
Username and password are the credentials of the User
*/

func CreateSession(w http.ResponseWriter, r *http.Request, username string, password string) *sessions.Session {

	// The session name will be the unique ID of the session
	sess_token := util.GenerateRandomString(32)
	// errorHandler.ErrorHandler(err)

	fmt.Print(sess_token)
	var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

	session, _ := store.Get(r, sess_token)

	// Set all the variables here that will be needed to access to authorize
	session.Values["ID"] = sess_token
	session.Values["loggedin"] = true

	session.Save(r, w)
	return session

}

// func GetCurrentSession() *sessions.Session {
// 	return
// }
