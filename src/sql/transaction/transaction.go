package main

import(
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err) // ends excecution of program
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into users(id, name) values (?, ?)")

	_ , err = stmt.Exec(999, "Wilton")
	if err != nil {
		tx.Rollback()
		fmt.Println("Error in transaction, intiating rollback")
		log.Fatal(err)
	}
	tx.Commit()

	_, err = stmt.Exec(1000, "Katia")
	if err != nil {
		tx.Rollback()
		fmt.Println("Error in transaction, intiating rollback")
		log.Fatal(err)
	}

	_, err = stmt.Exec(1, "TesteErro")

	if err != nil {
		tx.Rollback()
		fmt.Println("Error in transaction, intiating rollback")
		log.Fatal(err)
	}
	tx.Commit()
	fmt.Println("Successfully commited transaction")
}