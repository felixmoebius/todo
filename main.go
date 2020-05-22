package main

import (
	"os"
	"time"
	"html/template"
	"github.com/gin-gonic/gin"
)

const DATE_FORMAT = "2006-01-02"

func main() {
	router := gin.Default()

	db := DBConnect(os.Getenv("DATABASE_URL"))

	// make database available through context
	router.Use(db.Middleware)

	// make date formating available in templates
	router.SetFuncMap(template.FuncMap{
		"formatDate": func (t time.Time) string {
			return t.Format(DATE_FORMAT)
		},
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
