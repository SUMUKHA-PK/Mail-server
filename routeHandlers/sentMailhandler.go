package routeHandlers

import (
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/Mail-Server/sessionHandler"
)

// SentMailHandler works
func SentMailHandler(w http.ResponseWriter, r *http.Request) {
	user, val := sessionHandler.CheckActiveSession(r)
	if val {
		username := user[0].UserName
		sessionHandler.SessionHandlerNew(w, r, username, "0")
		log.Print("Routed to Sentmail page\n")
	} else {

	}
}
