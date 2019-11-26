package main

import "crypto/rand"
import "database/sql"
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
	ENCRYPTION_FAILED
)

func main() {
	k := os.Getenv("AES_KEY")

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

		AddAccount(k, s, "Alpha", "a@someserver.com")
		AddAccount(k, s, "Beta", "b@someserver.com")
		AddAccount(k, s, "Gamma", "g@someserver.com")

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

func AddAccount(k string, s *sql.Stmt, p ...interface{}) {
	var e error

	t := append([]interface{}{RandomToken(32)}, EncryptFields(k, p...)...)
	_, e = s.Exec(t...)
	if e != nil {
		t[0] = RandomToken(32)
		_, e = s.Exec(t...)
	}
	ExitOnError(e, DB_INSERT_FAILED)
}

func EncryptFields(k string, p ...interface{}) (r []interface{}) {
	for _, v := range p {
		b, e := AESEncrypt(k, fmt.Sprint(v))
		ExitOnError(e, ENCRYPTION_FAILED)
		r = append(r, EncodeToString(b))
	}
	return
}

func RandomToken(n int) (s string) {
	b := make([]byte, n)
	_, e := rand.Read(b)
	ExitOnError(e, MAKE_TOKEN_FAILED)
	return EncodeToString(b)
}

func OpenDB(n string, f func(*sql.DB)) {
	db, e := sql.Open("sqlite3", n)
	ExitOnError(e, DB_OPEN_FAILED)
	defer db.Close()
	f(db)
}
