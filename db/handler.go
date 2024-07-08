package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

type Handler struct {
	M  sync.Mutex
	DB *sqlx.DB
}

var H Handler

func SetHandler(DSN string) error {
	H.M.Lock()
	defer H.M.Unlock()
	db, err := sqlx.Connect("sqlite3", DSN)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	H.DB = db
	return nil
}

func Exists[K any](filter, col, table string) (bool, error) {
	H.M.Lock()
	defer H.M.Unlock()
	var data []K
	err := H.DB.Select(&data, "select * from $1 where $2=$3", table, col, filter)
	if err != nil {
		return false, err
	}
	return len(data) > 0, err
}
