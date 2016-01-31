package settings

import (
	"github.com/maiah/han.go/Godeps/_workspace/src/github.com/gorilla/sessions"
	"github.com/maiah/han.go/components/auth"
	"net/http"
)

func Settings(store sessions.Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if auth.IsAuthorized(store, r, "ADMIN") {
			file := "pages/settings.html"
			http.ServeFile(w, r, file)

		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}
