package user

func New(id int, username, password, firstname, lastname, role string) User {
	return User{id, username, password, firstname, lastname, role}
}

// User Manager
type User struct {
	id        int
	username  string
	password  string
	firstname string
	lastname  string
	role      string
}

func (u User) Id() int {
	return u.id
}

func (u User) Username() string {
	return u.username
}

func (u User) Password() string {
	return u.password
}

func (u User) Firstname() string {
	return u.firstname
}

func (u User) Lastname() string {
	return u.lastname
}

func (u User) Role() string {
	return u.role
}
