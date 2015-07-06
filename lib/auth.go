package lib

// Authentication and Authorization

import (
	"github.com/gorilla/sessions"
	"net/http"
)

func IsAuthorized(store sessions.Store, r *http.Request, roles ...string) bool {
	session, _ := store.Get(r, "user-session")
	if session.Values["username"] != nil {
		if len(roles) > 0 {
			for _, role := range roles {
				if session.Values["role"] == role {
					return true
				}
			}
		} else {
			return true
		}
	}

	return false
}

func IsAuthenticated(store sessions.Store, w http.ResponseWriter, r *http.Request) bool {
	authenticated := false
	r.ParseForm()

	username := r.PostForm["username"][0]
	password := r.PostForm["password"][0]

	theUser := GetUser(username)

	if theUser != nil && theUser.password == password {
		session, _ := store.Get(r, "user-session")

		session.Values["username"] = theUser.username
		session.Values["role"] = theUser.role

		session.Save(r, w)

		authenticated = true
	}

	return authenticated
}
