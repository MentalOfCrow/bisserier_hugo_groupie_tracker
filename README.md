# APIGroupieTracker
 
 Installer : npm install -g cors-anywhere
Si le Site est en maintenance alors faut attendre pour les données :  https://world.openfoodfacts.org/cgi/search.pl?search_terms=savon&search_simple=1&action=process&json=1

# Description

Open Food Facts est une base de données en ligne collaborative qui recueille des informations sur les produits alimentaires du monde entier. 

# Utilisation 
API OPEN FOOD FACT : API Tres Tres Tres complexe et souvent en maintenance ce qui durant certain heures les produits s'affiche pas !
Pour la Barre de Recherche : Veuillez Attendre que les produits chaques car Bases de Données énorme 
= Technique Console Double click sur la Barre de Recherche Puis attendre ....
	
exemple de Requete simple dans la Base De Donnée : chocolat = sans filtre = tablette de chocolat
exemple : Chocolat + choix france + choix Boisson = CandieUP

Particularité ! :  on peux mettre que des filtres sans rien marquer dans la barre de recherche ca fonctionne !


# Structure Du Projet : 

APIGroupieTracker
├── asset
│   ├── css
│   │   └── 404.css
│   │   └── search.css
│   │   └── style.css
│   │   └── favorites.css
│   │   └── pagination.css
│   │
│   └── img (Sans Images)
│       └── ... logo.jpg 
├── controllers
│   └── productController.go
│   └── favorites.go
├── models
│   └── product.go
├── data
│   └── favorites.json
│   └── SearchProductEndpoint.json
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


