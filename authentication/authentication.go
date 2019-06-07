package authentication

import (
	"github.com/SUMUKHA-PK/Mail-Server/authorisation"

	// should only be imported
	_ "github.com/go-sql-driver/mysql"
)

// Authentication is used to auth
func Authentication(username string, password string, decider int) int {

	DBPass := authorisation.ObtainPass()
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
