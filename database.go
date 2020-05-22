package main

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
)

// embed sql.DB to define receiver functions
type DB struct {
	*sql.DB
}

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

	CONTEXT_DB_KEY = "database"
)

func DBConnect(url string) *DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic("can't connect to database")
	}

	return &DB{db}
}

func (db *DB) Middleware(ctx *gin.Context) {
	ctx.Set(CONTEXT_DB_KEY, db)
}

func (db *DB) Get(id int) (todo, error) {
	var t todo

	row := db.QueryRow(SQL_QUERY_SELECT_ID, id)
	err := row.Scan(
		&t.Id, 
		&t.Description, 
		&t.Deadline, 
		&t.Progress)

	return t, err
}

func (db *DB) Delete(id int) error {
	_, err := db.Exec(
		SQL_QUERY_DELETE_ID,
		id)

	return err
}

func (db *DB) Update(t todo) error {
	_, err := db.Exec(
		SQL_QUERY_UPDATE_ID, 
		t.Description, 
		t.Deadline, 
		t.Progress, 
		t.Id)

	return err
}

func (db *DB) Insert(t todo) error {
	_, err := db.Exec(
		SQL_QUERY_INSERT,
		t.Description,
		t.Deadline,
		t.Progress)

	return err
}

func (db *DB) All() ([]todo, error) {
	todos := make([]todo, 0)

	rows, err := db.Query(SQL_QUERY_SELECT_ALL)
	if err != nil {
		return todos, err
	} else {
		defer rows.Close()
	}

	// loop over results and convert to todo	
	for rows.Next() {
		var t todo

		if err = rows.Scan(
			&t.Id, 
			&t.Description, 
			&t.Deadline, 
			&t.Progress,
		); err != nil {
			break
		}

		todos = append(todos, t)
	}

	return todos, err
}
