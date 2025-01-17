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

	stmt, _ := db.Prepare("update users set name = ? where id = ?")
	
	resp, _ := stmt.Exec("CLEIDSON", 1)
	resp2, _ := stmt.Exec("Hector", 2)

	var rows int64 = 0 

	id, _ := resp.LastInsertId()
	fmt.Println(id)

	lines, _ := resp.RowsAffected()
	rows += lines
	lines2, _ := resp2.RowsAffected()
	rows += lines2	
	fmt.Println("rows affected",rows)
}