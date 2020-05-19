package main

import (
	"net/http"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "database/sql"
	_ "github.com/lib/pq"
)

func routeIndex(ctx *gin.Context) {
	db := dbGet(ctx)

	rows, err := db.Query("SELECT id, description, deadline, progress FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	todos := make([]todo, 0)
	for rows.Next() {
		var t todo
		if err := rows.Scan(&t.Id, &t.Description, &t.Deadline, &t.Progress); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, t)
	}

	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H {
			"todos": todos,
		},
	)
}

func routeImprint(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"imprint.html",
		nil,
	)
}

func routeNew(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"new.html",
		nil,
	)
}

// todo resource
// create
func routeTodoPost(ctx *gin.Context) {
	db := dbGet(ctx)

	desc := ctx.PostForm("description")
	date := ctx.PostForm("deadline")
	prog := ctx.PostForm("progress")

	_, err := db.Exec("INSERT INTO todos(description, deadline, progress) VALUES ($1, $2, $3)", desc, date, prog)

	if err != nil {
		panic(err)
	}

	ctx.Redirect(http.StatusMovedPermanently, "/")
}

// read
func routeTodoGet(ctx *gin.Context) {
	db := dbGet(ctx)

	var t todo

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusNotFound, "404")
		return
	}

	row := db.QueryRow("SELECT id, description, deadline, progress FROM todos WHERE id = $1", id)
	err = row.Scan(&t.Id, &t.Description, &t.Deadline, &t.Progress)
	if err != nil {
		ctx.String(http.StatusNotFound, "404")
		return
	}

	ctx.HTML(
		http.StatusOK,
		"edit.html",
		gin.H {
			"todo": t,
		},
	)
}

// update
func routeTodoPut(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, id)
}

// delete
func routeTodoDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, id)
}
