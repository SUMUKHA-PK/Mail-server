package routeHandlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/email-server/Web_Server/DB"
	"github.com/email-server/Web_Server/authentication"
	"github.com/email-server/Web_Server/mailHandler"
	"github.com/email-server/Web_Server/sessionHandler"
	"github.com/email-server/Web_Server/util"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		util.RenderPage(w, "../webpages/authentication/signup.html")
	} else if r.Method == "POST" {
		r.ParseForm()

		username := r.Form["username"]
		password := r.Form["password"]

		usernamestr := util.GetString(username)
		passwordstr := util.GetString(password)

		// We must get OTP from here
		otp := "1234"

		x := authentication.Authentication(usernamestr, passwordstr, 0, otp)

		if x == 1 {
			util.RenderPage(w, "../webpages/static/signupLogin.html")
		} else {
			util.RenderPage(w, "../webpages/static/signupFail.html")
		}
		// Things to do:
		// 1. Redirect to a page where then can enter their phonenumber
		// 2. Verify that phone number via OTP
		// 3. If OTP Matches, add a DB entry of username and password
		// 4. Then redirect them to the login page!

	}
}

// This, though a global variable exists in its own thread when running, so is fine.
var User util.UserData

// loginHandler handles authentication and session creation for every login
func loginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		util.RenderPage(w, "../webpages/authentication/login.html")
	} else if r.Method == "POST" {
		r.ParseForm()

		//authentication.LoginHelper("dbPass", )
		username := r.Form["username"]
		password := r.Form["password"]

		User.UserName = util.GetString(username)
		passwordstr := util.GetString(password)

		x := authentication.Authentication(User.UserName, passwordstr, 1, "")
		var sessionVar util.UserData
		
		if x == 2 {
			if sessionHandler.CheckActiveSession(r,User.UserName){
				log.Println("Found an active session")
				sessionVar = sessionHandler.GetActiveSession(User.UserName)
			} else {
				log.Println("Creating a new session")
				sessionVar = sessionHandler.CreateSession(w,User.UserName)
			}
			fmt.Println(sessionVar)
			sessionHandler.SessionHandlerNew(w, r, User.UserName, "1")
		} else {
			util.RenderPage(w, "../webpages/static/loginfail.html")
		}

	}
}

// logoutHandler renders the logout page (home page) on button click
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	user,err := DB.CheckActiveSession(User.UserName)
	if user!=nil && err==nil {
		sessionHandler.DestroySession(w,user)
	}
	util.RenderPage(w, "../webpages/static/index.html")
}

// HandlerFunc is the main route handler function that routes to different paths
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Home page of server
	if r.URL.Path == "/" {
		util.RenderPage(w, "../webpages/static/index.html")
		log.Print("Routed to Home page\n")
		// 2 paths :one hanldes the POST other handles the GET
	} else if r.URL.Path == "/login.html" || r.URL.Path == "/login" {
		loginHandler(w, r)
		log.Print("Routed to Login page\n")
	} else if r.URL.Path == "/signup.html" || r.URL.Path == "/signup" {
		signupHandler(w, r)
		log.Print("Routed to Signup page\n")
	} else if r.URL.Path == "/compose.html" {
		util.RenderPage(w, "../webpages/static/compose.html")
		log.Print("Routed to Compose page\n")
	} else if r.URL.Path == "/compose" {
		var data [][]string = mailHandler.ComposeHandler(w, r)
		DB.UpdateDB(data)
		username := User.UserName //Username must be obtained from the cookie not from the botched struct job
		sessionHandler.SessionHandlerNew(w, r, username, "1")
	} else if r.URL.Path == "/sentmail.html" {
		username := User.UserName //Username must be obtained from the cookie not from the botched struct job
		sessionHandler.SessionHandlerNew(w, r, username, "0")
		log.Print("Routed to Sentmail page\n")
	} else if r.URL.Path == "/loggedin.html" {
		username := User.UserName //Username must be obtained from the cookie not from the botched struct job
		sessionHandler.SessionHandlerNew(w, r, username, "1")
		log.Print("Routed to loggedin page\n")
	} else if r.URL.Path == "/logout" {
		logoutHandler(w, r)
		log.Print("Routed to Home page on logout\n")
	} else {
		w.WriteHeader(http.StatusNotFound) // Status code 404
		fmt.Fprint(w, "<h1>Error 404 : Page not found</h1>")
	}
}
