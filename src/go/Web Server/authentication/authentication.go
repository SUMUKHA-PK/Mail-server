package authentication

import (
	_ "../../mysql"
)

var DBPass string

func GetPass(dbPass string) {
	DBPass = dbPass
}

func Authentication(username string, password string, decider int) int {

	// Decider 0 for Signup, 1 for Login
	if decider == 0 {
		if SignUpHelper(DBPass, username, password) == 1 {
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
