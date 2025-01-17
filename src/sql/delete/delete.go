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

	stmt, _ := db.Prepare("delete from users where id = ?")
	
	resp, _ := stmt.Exec(1)


	id, _ := resp.LastInsertId()
	fmt.Println(id)
	lines, _ := resp.RowsAffected()
	fmt.Println("rows affected",lines)
}