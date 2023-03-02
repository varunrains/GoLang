package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// Capital letter infront of a method will allow to use this package in any other package
// It is kind of public key word in GO module
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

}
