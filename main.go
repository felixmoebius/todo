package main

import (
	"os"
	"html/template"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// make database available through context
	router.Use(dbMiddleware())

	// make date formating available in templates
	router.SetFuncMap(template.FuncMap{
		"formatDate": formatDate,
	})

	router.LoadHTMLGlob("views/*")

	// routes
	router.GET(	"/",			routeIndex)
	router.GET(	"/imprint",		routeImprint)
	router.GET(	"/new",			routeNew)
	router.POST(	"/t",			routeTodoCreate)
	router.GET(	"/t/:id",		routeTodoRead)
	router.POST(	"/t/:id/update",	routeTodoUpdate)
	router.POST(	"/t/:id/delete",	routeTodoDelete)

	router.Run(":" + os.Getenv("PORT"))
}