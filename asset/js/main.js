document.addEventListener('DOMContentLoaded', function() {
    // Ceci sont des ecouteurs d"évenemments permet de faire la recherche lorsqu'un filtre est modifié
    document.getElementById('category-filter').addEventListener('change', function() {
        searchProducts();
    });
    document.getElementById('country-filter').addEventListener('change', function() {
        searchProducts();
    });
    document.getElementById('nutriscore-filter').addEventListener('change', function() {
        searchProducts();
    });
    document.getElementById('nova-filter').addEventListener('change', function() {
        searchProducts();
    });
    document.getElementById('ecoscore-filter').addEventListener('change', function() {
        searchProducts();
    });
});

function searchProducts(page = 1) {
    var query = document.getElementById('search-input').value;
    var category = document.getElementById('category-filter').value;
    var country = document.getElementById('country-filter').value;
    var nutriscore = document.getElementById('nutriscore-filter').value;
    var nova = document.getElementById('nova-filter').value;
    var ecoscore = document.getElementById('ecoscore-filter').value;
    
    var endpoint = `https://world.openfoodfacts.org/cgi/search.pl?search_terms=${query}&page=${page}&page_size=20&search_simple=1&action=process&json=1`;
    
    // Ce sont les filtres à l'URL si besoin
    if(category) {
        endpoint += `&tagtype_0=categories&tag_contains_0=contains&tag_0=${category}`;
    }
    if(country) {
        endpoint += `&tagtype_1=countries&tag_contains_1=contains&tag_1=${country}`;
    }
    if(nutriscore) {
        endpoint += `&tagtype_2=nutrition_grades&tag_contains_2=contains&tag_2=${nutriscore}`;
    }
    if(nova) {
        endpoint += `&tagtype_3=nova_groups&tag_contains_3=contains&tag_3=${nova}`;
    }
    if(ecoscore) {
        endpoint += `&tagtype_4=ecoscore_grade&tag_contains_4=contains&tag_4=${ecoscore}`;
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
    searchResultsDiv.innerHTML = ''; // Permet d'efface les résultats précedent 

    if (results.length === 0) {
        searchResultsDiv.innerHTML = 'Aucun résultat trouvé';
        return;
    }

    results.forEach(result => {
        var productBlock = document.createElement('div');
        productBlock.classList.add('product-block');

        // Permet laCréation et ajout de l'image du produit
        var productImage = document.createElement('img');
        productImage.src = result.image_front_url;
        productImage.alt = `Image de ${result.product_name}`;
        productImage.classList.add('product-image');
        productBlock.appendChild(productImage);

        // C'est durant la Création et ajout du nom du produit et de la marque
        var productName = document.createElement('div');
        productName.classList.add('product-name');
        productName.innerHTML = `${result.product_name} - ${result.brands}`;
        productBlock.appendChild(productName);
        
        // Ca permet de crée un bouton de favori
        var favoriteButton = document.createElement('button');
        favoriteButton.innerText = '★';
        favoriteButton.classList.add('favorite-icon');
        favoriteButton.setAttribute('data-product', result.code);
        favoriteButton.addEventListener('click', function() {
            toggleFavorite(this, result.code, result.product_name, result.image_front_url);
        });
        productBlock.appendChild(favoriteButton);

        // Alors ca c'est la Création et ajout des icônes des scores
        var scoresContainer = document.createElement('div');
        scoresContainer.classList.add('scores-container');

         // Ca permet Ajoute de l'icône Nutriscore meme ceux qui ont pas de données
         var nutriscoreIcon = document.createElement('img');
         var nutriscoreImageSrc = result.nutrition_grade_fr ? `/asset/img/nutriscore-${result.nutrition_grade_fr}.svg` : `/asset/img/nutriscore-unknown.svg`;
         nutriscoreIcon.src = nutriscoreImageSrc;
         nutriscoreIcon.alt = `Nutriscore ${result.nutrition_grade_fr || 'Inconnu'}`;
         nutriscoreIcon.classList.add('score-icon');
         scoresContainer.appendChild(nutriscoreIcon);
 
         // Ca permet Ajoute de l'icône NOVA meme ceux qui ont pas de données
         var novaIcon = document.createElement('img');
         var novaImageSrc = result.nova_group ? `/asset/img/nova-group-${result.nova_group}.svg` : `/asset/img/nova-group-unknown.svg`;
         novaIcon.src = novaImageSrc;
         novaIcon.alt = `Nova Group ${result.nova_group || 'Inconnu'}`;
         novaIcon.classList.add('score-icon');
         scoresContainer.appendChild(novaIcon);
 
         // Ca permet Ajoute de l'icône EcoScore meme ceux qui ont pas de données
         var ecoscoreIcon = document.createElement('img');
         var ecoscoreImageSrc = result.ecoscore_grade ? `/asset/img/ecoscore-${result.ecoscore_grade}.svg` : `ecoscore-not-applicable.svg`;
         ecoscoreIcon.src = ecoscoreImageSrc;
         ecoscoreIcon.alt = `EcoScore ${result.ecoscore_grade || 'Inconnu'}`;
         ecoscoreIcon.classList.add('score-icon');
         scoresContainer.appendChild(ecoscoreIcon);

        productBlock.appendChild(scoresContainer);

        searchResultsDiv.appendChild(productBlock);
    });
}

// La Fonction en résumer =  permet d'afficher la pagination
function displayPagination(currentPage, totalPages) {
    var paginationContainer = document.getElementById('pagination-container');
    paginationContainer.innerHTML = ''; // Ca Effacer la pagination précédente
    
    //C'est l' Ajouter le bouton "Précédent"
    var previousButton = document.createElement('button');
    previousButton.innerText = 'Précédent';
    previousButton.addEventListener('click', function() {
        if (currentPage > 1) {
            console.log('Page précédente:', currentPage - 1);
            searchProducts(currentPage - 1);
        }
    });
    paginationContainer.appendChild(previousButton);
    
    // C'est l'Ajouter les numéros de page
    for (var i = 1; i <= totalPages; i++) {
        var pageButton = document.createElement('button');
        pageButton.innerText = i;
        if (i === currentPage) {
            pageButton.classList.add('current');
        }
        //Ca permet tous simple l' Ajout de l'événement de clic sur le bouton de page
        pageButton.addEventListener('click', function() {
            console.log('Page sélectionnée:', parseInt(this.innerText));
            searchProducts(parseInt(this.innerText));
        });
        paginationContainer.appendChild(pageButton);
    }
    
    // Comme précedent sauf que ca permet l'Ajouter le bouton "Suivant"
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

// En résumer ca permet de gérer les clics sur les produits
function onProductClick(productCode) {
    // Ca Rediriger l'utilisateur vers la page du produit sélectionné important
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
        element.classList.remove('fas'); // Awesome bibliotheque 'fas' est la classe pour un favori
    } else {
        favorites.push(favoriteData);
        element.classList.add('fas'); // Awesome bibliotheque permet d'ajouter la classe pour indiquer un favori
    }
    localStorage.setItem('favorites', JSON.stringify(favorites));
}

// Ca permet en résumer la Fonction pour mettre à jour l'icône de favori lors du chargement de la page
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

//Alors ca c'est Événement pour gérer les clics sur l'icône de favori
document.addEventListener('click', function(event) {
    if (event.target && event.target.classList.contains('favorite-icon')) {
        const productCode = event.target.getAttribute('data-product');
        console.log('Clic sur l\'icône de favori:', productCode);
        toggleFavorite(event.target, productCode);
    }
});

// Ca permet d'Appel de la fonction pour mettre à jour les icônes de favori lors du chargement de la page
updateFavoriteIcon(productCode);

function onProductClick(productCode) {
    window.location.href = '/product?code=' + productCode;
}
