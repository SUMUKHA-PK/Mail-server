package routeHandlers

import (
	"net/http"

	"github.com/SUMUKHA-PK/Mail-Server/authentication"
	"github.com/SUMUKHA-PK/Mail-Server/util"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		util.RenderPage(w, "webpages/authentication/signup.html")
	} else if r.Method == "POST" {

		r.ParseForm()

		User.UserName = util.GetString(r.Form["username"])
		User.Password = util.GetString(r.Form["password"])
		User.PhoneNo = util.GetString(r.Form["phno"])
		User.Auth = false

		User.OTP = authentication.GenerateOTP()
		util.RenderPage(w, "webpages/authentication/otp.html")
	}
}

func signupHelper(username string, password string, phno string, w http.ResponseWriter, r *http.Request) {
	if User.Auth {
		x := authentication.Authentication(User.UserName, User.Password, 0)

		if x == 1 {
			util.RenderPage(w, "webpages/static/signupLogin.html")
		} else {
			util.RenderPage(w, "webpages/static/signupFail.html")
		}
	} else {
		//Not sure
	}
}

func VerifyAndRoute(username string, password string, phno string, w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		util.RenderPage(w, "webpages/authentication/otp.html")
	} else if r.Method == "POST" {
		r.ParseForm()
		otpUser := util.GetString(r.Form["otp"])

		if otpUser == User.OTP {
			User.Auth = true
			signupHelper(username, password, phno, w, r)
		} else {
			util.RenderPage(w, "webpages/static/otpFail.html")
		}
	}
}
