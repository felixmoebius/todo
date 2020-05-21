package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	SQL_QUERY_SELECT_ID = `
		SELECT id, description, deadline, progress 
		FROM todos 
		WHERE id = $1`

	SQL_QUERY_SELECT_ALL = `
		SELECT id, description, deadline, progress 
		FROM todos`

	SQL_QUERY_DELETE_ID = `
		DELETE FROM todos 
		WHERE id = $1`

	SQL_QUERY_UPDATE_ID = `
		UPDATE todos 
		SET description = $1, deadline = $2, progress = $3 
		WHERE id = $4`

	SQL_QUERY_INSERT = `
		INSERT 
		INTO todos(description, deadline, progress) 
		VALUES ($1, $2, $3)`
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

func dbInstance(ctx *gin.Context) *sql.DB {
	return ctx.MustGet("database").(*sql.DB)
}

func dbGet(db *sql.DB, id int) (todo, error) {
	var t todo

	row := db.QueryRow(SQL_QUERY_SELECT_ID, id)
	err := row.Scan(&t.Id, &t.Description, &t.Deadline, &t.Progress)
	if err != nil {
		return t, err
	}

	return t, nil
}

func dbGetAll(db *sql.DB) ([]todo, error) {
	todos := make([]todo, 0)

	rows, err := db.Query(SQL_QUERY_SELECT_ALL)
	if err != nil {
		return todos, err
	} else {
		defer rows.Close()
	}

	for rows.Next() {
		var t todo

		err = rows.Scan(&t.Id, &t.Description, &t.Deadline, &t.Progress)
		if err != nil {
			return todos, err
		}

		todos = append(todos, t)
	}

	return todos, nil
}

func dbDelete(db *sql.DB, id int) error {
	_, err := db.Exec(SQL_QUERY_DELETE_ID, id)
	return err
}

func dbUpdate(db *sql.DB, t todo) error {
	_, err := db.Exec(SQL_QUERY_UPDATE_ID, t.Description, t.Deadline, t.Progress, t.Id)
	return err
}

func dbInsert(db *sql.DB, t todo) error {
	_, err := db.Exec(SQL_QUERY_INSERT, t.Description, t.Deadline, t.Progress)
	return err
}
