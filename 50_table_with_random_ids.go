package main

import "crypto/rand"
import "database/sql"
import "encoding/base64"
import _ "github.com/mattn/go-sqlite3"
import "fmt"
import "os"
import "strings"

const (
  _ = iota
  DB_OPEN_FAILED
  DB_PREPARE_FAILED
  DB_QUERY_FAILED
  DB_INSERT_FAILED
  DB_CREATE_INDEX
  MAKE_TOKEN_FAILED
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

    CreateIndex(db, "PrimKey", "id")
    CreateIndex(db, "NameIndex", "Name")
    CreateIndex(db, "EmailIndex", "Email")

    s, e = db.Prepare("INSERT INTO Account(id, Name, Email) VALUES(?, ?, ?)")
    ExitOnError(e, DB_PREPARE_FAILED)

    AddAccount(s, "Alpha", "a@someserver.com")
    AddAccount(s, "Beta", "b@someserver.com")
    AddAccount(s, "Gamma", "g@someserver.com")

    r = db.QueryRow("SELECT count(*) FROM Account")
    ExitOnError(e, DB_QUERY_FAILED)

    r.Scan(&i)
    fmt.Println("rows in Account table =", i)
  })
}

func CreateIndex(db *sql.DB, n string, f ...string) {
  q := strings.Join(f, ",")
  _, e := db.Exec(
    fmt.Sprintf("CREATE UNIQUE INDEX %s ON Account (%s)", n, q))
  ExitOnError(e, DB_CREATE_INDEX)
}

func AddAccount(s *sql.Stmt, p ...interface{}) {
  var e error

  t := RandomToken(16)
  _, e = s.Exec(PushValue(t, p...)...)
  if e != nil {
    t = RandomToken(16)
    _, e = s.Exec(PushValue(t, p...)...)
  }
  ExitOnError(e, DB_INSERT_FAILED)
}

func PushValue(v interface{}, p ...interface{}) []interface{} {
  return append([]interface{}{v}, p...)
}

func RandomToken(n int) (s string) {
  b := make([]byte, n)
  _, e := rand.Read(b)
  ExitOnError(e, MAKE_TOKEN_FAILED)
  return base64.StdEncoding.EncodeToString(b)
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
