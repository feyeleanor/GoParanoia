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
  var db *sql.DB
  var s *sql.Stmt
  var e error

  db, e = sql.Open("sqlite3", ":memory:")
  ExitOnError(e, DB_OPEN_FAILED)
  defer db.Close()

  s, e = db.Prepare("ATTACH DATABASE ? AS ?")
  ExitOnError(e, DB_PREPARE_FAILED)

  q := os.Args[1] + ".db"
  _, e = s.Exec(q, os.Args[1])
  ExitOnError(e, DB_EXEC_FAILED)
}

func ExitOnError(e error, n int) {
  if e != nil {
    fmt.Println(e)
    os.Exit(n)
  }
}
