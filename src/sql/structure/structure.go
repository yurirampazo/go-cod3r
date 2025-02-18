package main

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists users")
	exec(db, `create table if not exists users(
			 id integer auto_increment primary key,
			 name varchar(80)
			 )`)
}