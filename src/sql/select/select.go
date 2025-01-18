package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"

)

type user struct {
	id int
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, name from users where id < ?", 5)
	defer rows.Close()

	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.name)
		fmt.Println(u)	
	}
	
}