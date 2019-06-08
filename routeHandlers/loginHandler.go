package routeHandlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/Mail-Server/authentication"
	"github.com/SUMUKHA-PK/Mail-Server/sessionHandler"
	"github.com/SUMUKHA-PK/Mail-Server/util"
)

// LoginHandler handles authentication and session creation for every login
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		util.RenderPage(w, "webpages/authentication/login.html")
	} else if r.Method == "POST" {
		r.ParseForm()

		User.UserName = util.GetString(r.Form["username"])
		User.Password = util.GetString(r.Form["password"])

		x := authentication.Authentication(User.UserName, User.Password, 1)
		var sessionVar util.UserData

		if x == 2 {
			log.Println("Creating a new session")
			sessionVar = sessionHandler.CreateSession(w, User.UserName)
			fmt.Println(sessionVar)
			sessionHandler.SessionHandlerNew(w, r, User.UserName, "1")
		} else {
			util.RenderPage(w, "webpages/static/loginfail.html")
		}
	}
}
