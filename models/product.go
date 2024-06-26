package models

// ProductResponse  = représente la structure complète de la réponse pour un produit
type ProductResponse struct {
	Code      string  `json:"code"`
	Product   Product `json:"product"`
	Status    int     `json:"status"`
	Page      int     `json:"page"`
	PageSize  int     `json:"page_size"`
	PageCount int     `json:"page_count"`
	Count     int     `json:"count"`
}

// Product = contient des informations détaillées sur un produit
type Product struct {
	ID                      string `json:"_id"`
	Allergens               string `json:"allergens"`
	BrandOwner              string `json:"brand_owner"`
	Brands                  string `json:"brands"`
	Categories              string `json:"categories"`
	GenericName             string `json:"generic_name"`
	ProductName             string `json:"product_name"`
	ProductNameEn           string `json:"product_name_en"`
	IngredientsText         string `json:"ingredients_text"`
	NutritionGradeFr        string `json:"nutrition_grade_fr"`
	Origins                 string `json:"origins"`
	ManufacturingPlaces     string `json:"manufacturing_places"`
	Quantity                string `json:"quantity"`
	Packaging               string `json:"packaging"`
	Labels                  string `json:"labels"`
	EcoScore                string `json:"ecoscore"`
	NovaGroup               string `json:"nova_group"`
	EnvironmentalImpact     string `json:"environmental_impact"`
	IngredientsAnalysis     string `json:"ingredients_analysis"`
	NutrientLevels          string `json:"nutrient_levels"`
	ProductQuantity         string `json:"product_quantity"`
	PackagingTags           string `json:"packaging_tags"`
	LabelsTags              string `json:"labels_tags"`
	AllergensTags           string `json:"allergens_tags"`
	TracesTags              string `json:"traces_tags"`
	StoresTags              string `json:"stores_tags"`
	CountriesTags           string `json:"countries_tags"`
	IngredientsFromPalmOilN int    `json:"ingredients_from_palm_oil_n"`

	// Champs pour les images = Permet d'afficher les images
	ImageFrontSmallURL       string  `json:"image_front_small_url"`
	ImageFrontThumbURL       string  `json:"image_front_thumb_url"`
	ImageFrontURL            string  `json:"image_front_url"`
	ImageIngredientsSmallURL string  `json:"image_ingredients_small_url"`
	ImageIngredientsThumbURL string  `json:"image_ingredients_thumb_url"`
	ImageIngredientsURL      string  `json:"image_ingredients_url"`
	ImageNutritionSmallURL   string  `json:"image_nutrition_small_url"`
	ImageNutritionThumbURL   string  `json:"image_nutrition_thumb_url"`
	ImageNutritionURL        string  `json:"image_nutrition_url"`
	ImageSmallURL            string  `json:"image_small_url"`
	ImageThumbURL            string  `json:"image_thumb_url"`
	ImageURL                 string  `json:"image_url"`
	Images                   []Image `json:"images"`
}

// ProductDetail = Ca représente les détails complets d'un produit
type ProductDetail struct {
	ID            string `json:"_id"`
	Allergens     string `json:"allergens"`
	Brand         string `json:"brands"`
	Categories    string `json:"categories"`
	GenericName   string `json:"generic_name"`
	ProductNameFr string `json:"product_name"`
	ProductNameEn string `json:"product_name_en"`
	Ingredients   string `json:"ingredients_text"`
	// Pas besoin d'en rajouter pour le moment pour éviter de surcharger l'api Open Food Facts
}

// Image = représente la structure pour les données d'image liées au produit
type Image struct {
	Uploader  string     `json:"uploader"`
	UploadedT int        `json:"uploaded_t"`
	Sizes     ImageSizes `json:"sizes"`
}

// ImageSizes = contient les différentes tailles d'image disponibles
type ImageSizes struct {
	Thumb  ImageSize `json:"thumb"`
	Small  ImageSize `json:"small"`
	Medium ImageSize `json:"medium"`
	Large  ImageSize `json:"large"`
}

// ImageSize =  détient l'URL et les dimensions pour une taille spécifique d'image
type ImageSize struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Favorite des produits
type Favorite struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	ImageURL      string `json:"image_url"`
	Brand         string `json:"brand"`
	Categories    string `json:"categories"`
	GenericName   string `json:"generic_name"`
	Ingredients   string `json:"ingredients"`
	Quantity      string `json:"quantity"`
	Packaging     string `json:"packaging"`
	Labels        string `json:"labels"`
	Allergens     string `json:"allergens"`
	NutriScore    string `json:"nutri_score"`
	NOVA          string `json:"nova"`
	EcoScore      string `json:"eco_score"`
	Origins       string `json:"origins"`
	Manufacturing string `json:"manufacturing"`
	Stores        string `json:"stores"`
	Countries     string `json:"countries"`
}

// Permet ProductDetails probleme sur API open food facts problmes
type ProductDetails struct {
	ID            string `json:"_id"`
	Allergens     string `json:"allergens"`
	Brand         string `json:"brand"`
	Categories    string `json:"categories"`
	GenericName   string `json:"generic_name"`
	ProductNameFr string `json:"product_name_fr"`
	ProductNameEn string `json:"product_name_en"`
	Ingredients   string `json:"ingredients"`
}

// Endpoints API OPEN FOOD FACTS
const (

	// 3 Requetes Minimums
	SearchProductEndpoint     = "https://world.openfoodfacts.org/api/v0/product/%s.json"
	ListProductsByCategory    = "https://world.openfoodfacts.org/category/%s.json"
	ProductsByCountryEndpoint = "https://world.openfoodfacts.org/country/%s.json"
	// Bonus
	ProductsByBrandEndpoint    = "https://world.openfoodfacts.org/brand/%s.json"
	ProductsByLabelEndpoint    = "https://world.openfoodfacts.org/label/%s.json"
	ProductsByAdditiveEndpoint = "https://world.openfoodfacts.org/additive/%s.json"
	ProductsByAllergenEndpoint = "https://world.openfoodfacts.org/allergen/%s.json"
	UserContributionsEndpoint  = "https://world.openfoodfacts.org/contribtour/%s.json"
	GeneralSearchEndpoint      = "https://world.openfoodfacts.org/cgi/search.pl"
)
