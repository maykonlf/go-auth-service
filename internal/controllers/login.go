package controllers

import (
	"errors"
	"github.com/maykonlf/go-auth-service/internal/entities"
)

type LoginControllerI interface {
	Authenticate(user *entities.User) error
}

func NewLoginController() LoginControllerI {
	return &LoginController{}
}

type LoginController struct {

}

func (l *LoginController) Authenticate(user *entities.User) error {
	if user.Username != "username" && user.Password != "password" {
		return errors.New("invalid credentials")
	}

	return nil
}
