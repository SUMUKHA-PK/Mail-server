package routeHandlers

import (
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/Mail-Server/sessionHandler"
)

// OTPHandler works
func OTPHandler(w http.ResponseWriter, r *http.Request) {
	user, val := sessionHandler.CheckActiveSession(r)
	if val {
		log.Print("Routed to OTP verification page\n")
		VerifyAndRoute(user[0].UserName, user[0].Password, user[0].PhoneNo, w, r)
	} else {

	}
}
