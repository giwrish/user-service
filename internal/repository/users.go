package repository

import (
	"errors"
	"fmt"

	"github.com/giwrish/user-service/internal/models"
)

var users = make(map[string]*models.User)

func FindByUserName(username string) (*models.User, error) {

	user, ok := users[username]

	if !ok {
		return nil, errors.New(fmt.Sprintf("User %v does not exist", username))
	}

	return user, nil
}

func CreateUser(user *models.User) {
	users[user.Username] = user
}
