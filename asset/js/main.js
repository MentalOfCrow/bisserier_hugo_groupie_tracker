// Fonction pour rechercher des produits
function searchProducts(page = 1) {
    var query = document.getElementById('search-input').value;
    var category = document.getElementById('category-filter').value;
    var country = document.getElementById('country-filter').value;
    
    var endpoint = `https://world.openfoodfacts.org/cgi/search.pl?search_terms=${query}&page=${page}&page_size=20&search_simple=1&action=process&json=1`;
    
    // Ajouter des filtres à l'URL si nécessaire
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

// Fonction pour afficher les résultats de la recherche
function displaySearchResults(results) {
    var searchResultsDiv = document.getElementById('search-results');
    searchResultsDiv.innerHTML = ''; // Efface les résultats précédents

    if (results.length === 0) {
        searchResultsDiv.innerHTML = 'Aucun résultat trouvé';
        return;
    }

    results.forEach(result => {
        var productBlock = document.createElement('div');
        productBlock.classList.add('product-block');

        // Création et ajout de l'image du produit
        var productImage = document.createElement('img');
        productImage.src = result.image_front_url;
        productImage.alt = `Image de ${result.product_name}`;
        productImage.classList.add('product-image');
        productBlock.appendChild(productImage);

        // Création et ajout du nom du produit et de la marque
        var productName = document.createElement('div');
        productName.classList.add('product-name');
        productName.innerHTML = `${result.product_name} - ${result.brands}`;
        productBlock.appendChild(productName);
        
        // Création du bouton de favori
        var favoriteButton = document.createElement('button');
        favoriteButton.innerText = '★';
        favoriteButton.classList.add('favorite-icon');
        favoriteButton.setAttribute('data-product', result.code);
        favoriteButton.addEventListener('click', function() {
            toggleFavorite(this, result.code, result.product_name, result.image_front_url);
        });
        productBlock.appendChild(favoriteButton);

        // Création et ajout des icônes des scores
        var scoresContainer = document.createElement('div');
        scoresContainer.classList.add('scores-container');

        // Ajout de l'icône Nutriscore
        var nutriscoreIcon = document.createElement('img');
        nutriscoreIcon.src = `/asset/img/nutriscore-${result.nutrition_grade_fr}.svg`;
        nutriscoreIcon.alt = `Nutriscore ${result.nutrition_grade_fr}`;
        nutriscoreIcon.classList.add('score-icon');
        scoresContainer.appendChild(nutriscoreIcon);

        // Ajout de l'icône NOVA
        var novaIcon = document.createElement('img');
        novaIcon.src = `/asset/img/nova-group-${result.nova_group}.svg`;
        novaIcon.alt = `Nova Group ${result.nova_group}`;
        novaIcon.classList.add('score-icon');
        scoresContainer.appendChild(novaIcon);

        // Ajout de l'icône EcoScore
        var ecoscoreIcon = document.createElement('img');
        ecoscoreIcon.src = `/asset/img/ecoscore-${result.ecoscore_grade}.svg`;
        ecoscoreIcon.alt = `EcoScore ${result.ecoscore_grade}`;
        ecoscoreIcon.classList.add('score-icon');
        scoresContainer.appendChild(ecoscoreIcon);

        productBlock.appendChild(scoresContainer);

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

// Fonction pour ajouter ou supprimer un produit des favoris
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

// Fonction pour mettre à jour la liste des favoris dans le stockage local
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

// Appel de la fonction pour mettre à jour les icônes de favori lors du chargement de la page
updateFavoriteIcon(productCode);

// Gestion des clics sur l'icône de favori
document.addEventListener('click', function(event) {
    if (event.target && event.target.classList.contains('favorite-icon')) {
        const productCode = event.target.getAttribute('data-product');
        console.log('Clic sur l\'icône de favori:', productCode);
        toggleFavorite(event.target, productCode);
    }
});
