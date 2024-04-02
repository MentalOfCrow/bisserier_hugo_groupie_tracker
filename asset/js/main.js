function searchProducts(page = 1) {
    var query = document.getElementById('search-input').value;
    var category = document.getElementById('category-filter').value;
    var country = document.getElementById('country-filter').value;
    
    var endpoint = `https://world.openfoodfacts.org/cgi/search.pl?search_terms=${query}&page=${page}&page_size=20&search_simple=1&action=process&json=1`;
    
    // Ajoute des filtres à l'URL si nécessaire
    if(category) {
        endpoint += `&tagtype_0=categories&tag_contains_0=contains&tag_0=${category}`;
    }
    if(country) {
        endpoint += `&tagtype_1=countries&tag_contains_1=contains&tag_1=${country}`;
    }

    fetch(endpoint)
        .then(response => {
            if (!response.ok) {
                throw new Error('Échec de la requête');
            }
            return response.json();
        })
        .then(data => {
            console.log('Résultats de la recherche:', data.products);
            displaySearchResults(data.products);
            displayPagination(page, data.page_count);
        })
        .catch(error => {
            console.error('Erreur lors de la recherche:', error);
            alert('Une erreur s\'est produite lors de la recherche');
        });
}


function displaySearchResults(results) {
    var searchResultsDiv = document.getElementById('search-results');

    // Effacer les résultats précédents
    searchResultsDiv.innerHTML = '';

    // Vérifier s'il y a des résultats
    if (results.length === 0) {
        console.log('Aucun résultat trouvé');
        searchResultsDiv.innerHTML = 'Aucun résultat trouvé';
        return;
    }

    // Parcourir les résultats et les ajouter à la liste
    results.forEach(result => {
        console.log('Résultat:', result);
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
            console.log('Clic sur le produit:', result.code);
            onProductClick(result.code);
        });

        // Ajout de l'icône de favori
        var favoriteIcon = document.createElement('i');
        favoriteIcon.classList.add('far', 'fa-star', 'favorite-icon');
        favoriteIcon.setAttribute('data-product', result.code);
        favoriteIcon.setAttribute('data-name', result.product_name);
        favoriteIcon.setAttribute('data-image', result.image_front_url);
        favoriteIcon.addEventListener('click', function(event) {
            event.stopPropagation(); // Empêche la propagation de l'événement de clic sur le produit
            console.log('Clic sur l\'icône de favori:', result.code);
            toggleFavorite(this, result.code, result.product_name, result.image_front_url);
        });

        productBlock.appendChild(productImage); // Ajoutez cette ligne pour inclure l'image
        productBlock.appendChild(productName);
        productBlock.appendChild(productBrand);
        productBlock.appendChild(favoriteIcon); // Ajoute l'icône de favori

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
            console.log('Page précédente:', currentPage - 1);
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
            console.log('Page sélectionnée:', parseInt(this.innerText));
            searchProducts(parseInt(this.innerText));
        });
        paginationContainer.appendChild(pageButton);
    }
    
    // Ajouter le bouton "Suivant"
    var nextButton = document.createElement('button');
    nextButton.innerText = 'Suivant';
    nextButton.addEventListener('click', function() {
        if (currentPage < totalPages) {
            console.log('Page suivante:', currentPage + 1);
            searchProducts(currentPage + 1);
        }
    });
    paginationContainer.appendChild(nextButton);
}

// Fonction pour gérer les clics sur les produits (à remplir avec le code nécessaire)
function onProductClick(productCode) {
    // Rediriger l'utilisateur vers la page du produit sélectionné
    window.location.href = `/product?code=${productCode}`;
}

function toggleFavorite(element, productCode, productName, imageUrl) {
    let favorites = JSON.parse(localStorage.getItem('favorites')) || [];
    let isFavorite = favorites.some(fav => fav.code === productCode);
    
    const favoriteData = { code: productCode, name: productName, image_url: imageUrl };
    const url = isFavorite ? '/api/favorites/remove' : '/api/favorites/add';
    
    fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(favoriteData)
    }).then(response => {
        if (!response.ok) {
            return response.text().then(text => { throw new Error(`Failed to ${isFavorite ? 'remove' : 'add'} favorite: ${text}`); });
        }
        return response.json(); // Si vous envoyez une réponse JSON depuis le serveur
    }).then(data => {
        // Mettre à jour l'icône de l'étoile ici si nécessaire
        updateFavoritesList(isFavorite, favorites, favoriteData, element, productCode);
    }).catch(error => {
        console.error('Error toggling favorite:', error);
    });
}

function updateFavoritesList(isFavorite, favorites, favoriteData, element, productCode) {
    if (isFavorite) {
        favorites = favorites.filter(fav => fav.code !== productCode);
        element.classList.remove('fas'); // Suppose que 'fas' est la classe pour un favori
    } else {
        favorites.push(favoriteData);
        element.classList.add('fas'); // Ajoute la classe pour indiquer un favori
    }
    localStorage.setItem('favorites', JSON.stringify(favorites));
}



// Fonction pour mettre à jour l'icône de favori lors du chargement de la page
function updateFavoriteIcon(productCode) {
    let favorites = JSON.parse(localStorage.getItem('favorites')) || [];
    const index = favorites.findIndex(favorite => favorite.code === productCode);
    if (index !== -1) {
        const favoriteIcon = document.querySelector(`[data-product="${productCode}"]`);
        if (favoriteIcon) {
            favoriteIcon.classList.add('fas');
            console.log('Produit trouvé dans les favoris:', productCode);
        }
    }
}

// Événement pour gérer les clics sur l'icône de favori
document.addEventListener('click', function(event) {
    if (event.target && event.target.classList.contains('favorite-icon')) {
        const productCode = event.target.getAttribute('data-product');
        console.log('Clic sur l\'icône de favori:', productCode);
        toggleFavorite(event.target, productCode);
    }
});

// Appel de la fonction pour mettre à jour les icônes de favori lors du chargement de la page
updateFavoriteIcon(productCode);

function onProductClick(productCode) {
    window.location.href = '/product?code=' + productCode;
}

