package routeHandlers

import (
	"../authentication"
)

func CreateSession(username string, password string) int {
	x := authentication.Authentication(username, password, 1)
	return x
}
