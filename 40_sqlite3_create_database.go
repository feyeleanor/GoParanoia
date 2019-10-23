package main

import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "os"

const (
	_ = iota
	DB_OPEN_FAILED
	DB_PREPARE_FAILED
	DB_EXEC_FAILED
)

func main() {
	db, e := sql.Open("sqlite3", os.Args[1])
	ExitOnError(e, DB_OPEN_FAILED)
	defer db.Close()

	_, e = db.Prepare("")
	ExitOnError(e, DB_PREPARE_FAILED)
}
