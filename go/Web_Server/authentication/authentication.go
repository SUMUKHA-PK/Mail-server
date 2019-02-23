package authentication

import (
	_ "../../mysql"
	"../authorisation"
)

func Authentication(username string, password string, decider int, otp string) int {

	DBPass := authorisation.ObtainPass()
	// Decider 0 for Signup, 1 for Login
	if decider == 0 {
		if SignUpHelper(DBPass, username, password, otp) == 1 {
			return 1
		} else {
			return -1
		}
	} else if decider == 1 {
		if LoginHelper(DBPass, username, password) == 1 {
			return 2
		} else {
			return -2
		}
	} else {
		return 0
	}
}
