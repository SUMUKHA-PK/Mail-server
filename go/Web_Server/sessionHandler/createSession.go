package sessionHandler

import (
	"net/http"

	"../../securecookie"
	"../../sessions"
)

func CreateSession(w http.ResponseWriter, r *http.Request, username string, password string) *sessions.Session {

	// The session name will be the unique ID of the session
	sess_name := securecookie.GenerateRandomKey(32)

	var store = sessions.NewCookieStore(sess_name)

	session, _ := store.Get(r, "Session_name")

	// fmt.Println(string(sess_name))
	// errorHandler.HttpError(w, err)

	// Set all the variables here that will be needed to access to authorize
	session.Values["ID"] = sess_name

	return session // There will be no error here, you WILL get a session pointer back

}
