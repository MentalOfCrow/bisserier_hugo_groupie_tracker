// Fonction pour effectuer une recherche
function searchProducts(page = 1) {
    var query = document.getElementById('search-input').value;
    if (!query.trim()) {
        alert('Veuillez entrer un terme de recherche');
        return;
    }

    // Ajout du paramètre de page à l'URL de recherche
    fetch(`https://world.openfoodfacts.org/cgi/search.pl?search_terms=${query}&page=${page}&page_size=20&search_simple=1&action=process&json=1`)
        .then(response => {
            // Vérifier si la réponse est OK
            if (!response.ok) {
                throw new Error('Échec de la requête');
            }
            // Convertir la réponse en JSON
            return response.json();
        })
        .then(data => {
            displaySearchResults(data.products);
            displayPagination(page, data.page_count);
        })
        .catch(error => {
            console.error('Erreur:', error);
            alert('Une erreur s\'est produite lors de la recherche');
        });
}

// Fonction pour afficher les résultats de recherche
function displaySearchResults(results) {
    var searchResultsDiv = document.getElementById('search-results');

    // Effacer les résultats précédents
    searchResultsDiv.innerHTML = '';

    // Vérifier s'il y a des résultats
    if (results.length === 0) {
        searchResultsDiv.innerHTML = 'Aucun résultat trouvé';
        return;
    }

    // Parcourir les résultats et les ajouter à la liste
    results.forEach(result => {
        var productBlock = document.createElement('div');
        productBlock.classList.add('product-block');

        var productName = document.createElement('p');
        productName.classList.add('product-name');
        productName.textContent = result.product_name;

        var productBrand = document.createElement('p');
        productBrand.classList.add('product-brand');
        productBrand.textContent = result.brands;

        var productImage = document.createElement('img');
        productImage.classList.add('product-image');
        productImage.src = result.image_front_url || 'path/to/default/image.jpg'; // Assurez-vous que ce champ correspond à celui de votre API
        productImage.alt = "Image du produit " + result.product_name;

        // Ajout de l'événement de clic sur le bloc de produit
        productBlock.addEventListener('click', function() {
            onProductClick(result.code);
        });

        productBlock.appendChild(productImage); // Ajoutez cette ligne pour inclure l'image
        productBlock.appendChild(productName);
        productBlock.appendChild(productBrand);

        searchResultsDiv.appendChild(productBlock);
    });
}

// Fonction pour afficher la pagination
function displayPagination(currentPage, totalPages) {
    var paginationContainer = document.getElementById('pagination-container');
    paginationContainer.innerHTML = ''; // Effacer la pagination précédente
    
    // Ajouter le bouton "Précédent"
    var previousButton = document.createElement('button');
    previousButton.innerText = 'Précédent';
    previousButton.addEventListener('click', function() {
        if (currentPage > 1) {
            searchProducts(currentPage - 1);
        }
    });
    paginationContainer.appendChild(previousButton);
    
    // Ajouter les numéros de page
    for (var i = 1; i <= totalPages; i++) {
        var pageButton = document.createElement('button');
        pageButton.innerText = i;
        if (i === currentPage) {
            pageButton.classList.add('current');
        }
        // Ajout de l'événement de clic sur le bouton de page
        pageButton.addEventListener('click', function() {
            searchProducts(parseInt(this.innerText));
        });
        paginationContainer.appendChild(pageButton);
    }
    
    // Ajouter le bouton "Suivant"
    var nextButton = document.createElement('button');
    nextButton.innerText = 'Suivant';
    nextButton.addEventListener('click', function() {
        if (currentPage < totalPages) {
            searchProducts(currentPage + 1);
        }
    });
    paginationContainer.appendChild(nextButton);
}

// Fonction pour gérer les clics sur les produits (à remplir avec le code nécessaire)
function onProductClick(productCode) {
    // Ajoutez ici le code pour gérer les clics sur les produits
    console.log("Clic sur le produit avec le code :", productCode);
}
