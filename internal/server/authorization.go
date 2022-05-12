// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package server

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	ErrBadAuthHeader = errors.New("BAD_AUTH_HEADER")
)

func CreateJWTToken(subjectID int64, secret string, d time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(d)),
		Subject:   fmt.Sprint(subjectID),
		ID:        uuid.New().String(),
	}

	return jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(secret))
}

func CreateJWTAccessToken(userID int64, secret string) (string, error) {
	return CreateJWTToken(userID, secret, 1*time.Hour)
}

func CreateJWTRefreshToken(userID int64, secret string) (string, error) {
	return CreateJWTToken(userID, secret, 72*time.Hour)
}

func CreateJWTPairTokens(userID int64, secret string) (string, string, error) {
	accessToken, err := CreateJWTAccessToken(userID, secret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := CreateJWTRefreshToken(userID, secret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func ParseJWTToken(tokenString, secret string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

	if err != nil || !token.Valid {
		return nil, err
	}

	return token.Claims.(*jwt.RegisteredClaims), nil
}

func ExtractJWTToken(authHeader []byte) (string, error) {
	bearerBytes := []byte("Bearer ")
	if !bytes.HasPrefix(authHeader, bearerBytes) {
		return "", ErrBadAuthHeader
	}

	splitted := bytes.SplitN(authHeader, bearerBytes, 2)
	if len(splitted) != 2 {
		return "", ErrBadAuthHeader
	}

	token := splitted[1]

	return string(token), nil
}
