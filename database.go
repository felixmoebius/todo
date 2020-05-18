package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
)

func dbSetup() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("can't connect to database")
	}

	return db
}

func dbMiddleware() gin.HandlerFunc {
	db := dbSetup()
	return func (ctx *gin.Context) {
		ctx.Set("database", db)
	}
}

func dbGet(ctx *gin.Context) *sql.DB {
	return ctx.MustGet("database").(*sql.DB)
}