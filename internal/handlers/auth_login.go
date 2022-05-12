// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package handlers

import (
	"api/internal/config"
	"api/internal/models/auth"
	"api/internal/server"
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

const (
	InvalidLoginOrPassword        = "invalid_login_or_password"
	InvalidLoginOrPasswordMessage = "Вы ввели неправильный логин или пароль"

	UserBlocked        = "user_blocked"
	UserBlockedMessage = "Аккаунт заблокирован"
)

func Login(ctx *fasthttp.RequestCtx) {
	requestID := ctx.UserValue("request_id").(string)

	data := new(struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	})

	if err := json.Unmarshal(ctx.Request.Body(), data); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.Response.SetBody(server.NewHTTPBadRequestError(requestID))
		return
	}

	user, err := auth.Login(data.Login, data.Password)
	if err != nil {
		if errors.Is(err, auth.ErrInvalidLoginOrPassword) {
			ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
			ctx.Response.SetBody(server.NewHTTPError(
				InvalidLoginOrPassword,
				InvalidLoginOrPasswordMessage,
				requestID,
			))
			return
		} else if errors.Is(err, auth.ErrUserBlocked) {
			ctx.Response.SetStatusCode(fasthttp.StatusForbidden)
			ctx.Response.SetBody(server.NewHTTPError(
				UserBlocked,
				UserBlockedMessage,
				requestID,
			))
			return
		} else {
			ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.Response.SetBody(server.NewHTTPInternalServerError(requestID))
			return
		}
	}

	config := config.GetConfig()

	accessToken, refreshToken, err := server.CreateJWTPairTokens(user.ID, config.JwtSecret)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.Response.SetBody(server.NewHTTPInternalServerError(requestID))
		return
	}

	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(server.NewHTTPAuthSuccess(accessToken, refreshToken))
}
