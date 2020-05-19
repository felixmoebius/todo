package main

import (
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	// make database available through context
	router.Use(dbMiddleware())

	// routes
	router.GET(	"/", 		routeIndex)
	router.GET(	"/imprint", 	routeImprint)
	router.GET(	"/new", 	routeNew)
	router.POST(	"t/:id",	routeTodoDelete)
	router.GET(	"/t/:id",	routeTodoGet)
	router.POST(	"/t",		routeTodoPost)
	router.PUT(	"t/:id",	routeTodoPut)

	router.Run(":" + os.Getenv("PORT"))
}