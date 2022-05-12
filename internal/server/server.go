// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package server

import (
	"github.com/fasthttp/router"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

type server struct {
	*router.Router
}

func injector(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		requestID := uuid.New().String()

		ctx.Response.Header.Add("Content-Type", "application/json")
		ctx.Response.Header.Add("X-Request-ID", requestID)
		ctx.SetUserValue("request_id", requestID)

		next(ctx)
	}
}

func (s *server) Start(serverAddr string) error {
	return fasthttp.ListenAndServe(serverAddr, injector(s.Handler))
}

func New() *server {
	return &server{
		router.New(),
	}
}
