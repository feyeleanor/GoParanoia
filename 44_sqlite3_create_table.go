package main

import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "fmt"
import "os"

const (
  _ = iota
  DB_OPEN_FAILED
  DB_PREPARE_FAILED
  DB_EXEC_FAILED
)

func main() {
  d := os.Args[1]
  os.Remove(d)

  OpenDB(d, func(db *sql.DB) {
    var s *sql.Stmt
    var e error

    s, e = db.Prepare(`
      CREATE TABLE Account (
        id    INTEGER PRIMARY KEY,
        Name  TEXT NOT NULL,
        Email TEXT UNIQUE NOT NULL
      )`)
    ExitOnError(e, DB_PREPARE_FAILED)

    _, e = s.Exec()
    ExitOnError(e, DB_EXEC_FAILED)
  })
}

func OpenDB(n string, f func(*sql.DB)) {
  db, e := sql.Open("sqlite3", n)
  ExitOnError(e, DB_OPEN_FAILED)
  defer db.Close()
  f(db)
}

func ExitOnError(e error, n int) {
  if e != nil {
    fmt.Println(e)
    os.Exit(n)
  }
}
