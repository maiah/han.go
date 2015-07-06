package lib

// User Manager

type User struct {
	id        int
	username  string
	password  string
	firstname string
	lastname  string
	role      string
}

var users = []User{
	User{0, "gohan", "gohan", "Gohan", "Macariola", "ADMIN"},
	User{1, "maiah", "maiah", "Maiah", "Macariola", "USER"},
}

func GetUser(username string) (theUser *User) {
	for _, aUser := range users {
		if aUser.username == username {
			theUser = &aUser
			break
		}
	}

	return
}
