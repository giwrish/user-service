package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/giwrish/user-service/internal/repository"
	"github.com/giwrish/user-service/pkg/http/api"
	"golang.org/x/crypto/bcrypt"
)

const (
	UsernameMustNotBeEmpty            = "username must not be empty"
	UserDoesNotExist                  = "user %v does not exist"
	UsernameAndPasswordMustNotBeEmpty = "username and password cannot be empty"
	PasswordCannotBeEmpty             = "password cannot be empty"
	InvalidRequestBody                = "invalid request body, %v"
	InternalServerErr                 = "internal server error, %v"
)

type Users struct {
	mu      sync.Mutex
	queries *repository.Queries
}

func NewUserHandler(queries *repository.Queries) *Users {
	return &Users{
		queries: queries,
	}
}

func (u *Users) GetUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	if username == "" {
		api.Err(w, UsernameMustNotBeEmpty, http.StatusBadRequest)
		return
	}

	user, err := u.queries.GetUser(r.Context(), username)

	if errors.Is(err, sql.ErrNoRows) {
		api.Err(w, fmt.Sprintf(UserDoesNotExist, username), http.StatusNotFound)
		return
	}

	if err != nil {
		api.Err(w, fmt.Sprintf(InternalServerErr, err.Error()), http.StatusInternalServerError)
		return
	}

	api.Success(w, user, http.StatusOK)

}

func (u *Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser repository.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		api.Err(w, fmt.Sprintf(InvalidRequestBody, err.Error()), http.StatusBadRequest)
		return
	}

	if newUser.Username == "" || newUser.Password == "" {
		api.Err(w, UsernameAndPasswordMustNotBeEmpty, http.StatusBadRequest)
		return
	}

	encrypted, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	if err != nil {
		log.Printf("error while encrypting the password, %v", err.Error())
		api.Err(w, fmt.Sprintf(InternalServerErr, err.Error()), http.StatusInternalServerError)
		return
	}

	newUser.Password = string(encrypted)
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	u.mu.Lock()
	user, err := u.queries.CreateUser(r.Context(), newUser)
	u.mu.Unlock()

	if err != nil {
		api.Err(w, fmt.Sprintf(InternalServerErr, err.Error()), http.StatusBadRequest)
		return
	}

	api.Success(w, user, http.StatusCreated)
}

func (u *Users) UpdateUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	if username == "" {
		api.Err(w, UsernameMustNotBeEmpty, http.StatusBadRequest)
		return
	}

	var reqBody struct {
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		api.Err(w, fmt.Sprintf(InvalidRequestBody, err.Error()), http.StatusBadRequest)
		return
	}

	if reqBody.Password == "" {
		api.Err(w, PasswordCannotBeEmpty, http.StatusBadRequest)
		return
	}

	encrypted, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 10)
	if err != nil {
		log.Printf("error while encrypting the password, %v", err.Error())
		api.Err(w, fmt.Sprintf(InternalServerErr, err.Error()), http.StatusInternalServerError)
		return
	}
	_, err = u.queries.GetUser(r.Context(), username)
	if errors.Is(err, sql.ErrNoRows) {
		api.Err(w, fmt.Sprintf(UserDoesNotExist, username), http.StatusNotFound)
		return
	}
	if err != nil {
		api.Err(w, fmt.Sprintf(InternalServerErr, err.Error()), http.StatusInternalServerError)
		return
	}

	userReq := repository.UpdateUserPasswordParams{
		Username:  username,
		Password:  string(encrypted),
		UpdatedAt: time.Now(),
	}

	u.mu.Lock()
	_, err = u.queries.UpdateUserPassword(r.Context(), userReq)
	u.mu.Unlock()

	if err != nil {
		api.Err(w, fmt.Sprintf(InternalServerErr, err.Error()), http.StatusBadRequest)
		return
	}
	api.Success(w, nil, http.StatusNoContent)
}

func (u *Users) DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	if username == "" {
		api.Err(w, UsernameMustNotBeEmpty, http.StatusBadRequest)
		return
	}

	_, err := u.queries.GetUser(r.Context(), username)
	if errors.Is(err, sql.ErrNoRows) {
		api.Err(w, fmt.Sprintf(UserDoesNotExist, username), http.StatusNotFound)
		return
	}
	if err != nil {
		api.Err(w, fmt.Sprintf(InternalServerErr, err.Error()), http.StatusInternalServerError)
		return
	}
	u.mu.Lock()
	_, err = u.queries.DeleteUser(r.Context(), username)
	u.mu.Unlock()
	if err != nil {
		api.Err(w, fmt.Sprintf(InternalServerErr, err.Error()), http.StatusBadRequest)
		return
	}
	api.Success(w, nil, http.StatusNoContent)
}
