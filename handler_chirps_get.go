package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {

	authorIdStr := r.URL.Query().Get("author_id")
	sortOrder := r.URL.Query().Get("sort")

	chirps := []Chirp{}

	if authorIdStr == "" {
		dbChirps, err := cfg.db.GetChirps(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
			return
		}
		//chirps := []Chirp{}
		for _, dbChirp := range dbChirps {
			chirps = append(chirps, Chirp{
				ID:        dbChirp.ID,
				CreatedAt: dbChirp.CreatedAt,
				UpdatedAt: dbChirp.UpdatedAt,
				UserID:    dbChirp.UserID,
				Body:      dbChirp.Body,
			})
		}

	} else {

		authorID, err := uuid.Parse(authorIdStr)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "couldn't parse author id", err)
		}
		dbChirps, err := cfg.db.GetChirpsByUserId(r.Context(), authorID)
		//chirps := []Chirp{}
		for _, dbChirp := range dbChirps {
			chirps = append(chirps, Chirp{
				ID:        dbChirp.ID,
				CreatedAt: dbChirp.CreatedAt,
				UpdatedAt: dbChirp.UpdatedAt,
				UserID:    dbChirp.UserID,
				Body:      dbChirp.Body,
			})
		}

	}

	if sortOrder == "desc" {
		sort.Slice(chirps, func(i, j int) bool { return chirps[i].CreatedAt.After(chirps[j].CreatedAt) })
	}

	if sortOrder == "asc" || sortOrder == "" {
		sort.Slice(chirps, func(i, j int) bool { return chirps[i].CreatedAt.Before(chirps[j].CreatedAt) })
	}

	respondWithJSON(w, http.StatusOK, chirps)

}

func (cfg *apiConfig) handlerChirpsById(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	dbChirp, err := cfg.db.GetChirpById(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		UserID:    dbChirp.UserID,
		Body:      dbChirp.Body,
	})
}
