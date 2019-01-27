package routeHandlers

import (
	"os"
	"fmt"
	"net/http"
	// "strings"
	// "reflect"
	
	"../errorHandler"

	//"../authentication"
)


func renderPage(w http.ResponseWriter, pageName string) {

	f, err := os.Open(pageName)
	errorHandler.ErrorHandler(err)
	b1:= make([]byte, 100000)
	 n1, err := f.Read(b1)
	 errorHandler.ErrorHandler(err)
	fmt.Fprintf(w, string(b1))
	fmt.Printf("n1 = %d\n", n1)
	// check(err)
}


func signupHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		renderPage(w, "../webpages/authentication/signup.html")
	} else if r.Method == "POST" {
		r.ParseForm()

		newusername := r.Form["username"]
		newpassword := r.Form["password"]

		usernamestr := ""
		passwordstr := ""

		for i:= 0; i < len(username); i++ {
			usernamestr = usernamestr + username[i]
		}

		for i:= 0; i < len(password); i++ {
			passwordstr = passwordstr + password[i]
		}
		
		// Things to do: 
		// 1. Redirect to a page where then can enter their phonenumber
		// 2. Verify that phone number via OTP
		// 3. If OTP Matches, add a DB entry of username and password
		// 4. Then redirect them to the login page!

	}
}


func loginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET"  {
		renderPage(w, "../webpages/authentication/login.html")
	} else if r.Method == "POST" {
		r.ParseForm()

		//authentication.LoginHelper("dbPass", )
		username := r.Form["username"]
		password := r.Form["password"]

		usernamestr := ""
		passwordstr := ""

		for i:= 0; i < len(username); i++ {
			usernamestr = usernamestr + username[i]
		}

		for i:= 0; i < len(password); i++ {
			passwordstr = passwordstr + password[i]
		}

		// If x is 1, then create session. else don't create session
		// x := LoginHelper("dbPass", username, password)

		x := 0
		if x == 1 {
			// createSession(username, password);
		} else {
			renderPage(w, "../webpages/static/loginfail.html")
		}


	}
}


func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		renderPage(w, "../webpages/static/index.html")
		
	} else if r.URL.Path == "/login.html" || r.URL.Path == "/login" {
		loginHandler(w, r);
	
	} else if r.URL.Path == "/signup.html" {
		signupHandler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound) // Status code 404
		fmt.Fprint(w, "<h1>Error 404 : Page not found</h1>")
	}
}


