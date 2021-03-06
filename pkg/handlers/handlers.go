package handlers

import (
	"net/http"

	"github.com/HaoxuanXu/bookings/pkg/config"
	"github.com/HaoxuanXu/bookings/pkg/models"
	"github.com/HaoxuanXu/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}



func NewRepo(app *config.AppConfig) *Repository {
	return &Repository {
		App: app,
	}
}

func NewHandlers(repo *Repository) {
	Repo = repo
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

