package config

import (
	"database/sql"
	"fmt"
	_ `github.com/go-sql-driver/mysql`
)

const (
	username = `root`
	password = ``
	database = `akademik`
)

var dsn = fmt.Sprintf(`%s:%s@/%s`,username,password,database)

func Mysql() (db *sql.DB,err error) {
	db, err = sql.Open(`mysql`,dsn)
	if err != nil {
		fmt.Println(err)
	}
	return 
}
