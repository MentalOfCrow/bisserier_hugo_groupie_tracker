package main

import (
	routes "APIGroupieTracker/routes" // Importez le routeur.
	"APIGroupieTracker/templates"     // Importez les templates.
	"fmt"
	"net/http"
)

func main() {
	templates.InitTemplate() // Initialisez les templates.

	routes.SetRoutes() // Configurez les routes définies dans le fichier 'routes.go'.

	// Gère les fichiers statiques (CSS, JS, images). '/asset/' est le chemin dans l'URL.
	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset"))))

	// Démarrage du serveur sur le port 8080.
	fmt.Println("Serveur démarré sur http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil) // nil signifie que nous utilisons le routeur par défaut.
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur :", err)
	}
}
