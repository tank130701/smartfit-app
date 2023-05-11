package postrgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	usersTable        = "Users"
	//articlesInfoTable = "articles"
	sessionsTable     = "sessions"
)


func NewPostrgesConnection(host string, username string, password string, port int, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
