package api

import (
	unrolledRender "github.com/unrolled/render"
)

var render = unrolledRender.New(
	unrolledRender.Options{
		IndentJSON: true,
	},
)

type ApiError struct {
	Err string `json:"error"`
}
