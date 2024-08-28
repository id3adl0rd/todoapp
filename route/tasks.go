package route

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
	"to-do-app/logger"
	"to-do-app/models"
	"to-do-app/reason"
	"to-do-app/repository"
)

func CreateTasks(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		logger.Log.Errorf("Error decoding JSON: %s", err)
		RespondWithError(w, http.StatusInternalServerError, reason.ServerProblems)
		return
	}

	defer r.Body.Close()

	_, err = time.Parse(time.RFC3339, task.DueDate)
	if err != nil {
		logger.Log.Errorf("Error converting due date: %s", err)
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid due date: %v", err))
		return
	}

	t := time.Now().Format(time.RFC3339)

	task.CreatedAt = t
	task.UpdatedAt = t

	dx := repository.DB.Create(&task)
	if dx.Error != nil {
		logger.Log.Errorf("Error creating task: %s", dx.Error)
		RespondWithError(w, http.StatusInternalServerError, reason.ServerProblems)
		return
	}

	RespondWithJSON(w, http.StatusCreated, task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Tasks

	dx := repository.DB.Find(&tasks)
	if dx.Error != nil {
		logger.Log.Errorf("Error getting tasks: %s", dx.Error)
		RespondWithError(w, http.StatusInternalServerError, reason.ServerProblems)
		return
	}

	RespondWithJSON(w, http.StatusOK, tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		logger.Log.Errorf("Error converting task id: %s", err)
		RespondWithError(w, http.StatusInternalServerError, reason.ServerProblems)
		return
	}

	if id == 0 {
		logger.Log.Error("Invalid task id. Task ID can not be zero.")
		RespondWithError(w, http.StatusNotFound, reason.TaskNotExist)
		return
	}

	var task models.Tasks
	dx := repository.DB.First(&task, id)
	if dx.Error != nil {
		logger.Log.Errorf("Error getting task: %s", dx.Error)
		RespondWithError(w, http.StatusInternalServerError, reason.TaskNotExist)
		return
	}

	RespondWithJSON(w, http.StatusOK, task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, reason.ServerProblems)
		return
	}

	if id == 0 {
		RespondWithError(w, http.StatusNotFound, reason.TaskNotExist)
		return
	}

	var task models.Tasks
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		logger.Log.Errorf("Error decoding JSON: %s", err)
		RespondWithError(w, http.StatusInternalServerError, reason.ServerProblems)
		return
	}

	defer r.Body.Close()

	_, err = time.Parse(time.RFC3339, task.DueDate)
	if err != nil {
		logger.Log.Errorf("Error converting due date: %s", err)
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid due date: %v", err))
		return
	}

	var dbTask models.Tasks
	dx := repository.DB.First(&dbTask, id)
	if dx.Error != nil {
		logger.Log.Errorf("Error getting task: %s", dx.Error)
		RespondWithError(w, http.StatusInternalServerError, reason.ServerProblems)
		return
	}

	dx = repository.DB.Model(&dbTask).Updates(task)
	if dx.Error != nil {
		logger.Log.Errorf("Error updating task: %s", dx.Error)
		RespondWithError(w, http.StatusInternalServerError, reason.ServerProblems)
		return
	}

	RespondWithJSON(w, http.StatusOK, dbTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		logger.Log.Errorf("Error converting task id: %s", err)
		RespondWithError(w, http.StatusInternalServerError, reason.ServerProblems)
		return
	}

	if id == 0 {
		logger.Log.Error("Invalid task id. Task ID can not be zero.")
		RespondWithError(w, http.StatusNotFound, reason.TaskNotExist)
		return
	}

	dx := repository.DB.First(&models.Tasks{}, id)
	if dx.Error != nil {
		logger.Log.Errorf("Error getting task: %s", dx.Error)
		RespondWithError(w, http.StatusNotFound, reason.TaskNotExist)
		return
	}

	dx = repository.DB.Delete(&models.Tasks{ID: id})
	if dx.Error != nil {
		logger.Log.Errorf("Error deleting task: %s", dx.Error)
		RespondWithError(w, http.StatusInternalServerError, reason.ServerProblems)
		return
	}

	RespondWithJSON(w, http.StatusNoContent, []map[string]string{{"status": reason.TaskDeleted}})
}
