package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Id      int
	Model   string
	Company string
	Price   int
}

var database *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := database.Query("select * from products")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	products := []Product{}

	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, products)
}

func main() {

	db, err := sql.Open("mysql", "root:@/productdb")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	http.HandleFunc("/", IndexHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
