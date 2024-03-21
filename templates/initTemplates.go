package templates

import (
	"html/template"
	"log"
)

var Temp *template.Template

func InitTemplate() {
	var err error
	Temp, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation des templates: %v", err)
	}
}
