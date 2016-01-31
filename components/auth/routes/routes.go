package routes

import (
	"github.com/maiah/han.go/Godeps/_workspace/src/github.com/gorilla/sessions"
	"github.com/maiah/han.go/components/auth"
	"github.com/maiah/han.go/utils"
	"html/template"
	"log"
	"net/http"
)

func Login(store sessions.Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if auth.IsAuthorized(store, r) {
			http.Redirect(w, r, "/home", http.StatusFound)

		} else {
			loginPage := "pages/login.html"

			if r.Method == "GET" {
				t, _ := template.ParseFiles(loginPage)
				t.Execute(w, nil)

			} else if r.Method == "POST" {
				if auth.IsAuthenticated(store, w, r) {
					http.Redirect(w, r, "/home", http.StatusFound)

				} else {
					t, err := template.ParseFiles(loginPage)

					if err != nil {
						log.Fatal(err)
					}

					if t.Execute(w,
						utils.Page{Message: "Invalid username/password"}) != nil {
						log.Fatal(err)
					}
				}
			}
		}

	}

}

func Logout(store sessions.Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "user-session")
		session.Options = &sessions.Options{MaxAge: -1}
		session.Values = nil
		session.Save(r, w)

		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
