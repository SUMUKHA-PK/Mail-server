package routeHandlers

import (
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/Mail-Server/DataBase"
	"github.com/SUMUKHA-PK/Mail-Server/mailHandler"
	"github.com/SUMUKHA-PK/Mail-Server/sessionHandler"
	"github.com/SUMUKHA-PK/Mail-Server/util"
)

func ComposeHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Routed to Compose page\n")
	if r.Method == "GET" {
		util.RenderPage(w, "webpages/static/compose.html")
	} else if r.Method == "POST" {
		user, val := sessionHandler.CheckActiveSession(r)
		if val {
			var data [][]string = mailHandler.ComposeHandler(w, r)
			DataBase.UpdateDB(data, User.UserName)
			username := user[0].UserName
			sessionHandler.SessionHandlerNew(w, r, username, "1")
		} else {

		}
	}
}
