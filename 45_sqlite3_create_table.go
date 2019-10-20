package main

import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "fmt"
import "os"

const (
	_ = iota
	DB_OPEN_FAILED
	DB_PREPARE_FAILED
	DB_QUERY_FAILED
	DB_INSERT_FAILED
)

func main() {
	OpenDB(os.Args[1], func(db *sql.DB) {
		var s *sql.Stmt
		var r *sql.Row
		var i int64

		_, e := db.Exec(`DROP TABLE IF EXISTS Account`)
		ExitOnError(e, DB_QUERY_FAILED)

		_, e = db.Exec(`
      CREATE TABLE Account (
        id    VARCHAR PRIMARY KEY,
        Name  TEXT NOT NULL,
        Email TEXT UNIQUE NOT NULL
      ) WITHOUT ROWID`)
		ExitOnError(e, DB_QUERY_FAILED)

		s, e = db.Prepare("INSERT INTO Account(id, Name, Email) VALUES(?, ?, ?)")
		ExitOnError(e, DB_PREPARE_FAILED)

		_, e = s.Exec("a", "Ellie", "a@someserver.com")
		ExitOnError(e, DB_INSERT_FAILED)

		_, e = s.Exec("b", "Ellie", "b@someserver.com")
		ExitOnError(e, DB_INSERT_FAILED)

		_, e = db.Exec(
			"INSERT INTO Account(id, Name, Email) VALUES(?, ?, ?)",
			"c",
			"Ellie",
			"c@someserver.com")
		ExitOnError(e, DB_INSERT_FAILED)

		r = db.QueryRow("SELECT count(*) FROM Account")
		ExitOnError(e, DB_QUERY_FAILED)

		r.Scan(&i)
		fmt.Println("rows in Account table =", i)
	})
}

func OpenDB(n string, f func(*sql.DB)) {
	db, e := sql.Open("sqlite3", n)
	ExitOnError(e, DB_OPEN_FAILED)
	defer db.Close()
	f(db)
}
