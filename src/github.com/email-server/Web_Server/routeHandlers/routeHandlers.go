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

// This, though a global variable exists in its own thread when running, so is fine.
var User util.UserData

func signupHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		util.RenderPage(w, "../webpages/authentication/signup.html")
	} else if r.Method == "POST" {

		r.ParseForm()

		User.UserName = util.GetString(r.Form["username"])
		User.Password = util.GetString(r.Form["password"])
		User.PhoneNo = util.GetString(r.Form["phno"])
		User.Auth = false

		User.OTP = authentication.GenerateOTP()
		util.RenderPage(w,"../webpages/authentication/otp.html")
	}
}

func signupHelper(username string, password string, phno string,w http.ResponseWriter, r *http.Request){
	if User.Auth {
		x := authentication.Authentication(User.UserName,User.Password, 0)

		if x == 1 {
			util.RenderPage(w, "../webpages/static/signupLogin.html")
		} else {
			util.RenderPage(w, "../webpages/static/signupFail.html")
		}
	}else{

	}
}

// loginHandler handles authentication and session creation for every login
func loginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		util.RenderPage(w, "../webpages/authentication/login.html")
	} else if r.Method == "POST" {
		r.ParseForm()
	
		User.UserName = util.GetString(r.Form["username"])
		User.Password = util.GetString(r.Form["password"])

		x := authentication.Authentication(User.UserName, User.Password, 1)
		var sessionVar util.UserData
		
		if x == 2 {
			log.Println("Creating a new session")
			sessionVar = sessionHandler.CreateSession(w,User.UserName)
			fmt.Println(sessionVar)
			sessionHandler.SessionHandlerNew(w, r, User.UserName, "1")
		} else {
			util.RenderPage(w, "../webpages/static/loginfail.html")
		}
	}
}

func verifyAndRoute(username string, password string, phno string,w http.ResponseWriter, r *http.Request){

	if r.Method == "GET" {
		util.RenderPage(w,"../webpages/authentication/otp.html")
	}else if r.Method == "POST"{
		r.ParseForm()
		otpUser := util.GetString(r.Form["otp"])

		if otpUser == User.OTP {
			User.Auth = true
			signupHelper(username,password,phno,w,r)			
		}else{
			util.RenderPage(w, "../webpages/static/otpFail.html")
		}
	}
}

// logoutHandler renders the logout page (home page) on button click
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	user,err := DB.CheckActiveSession()
	if user!=nil && err==nil {
		sessionHandler.DestroySession(w,user)
	}
	util.RenderPage(w, "../webpages/static/index.html")
}

// HandlerFunc is the main route handler function that routes to different paths
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Home page of server
	if r.URL.Path == "/" {
		if sessionHandler.CheckActiveSession(r,User.UserName){
			log.Println("Found an active session")
			sessionVar := sessionHandler.GetActiveSession(User.UserName)
			sessionHandler.SessionHandlerNew(w,r,sessionVar.UserName,"1")
		}else{
			util.RenderPage(w, "../webpages/static/index.html")
			log.Print("Routed to Home page\n")
		}
		// 2 paths :one handles the POST other handles the GET
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
		username := User.UserName
		sessionHandler.SessionHandlerNew(w, r, username, "1")
	} else if r.URL.Path == "/sentmail.html" {
		username := User.UserName
		sessionHandler.SessionHandlerNew(w, r, username, "0")
		log.Print("Routed to Sentmail page\n")
	} else if r.URL.Path == "/loggedin.html" {
		username := User.UserName
		sessionHandler.SessionHandlerNew(w, r, username, "1")
		log.Print("Routed to loggedin page\n")
	} else if r.URL.Path == "/logout" {
		logoutHandler(w, r)
		log.Print("Routed to Home page on logout\n")
	} else if r.URL.Path == "/otpfail.html"{
		log.Print("Routed to OTP Fail page\n")
	} else if r.URL.Path == "/otp" || r.URL.Path == "/otp.html"{
		log.Print("Routed to OTP verification page\n")
		verifyAndRoute(User.UserName,User.Password,User.PhoneNo,w,r)
	}else {
		w.WriteHeader(http.StatusNotFound) // Status code 404
		fmt.Fprint(w, "<h1>Error 404 : Page not found</h1>")
	}
}

