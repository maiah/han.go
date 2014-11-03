package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Setup route handlers
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/home", home)

	// Start the server
	log.Println("Listening on port 5000")
	http.ListenAndServe(":5000", nil)
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

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//file := "pages/login.html"
		//http.ServeFile(w, r, file)

		fmt.Fprintf(w, createLoginForm(), "")

	} else if r.Method == "POST" {
		r.ParseForm()

		username := r.PostForm["username"][0]
		password := r.PostForm["password"][0]

		if username == "gohan" && password == "abc123" {
			http.Redirect(w, r, "/home", http.StatusFound)
		} else {
			fmt.Fprintf(w, createLoginForm(), "Invalid username/passowrd")
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	file := "pages/home.html"
	http.ServeFile(w, r, file)
}

func createLoginForm() string {
	return `<!DOCTYPE html>
            <html>
            <head>
                <title>Login</title>
            </head>
            <body>
                <p>%s</p>
                <form method="POST" action="/login">
                    <input type="text" name="username" placeholder="Username" />
                    <input type="password" name="password" placeholder="Password" />
                    <input type="submit" value="Login" />
                <form>
            </body>
            </html>`
}
