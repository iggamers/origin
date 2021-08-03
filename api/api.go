package api

import (
	"net/http"

	"github.com/i9code/xcho"
)

type api struct {
}

func Mount(g *xcho.Group, mfs ...xcho.MiddlewareFunc) {
	g.GET("", func(ctx *xcho.Context) (err error) {

		return ctx.JSON(http.StatusOK, "abc")
	}).Name = "测试"
}
