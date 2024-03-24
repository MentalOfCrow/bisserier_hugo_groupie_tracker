package controllers

import (
	"APIGroupieTracker/utils"
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/404.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Renvoie la page "À propos" au client.
	err = tmpl.ExecuteTemplate(w, "404", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HomeHandler gère la page d'accueil.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Chargez le template pour la page d'accueil.
	if r.URL.Path != "/" && r.URL.Path != "/flavicon.ico" {
		println("404 from home")
		tmpl, err := template.ParseFiles("templates/layout.html", "templates/404.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "404", nil)
		return
	}
	println("home valid")
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Renvoie la page d'accueil au client.
	err = tmpl.ExecuteTemplate(w, "layout.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// AboutHandler gère la page "À propos".
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Chargez le template pour la page "À propos".
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/about.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Renvoie la page "À propos" au client.
	err = tmpl.ExecuteTemplate(w, "layout.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// FavoritesHandler gère la page des favoris.
func FavoritesHandler(w http.ResponseWriter, r *http.Request) {
	// Chargez le template pour la page des favoris.
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/favorites.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Renvoie la page des favoris au client.
	err = tmpl.ExecuteTemplate(w, "layout.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// SearchHandler gère la page de recherche.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Chargez le template pour la page de recherche.
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/search.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Renvoie la page de recherche au client avec le template "search".
	err = tmpl.ExecuteTemplate(w, "search", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// ProductHandler gère la page d'un produit.
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	barcode := vars["barcode"]
	// Chargez le template pour la page du produit.
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/product.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Chargez les données du produit.
	product, err := utils.FetchProductData(barcode)
	if err != nil {
		http.Error(w, "Failed to fetch product data", http.StatusInternalServerError)
		return
	}
	// Renvoie la page du produit au client avec les données du produit.
	err = tmpl.ExecuteTemplate(w, "layout.html", product)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// FetchProductByBarcode gère la requête pour récupérer un produit par code-barres.
func FetchProductByBarcode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	barcode := vars["barcode"]
	product, err := utils.FetchProductData(barcode)
	if err != nil {
		http.Error(w, "Failed to fetch product data", http.StatusInternalServerError)
		return
	}
	// Renvoie les données du produit en tant que réponse JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// FetchAllProductsByCategory gère la requête pour récupérer tous les produits par catégorie.
func FetchAllProductsByCategory(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	products, err := utils.FetchProductsByCategory(category)
	if err != nil {
		http.Error(w, "Failed to fetch products by category", http.StatusInternalServerError)
		return
	}
	// Renvoie les produits en tant que réponse JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// FetchProductsByCountry gère la requête pour récupérer tous les produits par pays.
func FetchProductsByCountry(w http.ResponseWriter, r *http.Request) {
	country := mux.Vars(r)["country"]
	products, err := utils.FetchProductsByCountry(country)
	if err != nil {
		http.Error(w, "Failed to fetch products by country", http.StatusInternalServerError)
		return
	}
	// Renvoie les produits en tant que réponse JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// FetchSearchProducts effectue une recherche et retourne les produits
func FetchSearchProducts(w http.ResponseWriter, r *http.Request) {
	// Récupérer le paramètre de recherche depuis l'URL
	searchQuery := r.URL.Query().Get("search")

	// Appeler la fonction d'utilité pour effectuer la recherche
	products, err := utils.FetchSearchResults(searchQuery)
	if err != nil {
		http.Error(w, "Failed to fetch search results", http.StatusInternalServerError)
		return
	}

	// Écrire la réponse en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
