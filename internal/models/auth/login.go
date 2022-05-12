// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package auth

import (
	"api/internal/models"
	"errors"
)

var (
	ErrInvalidLoginOrPassword = errors.New("INVALID_LOGIN_OR_PASSWORD")
	ErrUserBlocked            = errors.New("USER_BLOCKED")
)

func Login(login, password string) (*models.User, error) {
	return nil, errors.New("auth.Login not implemented")
}
