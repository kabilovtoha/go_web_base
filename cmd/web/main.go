package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/kabilovtoha/go_web_base/pkg/config"
	"github.com/kabilovtoha/go_web_base/pkg/handlers"
	"github.com/kabilovtoha/go_web_base/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber string = "127.0.0.1:8000"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.IsProduction = false

	session = scs.New()
	session.Lifetime = 7 * 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return
	}
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

	//mux := routes(&app)
	//http.Handle("/", mux)
	//_ = http.ListenAndServe(portNumber, nil)
}
