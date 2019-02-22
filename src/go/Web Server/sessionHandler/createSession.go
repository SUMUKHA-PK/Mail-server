package sessionHandler

import (
	"net/http"

	"../../securecookie"
	"../../sessions"
)

func CreateSession(w http.ResponseWriter, r *http.Request, username string, password string) int {

	var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

	session, _ := store.Get(r, "session-name")
	session.Values["foo"] = "bar"

	return 2 // Return SUCCESS on completion
}
