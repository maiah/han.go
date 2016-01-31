package main

import (
	"github.com/maiah/han.go/Godeps/_workspace/src/github.com/gorilla/context"
	"github.com/maiah/han.go/Godeps/_workspace/src/github.com/gorilla/sessions"
	authRoutes "github.com/maiah/han.go/components/auth/routes"
	"github.com/maiah/han.go/components/home"
	"github.com/maiah/han.go/components/index"
	"github.com/maiah/han.go/components/settings"
	"log"
	"net/http"
	"os"
)

func main() {
	store := sessions.NewCookieStore([]byte("smallelephantandbigfly"))

	// Setup route handlers
	http.HandleFunc("/", index.Index)
	http.HandleFunc("/login", authRoutes.Login(store))
	http.HandleFunc("/logout", authRoutes.Logout(store))
	http.HandleFunc("/home", home.Home(store))
	http.HandleFunc("/settings", settings.Settings(store))

	port := os.Getenv("PORT")

	// Start the server
	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port,
		context.ClearHandler(http.DefaultServeMux)))
}
