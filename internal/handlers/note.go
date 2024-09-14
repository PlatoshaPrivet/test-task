package handlers

import (
	"encoding/json"
	"net/http"
	"note_API/internal/models"
	"note_API/internal/utils"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	var note models.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !utils.ValidateSpelling(note.Content) {
		http.Error(w, "Spelling error", http.StatusBadRequest)
		return
	}

	models.NoteID++
	note.ID = models.NoteID
	note.UserID = userID
	models.Notes = append(models.Notes, note)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	userNotes := []models.Note{}

	for _, note := range models.Notes {
		if note.UserID == userID {
			userNotes = append(userNotes, note)
		}
	}

	json.NewEncoder(w).Encode(userNotes)
}
