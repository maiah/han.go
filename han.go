package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/maiah/han.go/lib"
	"html/template"
	"log"
	"net/http"
)

var (
	store = sessions.NewCookieStore([]byte("smallelephantandbigfly"))
)

func main() {
	// Setup route handlers
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/home", home)
	http.HandleFunc("/settings", settings)

	// Start the server
	log.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000",
		context.ClearHandler(http.DefaultServeMux)))
}

// Route Handlers

func index(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if len(path) >= 7 {
		path = r.URL.Path[0:7]
	}

	if path == "/public" {
		public(w, r) // invoke static file handler
	} else {
		fmt.Fprintf(w, "welcome my han.go")
	}
}

// Custom static file handling

func public(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[1:]
	http.ServeFile(w, r, file)
}

type page struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	if lib.IsAuthorized(store, r) {
		http.Redirect(w, r, "/home", http.StatusFound)

	} else {
		loginPage := "pages/login.html"

		if r.Method == "GET" {
			t, _ := template.ParseFiles(loginPage)
			t.Execute(w, nil)

		} else if r.Method == "POST" {
			if lib.IsAuthenticated(store, w, r) {
				http.Redirect(w, r, "/home", http.StatusFound)

			} else {
				t, err := template.ParseFiles(loginPage)

				if err != nil {
					log.Fatal(err)
				}

				if t.Execute(w,
					&page{Message: "Invalid username/password"}) != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	session.Options = &sessions.Options{MaxAge: -1}
	session.Values = nil
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusFound)
}

func home(w http.ResponseWriter, r *http.Request) {
	if lib.IsAuthorized(store, r) {
		homePage := "pages/home.html"
		t, err := template.ParseFiles(homePage)
		if err != nil {
			log.Fatal(err)
		}

		session, _ := store.Get(r, "user-session")

		t.Execute(w, &page{Message: session.Values["username"].(string)})

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func settings(w http.ResponseWriter, r *http.Request) {
	if lib.IsAuthorized(store, r, "ADMIN") {
		file := "pages/settings.html"
		http.ServeFile(w, r, file)

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
