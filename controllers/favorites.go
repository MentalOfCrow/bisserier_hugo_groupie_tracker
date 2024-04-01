package controllers

import (
	"APIGroupieTracker/models"
	"APIGroupieTracker/utils"
	"encoding/json"
	"net/http"
)

// Ajouter un favori
func AddToFavorites(w http.ResponseWriter, r *http.Request) {
	var favorite models.Favorite
	err := json.NewDecoder(r.Body).Decode(&favorite)
	if err != nil {
		http.Error(w, "Erreur lors de la décodification du favori", http.StatusBadRequest)
		return
	}

	favorites, err := utils.ReadFavorites()
	if err != nil {
		http.Error(w, "Erreur lors de la lecture des favoris", http.StatusInternalServerError)
		return
	}

	favorites = append(favorites, favorite)
	err = utils.WriteFavorites(favorites)
	if err != nil {
		http.Error(w, "Erreur lors de l'écriture des favoris", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Définir le type de contenu de la réponse
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Favori ajouté avec succès"})
}

func RemoveFromFavorites(w http.ResponseWriter, r *http.Request) {
	var favoriteToRemove models.Favorite
	err := json.NewDecoder(r.Body).Decode(&favoriteToRemove)
	if err != nil {
		http.Error(w, "Erreur lors de la décodification du favori", http.StatusBadRequest)
		return
	}

	favorites, err := utils.ReadFavorites()
	if err != nil {
		http.Error(w, "Erreur lors de la lecture des favoris", http.StatusInternalServerError)
		return
	}

	// Trouver et supprimer le favori
	for i, favorite := range favorites {
		if favorite.Code == favoriteToRemove.Code {
			favorites = append(favorites[:i], favorites[i+1:]...)
			break
		}
	}

	err = utils.WriteFavorites(favorites)
	if err != nil {
		http.Error(w, "Erreur lors de l'écriture des favoris", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Définir le type de contenu de la réponse
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Favori supprimé avec succès"})
}
