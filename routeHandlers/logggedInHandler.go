package routeHandlers

import (
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/Mail-Server/sessionHandler"
)

// LoggedInHandler works
func LoggedInHandler(w http.ResponseWriter, r *http.Request) {
	user, val := sessionHandler.CheckActiveSession(r)
	if val {
		username := user[0].UserName
		sessionHandler.SessionHandlerNew(w, r, username, "1")
		log.Print("Routed to loggedin page\n")
	} else {

	}
}
