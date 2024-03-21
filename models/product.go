package models

// ProductResponse représente la structure de réponse globale pour un produit
type ProductResponse struct {
	Code    string  `json:"code"`
	Product Product `json:"product"`
	Status  int     `json:"status"`
}

// Product contient des informations détaillées sur le produit
type Product struct {
	ID                  string `json:"_id"`
	Allergens           string `json:"allergens"`
	BrandOwner          string `json:"brand_owner"`
	Brands              string `json:"brands"`
	Categories          string `json:"categories"`
	GenericName         string `json:"generic_name"`
	ProductName         string `json:"product_name"`
	ProductNameEn       string `json:"product_name_en"`
	IngredientsText     string `json:"ingredients_text"`
	NutritionGradeFr    string `json:"nutrition_grade_fr"`
	Origins             string `json:"origins"`
	ManufacturingPlaces string `json:"manufacturing_places"`
	Quantity            string `json:"quantity"`
	Packaging           string `json:"packaging"`
	Labels              string `json:"labels"`
	// Ajoutez d'autres champs si besoin de produit

	// Champs pour les images
	ImageFrontSmallURL       string   `json:"image_front_small_url"`
	ImageFrontThumbURL       string   `json:"image_front_thumb_url"`
	ImageFrontURL            string   `json:"image_front_url"`
	ImageIngredientsSmallURL string   `json:"image_ingredients_small_url"`
	ImageIngredientsThumbURL string   `json:"image_ingredients_thumb_url"`
	ImageIngredientsURL      string   `json:"image_ingredients_url"`
	ImageNutritionSmallURL   string   `json:"image_nutrition_small_url"`
	ImageNutritionThumbURL   string   `json:"image_nutrition_thumb_url"`
	ImageNutritionURL        string   `json:"image_nutrition_url"`
	ImageSmallURL            string   `json:"image_small_url"`
	ImageThumbURL            string   `json:"image_thumb_url"`
	ImageURL                 string   `json:"image_url"`
	Images                   []string `json:"images"`
}

// Endpoint pour la recherche de produit par code-barres
const (
	SearchProductEndpoint     = "https://world.openfoodfacts.org/api/v0/product/%s.json"
	ListProductsByCategory    = "https://world.openfoodfacts.org/category/%s.json"
	ProductsByCountryEndpoint = "https://world.openfoodfacts.org/country/%s.json"
)
