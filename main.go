package main

import (
	product3 "Catlog/handler/product"
	product2 "Catlog/service/product"
	"Catlog/store/brand"
	"Catlog/store/product"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "sunny:Sunny@9570@(127.0.0.1)/store")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in connection establishment!")
		return
	}
	p := product.New(db)
	b := brand.New(db)
	s := product2.New(p,b)
	ht := product3.Handler{s}
	fmt.Println("Server starting....")
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/product/{id}", ht.ProductWithId)
	r.HandleFunc("/product", ht.InsertProduct).Methods("POST")
	http.ListenAndServe(":8080",r)
}
