package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/geek?charset=utf8mb4")
}

type Result struct {
	Name string
}

func Query() (Result, error) {
	var res Result
	err := db.QueryRow("select name from week02 where id = 1").Scan(&res.Name)
	if err != nil {
		return res, errors.Wrap(err, "scan error")
	}
	return res, nil
}
