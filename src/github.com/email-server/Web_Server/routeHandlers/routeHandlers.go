package routeHandlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/email-server/Web_Server/DB"
	"github.com/email-server/Web_Server/mailHandler"
	"github.com/email-server/Web_Server/sessionHandler"
	"github.com/email-server/Web_Server/util"
)

// This, though a global variable exists in its own thread when running, so is fine.
var User util.UserData
var Room util.RoomData

// HandlerFunc is the main route handler function that routes to different paths
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Home page of server
	if r.URL.Path == "/" {
		user, val := sessionHandler.CheckActiveSession(r)
		if val {
			log.Println("Found an active session")
			sessionVar := sessionHandler.GetActiveSession(user[0].UserName)
			sessionHandler.SessionHandlerNew(w, r, sessionVar.UserName, "1")
		} else {
			util.RenderPage(w, "../webpages/static/index.html")
			log.Print("Routed to Home page\n")
		}
	} else if r.URL.Path == "/login.html" || r.URL.Path == "/login" {
		LoginHandler(w, r)
		log.Print("Routed to Login page\n")
	} else if r.URL.Path == "/signup.html" || r.URL.Path == "/signup" {
		SignupHandler(w, r)
		log.Print("Routed to Signup page\n")
	} else if r.URL.Path == "/compose.html" {
		util.RenderPage(w, "../webpages/static/compose.html")
		log.Print("Routed to Compose page\n")
	} else if r.URL.Path == "/compose" {
		user, val := sessionHandler.CheckActiveSession(r)
		if val {
			var data [][]string = mailHandler.ComposeHandler(w, r)
			DB.UpdateDB(data, User.UserName)
			username := user[0].UserName
			sessionHandler.SessionHandlerNew(w, r, username, "1")
		} else {

		}
	} else if r.URL.Path == "/sentmail.html" {
		user, val := sessionHandler.CheckActiveSession(r)
		if val {
			username := user[0].UserName
			sessionHandler.SessionHandlerNew(w, r, username, "0")
			log.Print("Routed to Sentmail page\n")
		} else {

		}

	} else if r.URL.Path == "/loggedin.html" {
		user, val := sessionHandler.CheckActiveSession(r)
		if val {
			username := user[0].UserName
			sessionHandler.SessionHandlerNew(w, r, username, "1")
			log.Print("Routed to loggedin page\n")
		} else {

		}

	} else if r.URL.Path == "/logout" {
		LogoutHandler(w, r)
		log.Print("Routed to Home page on logout\n")
	} else if r.URL.Path == "/otpfail.html" {
		log.Print("Routed to OTP Fail page\n")
	} else if r.URL.Path == "/otp" || r.URL.Path == "/otp.html" {
		user, val := sessionHandler.CheckActiveSession(r)
		if val {
			log.Print("Routed to OTP verification page\n")
			VerifyAndRoute(user[0].UserName, user[0].Password, user[0].PhoneNo, w, r)
		} else {

		}

	} else if r.URL.Path == "/createRoom" || r.URL.Path == "/createRoom.html" {
		log.Print("Routed to room creation page\n")
		RoomCreationHandler(w, r)
	} else if r.URL.Path == "/room" || r.URL.Path == "/room.html" {
		user, val := sessionHandler.CheckActiveSession(r)
		if val {
			log.Print("Routed to room page\n")
			RoomHandler(w, r, user)
		} else {

		}

	} else if r.URL.Path == "/rooms" || r.URL.Path == "/rooms.html" {
		user, val := sessionHandler.CheckActiveSession(r)
		if val {
			log.Print("Routed to room page\n")
			RenderRoomChoicePage(w, r, user)
		} else {

		}

	} else if r.URL.Path == "/userRoom.html" {
		log.Print("Routed to user room page\n")
		RenderUserRoom(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound) // Status code 404
		fmt.Fprint(w, "<h1>Error 404 : Page not found</h1>")
	}
}
