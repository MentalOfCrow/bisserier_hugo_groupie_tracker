// Fonction pour effectuer une recherche
function searchProducts() {
    // Récupérer le texte de la recherche
    var query = document.getElementById('search-input').value;

    // Vérifier si la requête est vide
    if (!query.trim()) {
        alert('Veuillez entrer un terme de recherche');
        return;
    }

    // Effectuer une requête GET à l'API avec la requête de recherche
    fetch('https://world.openfoodfacts.org/cgi/search.pl?search_terms=' + query + '&search_simple=1&action=process&json=1')
        .then(response => {
            // Vérifier si la réponse est OK
            if (!response.ok) {
                throw new Error('Échec de la requête');
            }
            // Convertir la réponse en JSON
            return response.json();
        })
        .then(data => {
            // Traiter les données de réponse
            displaySearchResults(data.products);
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

        productBlock.appendChild(productImage); // Ajoutez cette ligne pour inclure l'image
        productBlock.appendChild(productName);
        productBlock.appendChild(productBrand);

        searchResultsDiv.appendChild(productBlock);
    });
}

// Attacher un gestionnaire d'événements au formulaire de recherche
document.getElementById('search-form').addEventListener('submit', function(event) {
    event.preventDefault(); // Empêcher le formulaire de se soumettre normalement
    searchProducts(); // Appeler la fonction de recherche
});// Fonction pour effectuer une recherche avec pagination
function searchProducts(page = 1) {
    var query = document.getElementById('search-input').value;
    if (!query.trim()) {
        alert('Veuillez entrer un terme de recherche');
        return;
    }

    // Ajout du paramètre de page à l'URL de recherche
    fetch(`https://world.openfoodfacts.org/cgi/search.pl?search_terms=${query}&page=${page}&page_size=20&search_simple=1&action=process&json=1`)
        .then(response => response.json())
        .then(data => {
            displaySearchResults(data.products);
            displayPagination(data.page, data.page_count);
        })
        .catch(error => {
            console.error('Erreur:', error);
            alert('Une erreur s\'est produite lors de la recherche');
        });
}