package routeHandlers

import (
	"net/http"

	"github.com/email-server/Web_Server/sessionHandler"
	"github.com/email-server/Web_Server/util"
)

// logoutHandler renders the logout page (home page) on button click
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	user,val := sessionHandler.CheckActiveSession(r)
	if user!=nil && val {
		sessionHandler.DestroySession(w,user)
	}
	util.RenderPage(w, "../webpages/static/index.html")
}