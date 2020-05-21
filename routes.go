package main

import (
	"net/http"
	"strconv"
	"log"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	_ "database/sql"
	_ "github.com/lib/pq"
)

func routeIndex(ctx *gin.Context) {
	db := dbInstance(ctx)

	todos, err := db.All()
	if err != nil {
		log.Fatal(err)
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
func routeTodoCreate(ctx *gin.Context) {
	db := dbInstance(ctx)

	t, err := todoFromContext(ctx, true)

	err = db.Insert(t)
	if err != nil {
		panic(err)
	}

	ctx.Redirect(http.StatusMovedPermanently, "/")
}

// read
func routeTodoRead(ctx *gin.Context) {
	db := dbInstance(ctx)

	var t todo

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusNotFound, "404")
		return
	}

	t, err = db.Get(id)

	ctx.HTML(
		http.StatusOK,
		"edit.html",
		gin.H {
			"todo": t,
		},
	)
}

// update
func routeTodoUpdate(ctx *gin.Context) {
	db := dbInstance(ctx)

	t, err := todoFromContext(ctx, false)
	if err != nil {
		ctx.String(http.StatusNotFound, "404")
		return
	}
	
	err = db.Update(t)
	if err != nil {
		panic(err)
	}

	ctx.Redirect(http.StatusMovedPermanently, "/")
}

// delete
func routeTodoDelete(ctx *gin.Context) {
	db := dbInstance(ctx)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusNotFound, "404")
		return
	}

	err = db.Delete(id)
	if err != nil {
		panic(err)
	}

	ctx.Redirect(http.StatusMovedPermanently, "/")
}

func todoFromContext(ctx *gin.Context, noId bool) (todo, error)  {
	var t todo
	var err error

	id := "0"
	if !noId {
		id = ctx.Param("id")
	}
	
	desc := ctx.PostForm("description")
	date := ctx.PostForm("deadline")
	prog := ctx.PostForm("progress")

	// validate input
	if t.Id, err = strconv.Atoi(id); err != nil {
		return t, err
	}

	if len(desc) > 160 {
		return t, errors.New("Description longer than 160 characters")
	}
	t.Description = desc

	if t.Deadline, err = time.Parse("2006-01-02", date); err != nil {
		return t, err
	}

	if t.Progress, err = strconv.Atoi(prog); err != nil {
		return t, err
	}

	return t, nil
}
