package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, t string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + t)

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing on template", err)
		return
	}
}
