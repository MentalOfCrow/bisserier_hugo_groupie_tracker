package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// FetchProductData récupère les données d'un produit via son code-barres.
func FetchProductData(barcode string) (interface{}, error) {
	url := fmt.Sprintf("https://world.openfoodfacts.org/api/v0/product/%s.json", barcode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FetchProductsByCategory récupère tous les produits d'une catégorie donnée.
func FetchProductsByCategory(category string) ([]interface{}, error) {
	url := fmt.Sprintf("https://world.openfoodfacts.org/category/%s.json", category)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Products []interface{} `json:"products"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Products, nil
}

// FetchProductsByCountry récupère tous les produits d'un pays spécifique.
func FetchProductsByCountry(country string) ([]interface{}, error) {
	url := fmt.Sprintf("https://world.openfoodfacts.org/country/%s.json", country)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Products []interface{} `json:"products"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Products, nil
}

// HandleProductsByCategory affiche les produits par catégorie.
func HandleProductsByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	products, err := FetchProductsByCategory(category)
	if err != nil {
		http.Error(w, "Failed to fetch products by category", http.StatusInternalServerError)
		return
	}

	// Écrire la réponse en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// HandleProductsByCountry affiche les produits par pays.
func HandleProductsByCountry(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	products, err := FetchProductsByCountry(country)
	if err != nil {
		http.Error(w, "Failed to fetch products by country", http.StatusInternalServerError)
		return
	}

	// Écrire la réponse en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func FetchSearchResults(searchQuery string) ([]interface{}, error) {
	// Construire l'URL de recherche avec la requête de recherche
	searchURL := fmt.Sprintf("https://world.openfoodfacts.org/cgi/search.pl?search_terms=%s&search_simple=1&action=process&json=1", url.QueryEscape(searchQuery))

	// Effectuer une requête GET à l'API de recherche
	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Lire le corps de la réponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Analyser le corps de la réponse JSON
	var searchResponse struct {
		Products []interface{} `json:"products"`
	}
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		return nil, err
	}

	// Renvoyer les produits trouvés
	return searchResponse.Products, nil
}
