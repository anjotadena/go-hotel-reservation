package handlers

import (
	"net/http"

	"github.com/anjotadena/projectX/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}
