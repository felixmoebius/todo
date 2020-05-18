package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("views/*")

	// index
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"index.html",
			gin.H {
				"title": "Index",
			},
		)
	})

	// imprint
	router.GET("/imprint", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"imprint.html",
			nil,
		)
	})

	// new
	router.GET("/new", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"new.html",
			nil,
		)
	})

	router.POST("/t", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	// edit
	router.GET("/t/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.String(http.StatusOK, id)
	})

	router.PUT("t/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.String(http.StatusOK, id)
	})

	// delete
	router.DELETE("t/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.String(http.StatusOK, id)
	})

	router.Run()
}