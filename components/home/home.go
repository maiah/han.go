package home

import (
	"github.com/maiah/han.go/Godeps/_workspace/src/github.com/gorilla/sessions"
	"github.com/maiah/han.go/components/auth"
	"github.com/maiah/han.go/utils"
	"html/template"
	"log"
	"net/http"
)

func Home(store sessions.Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if auth.IsAuthorized(store, r) {
			homePage := "pages/home.html"
			t, err := template.ParseFiles(homePage)
			if err != nil {
				log.Fatal(err)
			}

			session, _ := store.Get(r, "user-session")

			t.Execute(w, utils.Page{Message: session.Values["username"].(string)})

		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}
