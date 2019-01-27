package sessionHandler

import (
	
	"net/http"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore(os.Getenv("SESSION_KE