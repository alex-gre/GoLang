package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// обновляем строку с id=1
	result, err := db.Exec("update Products set price = $1 where id = $2", 69000, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество обновленных строк
}
