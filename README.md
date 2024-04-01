# APIGroupieTracker
 
 Installer : npm install -g cors-anywhere
Si le Site est en maintenance alors faut attendre pour les données :  https://world.openfoodfacts.org/cgi/search.pl?search_terms=savon&search_simple=1&action=process&json=1

exemple de Requete dans la Base De Donnée : chocolat


APIGroupieTracker
├── asset
│   ├── css
│   │   └── 404.css
│   │   └── search.css
│   │   └── style.css
│   │   └── favorites.css
│   │   └── pagination.css
│   │
│   └── img
│       └── ... logo.jpg
├── controllers
│   └── productController.go
├── models
│   └── product.go
├── data
│   └── favorites.json
├── routes
│   └── routes.go
├── templates
│   ├── about.html
│   ├── category.html
│   ├── favorites.html
│   ├── index.html
│   ├── initTemplates.go
│   ├── layout.html
│   ├── product.html
│   └── search.html
├── utils
│   └── httpClient.go
├── go.mod
├── go.sum
└── main.go


