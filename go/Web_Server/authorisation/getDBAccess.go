package authorisation

var DBPass string

func GetPass(dbPass string) {
	DBPass = dbPass
}

// NEED TO check who is asking for the password
func ObtainPass() string {
	return DBPass
}
