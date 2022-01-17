package handlers

import (
	"net/http"

	"github.com/anjotadena/projectX/pkg/config"
	"github.com/anjotadena/projectX/pkg/render"
)

// TemplateData hold dataset
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

// Repo the repository used by the handlers
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	render.RenderTemplate(w, "home.page.html", &TemplateData{
		StringMap: stringMap,
	})
}
