## Overview
This is a simple todo list application written in Go.
It uses [Bootstrap](https://getbootstrap.com/) and [Postgresql](https://www.postgresql.org/).

You can try it out [here](https://shrouded-reaches-96469.herokuapp.com).

## Installation
To build the application you need a version of Go that supports [modules](https://blog.golang.org/using-go-modules).
With Go installed run the following commands.

```
$ git clone https://github.com/felixmoebius/todo && cd todo
$ go build
```
## Database
Setup a postgresql server and create a new table as follows.

```sql
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    deadline date,
    description text,
    progress integer
);
```
## Running
Before you run, you need to set the environment variables `PORT` and `DATABASE_URL`.

```
$ export PORT=8080
$ export DATABASE_URL=postgres://user:pass@example.org:port/path
```

Then simply run the application like this.
```
$ ./todo
```

You should now be able to access it at [localhost:8080](http://localhost:8080).
