package models

// ProductResponse represents the complete response structure for a product
type ProductResponse struct {
	Code      string  `json:"code"`
	Product   Product `json:"product"`
	Status    int     `json:"status"`
	Page      int     `json:"page"`
	PageSize  int     `json:"page_size"`
	PageCount int     `json:"page_count"`
	Count     int     `json:"count"`
}

// Product contains detailed information about the product
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

	EcoScore            string `json:"ecoscore"`
	NovaGroup           string `json:"nova_group"`
	EnvironmentalImpact string `json:"environmental_impact"`

	// Image fields
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

// Image represents the structure for image data related to the product
type Image struct {
	Uploader  string     `json:"uploader"`
	UploadedT int        `json:"uploaded_t"`
	Sizes     ImageSizes `json:"sizes"`
}

// ImageSizes contains different sizes of the image
type ImageSizes struct {
	Thumb  ImageSize `json:"thumb"`
	Small  ImageSize `json:"small"`
	Medium ImageSize `json:"medium"`
	Large  ImageSize `json:"large"`
}

// ImageSize holds the URL and dimensions for a particular size
type ImageSize struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Endpoints for API requests
const (
	SearchProductEndpoint     = "https://world.openfoodfacts.org/api/v0/product/%s.json"
	ListProductsByCategory    = "https://world.openfoodfacts.org/category/%s.json"
	ProductsByCountryEndpoint = "https://world.openfoodfacts.org/country/%s.json"
)
