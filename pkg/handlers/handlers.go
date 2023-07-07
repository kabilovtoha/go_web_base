package handlers

import (
	"github.com/kabilovtoha/go_web_base/pkg/config"
	"github.com/kabilovtoha/go_web_base/pkg/models"
	"github.com/kabilovtoha/go_web_base/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(a *config.AppConfig) (r *Repository) {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the Home page handler
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

// About is the About page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")

	StringMap := map[string]string{"remote_ip": remoteIP}

	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: StringMap,
	})
}
