package controllers

import (
	"APIGroupieTracker/models"
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
	err = tmpl.ExecuteTemplate(w, "404", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/favicon.ico" {
		tmpl, err := template.ParseFiles("templates/layout.html", "templates/404.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "404", nil)
		return
	}
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/about.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func FavoritesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/favorites.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/search.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "search", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	barcode := vars["barcode"]

	product, err := utils.FetchProductData(barcode)
	if err != nil {
		http.Error(w, "Failed to fetch product data", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/product.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout.html", product)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func FetchProductByBarcode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	barcode := vars["barcode"]
	product, err := utils.FetchProductData(barcode)
	if err != nil {
		http.Error(w, "Failed to fetch product data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func FetchAllProductsByCategory(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	products, err := utils.FetchProductsByCategory(category)
	if err != nil {
		http.Error(w, "Failed to fetch products by category", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func FetchProductsByCountry(w http.ResponseWriter, r *http.Request) {
	country := mux.Vars(r)["country"]
	products, err := utils.FetchProductsByCountry(country)
	if err != nil {
		http.Error(w, "Failed to fetch products by country", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func FetchSearchProducts(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("search")
	products, err := utils.FetchSearchResults(searchQuery)
	if err != nil {
		http.Error(w, "Failed to fetch search results", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func UpdateFavoritesHandler(w http.ResponseWriter, r *http.Request) {
	var updatedFavorites []models.Favorite
	err := json.NewDecoder(r.Body).Decode(&updatedFavorites)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = utils.WriteFavorites(updatedFavorites)
	if err != nil {
		http.Error(w, "Failed to update favorites", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
