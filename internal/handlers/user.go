package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giwrish/user-service/internal/models"
	"github.com/giwrish/user-service/internal/repository"
	"github.com/giwrish/user-service/pkg/http/api"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	if username == "" {
		api.Err(w, "username must not be empty", http.StatusBadRequest)
		return
	}

	user, err := repository.FindByUserName(username)

	if err != nil {
		api.Err(w, "user does not exist", http.StatusNotFound)
		return
	}

	api.Success(w, user, http.StatusOK)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		api.Err(w, fmt.Sprintf("invalid request body %v", err.Error()), http.StatusBadRequest)
		return
	}

	if newUser.Username == "" || newUser.Password == "" {
		api.Err(w, "username and password cannot be empty", http.StatusBadRequest)
		return
	}

	// encrypt password
	repository.CreateUser(&newUser)
	api.Success(w, newUser, http.StatusCreated)
}
