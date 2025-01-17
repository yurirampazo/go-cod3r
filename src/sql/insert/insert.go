package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"

)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into users(name) values (?)")
	stmt.Exec("Maria")
	stmt.Exec("John")
	stmt.Exec("Yuri")

	resp, _ := stmt.Exec("Jana")

	id, _ := resp.LastInsertId()
	fmt.Println(id)

	lines, _ := resp.RowsAffected()
	fmt.Println(lines)
}