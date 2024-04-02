package routes

import (
	"APIGroupieTracker/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// SetRoutes configure les routes de l'application.
func SetRoutes() {
	r := mux.NewRouter()

	// Routes pour les pages = des templates
	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	r.HandleFunc("/about", controllers.AboutHandler).Methods("GET")
	r.HandleFunc("/favorites", controllers.FavoritesHandler).Methods("GET")
	r.HandleFunc("/search", controllers.SearchHandler).Methods("GET")
	r.HandleFunc("/product/{barcode}", controllers.ProductHandler).Methods("GET")
	r.HandleFunc("/product", controllers.ProductHandler).Methods("GET")

	//favorites.go : POST passe les paramètres dans le corps de la requête
	r.HandleFunc("/api/favorites", controllers.GetFavorites).Methods("GET")
	r.HandleFunc("/api/favorites/add", controllers.AddToFavorites).Methods("POST") //
	r.HandleFunc("/api/favorites/remove", controllers.RemoveFromFavorites).Methods("POST")

	// Routes pour EndPoints de l'api Open Food Facts
	r.HandleFunc("/api/product/{barcode}", controllers.FetchProductByBarcode).Methods("GET")
	r.HandleFunc("/api/products/category/{category}", controllers.FetchAllProductsByCategory).Methods("GET")
	r.HandleFunc("/api/products/country/{country}", controllers.FetchProductsByCountry).Methods("GET")

	// Permet tout simplement de Gerer les fichiers statiques =  (CSS, JS, images). '/asset/' est le chemin dans l'URL.
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("asset"))))

	r.NotFoundHandler = http.HandlerFunc(controllers.ErrorHandler)

	http.Handle("/", r)
}
