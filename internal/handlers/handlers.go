package handlers

import (
	"encoding/json"
	"fmt"
	"genshin/internal/parser"
	"genshin/internal/repository"
	"net/http"
)

type Handlers struct {
	repo   repository.Repo
	parser *parser.Parse
}

func NewHandler(repo repository.Repo, parse *parser.Parse) *Handlers {
	return &Handlers{
		repo:   repo,
		parser: parse,
	}
}

func (h *Handlers) InsertCharacterHandler(w http.ResponseWriter, r *http.Request) {
	characters, _ := h.parser.FetchData()

	for _, character := range characters {
		if err := h.repo.InsertCharacter(character); err != nil {
			http.Error(w, fmt.Sprintf("Error inserting character: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Characters successfully inserted")
}

func (h *Handlers) GetCharacters(w http.ResponseWriter, r *http.Request) {
	characters, err := h.repo.GetCharacters()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error get character: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	if err := json.NewEncoder(w).Encode(characters); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
