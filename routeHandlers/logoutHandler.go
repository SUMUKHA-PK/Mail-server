package routeHandlers

import (
	"net/http"

	"github.com/SUMUKHA-PK/Mail-Server/sessionHandler"
	"github.com/SUMUKHA-PK/Mail-Server/util"
)

// logoutHandler renders the logout page (home page) on button click
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	user, val := sessionHandler.CheckActiveSession(r)
	if user != nil && val {
		sessionHandler.DestroySession(w, user)
	}
	util.RenderPage(w, "../webpages/static/index.html")
}
