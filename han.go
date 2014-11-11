package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
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
	http.ListenAndServe(":5000", context.ClearHandler(http.DefaultServeMux))
}

func index(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if len(path) >= 7 {
		path = r.URL.Path[0:7]
	}

	if path == "/public" {
		public(w, r)
	} else {
		fmt.Fprintf(w, "welcome my han.go")
	}
}

func public(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[1:]
	http.ServeFile(w, r, file)
}

type Page struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	if isAuthorized(r) {
		http.Redirect(w, r, "/home", http.StatusFound)

	} else {
		loginPage := "pages/login.html"

		if r.Method == "GET" {
			t, _ := template.ParseFiles(loginPage)
			t.Execute(w, nil)

		} else if r.Method == "POST" {
			if isAuthenticated(w, r) {
				http.Redirect(w, r, "/home", http.StatusFound)

			} else {
				t, _ := template.ParseFiles(loginPage)
				t.Execute(w, &Page{Message: "Invalid username/password"})
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
	if isAuthorized(r) {
		file := "pages/home.html"
		http.ServeFile(w, r, file)

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func settings(w http.ResponseWriter, r *http.Request) {
	if isAuthorized(r) {
		file := "pages/settings.html"
		http.ServeFile(w, r, file)

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func isAuthorized(r *http.Request) bool {
	session, _ := store.Get(r, "user-session")
	fmt.Printf("isAuthorized? %s\n", session.Values["username"])
	return session.Values["username"] != nil
}

func isAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	authenticated := false
	r.ParseForm()

	username := r.PostForm["username"][0]
	password := r.PostForm["password"][0]

	if username == "gohan" && password == "abc123" {
		session, _ := store.Get(r, "user-session")
		session.Values["username"] = username
		session.Save(r, w)

		authenticated = true
	}

	return authenticated
}
