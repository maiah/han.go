package usermanager

// User Manager

import "github.com/maiah/han.go/user"

var users = []user.User{
	user.New(0, "gohan", "gohan", "Gohan", "Macariola", "ADMIN"),
	user.New(1, "maiah", "maiah", "Maiah", "Macariola", "USER"),
}

func GetUser(username string) (theUser *user.User) {
	for _, aUser := range users {
		if aUser.Username() == username {
			theUser = &aUser
			break
		}
	}

	return
}
