package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/HaoxuanXu/bookings/pkg/config"
	"github.com/HaoxuanXu/bookings/pkg/handlers"
	"github.com/HaoxuanXu/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumer = ":8080"
var app config.AppConfig
var session *scs.SessionManager




func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Printf("Starting on port %s\n", portNumer)

	server := &http.Server {
		Addr: portNumer,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}