package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"to-do-app/logger"
)

type Tasks struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func CreateTasks(w http.ResponseWriter, r *http.Request) {
	var task Tasks
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		logger.Log.Errorf("Error decoding JSON: %s", err)
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
	}

	defer r.Body.Close()
	respondWithJSON(w, http.StatusCreated, task)
}
