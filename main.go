package main

import (
	"example/03-mux-gorm-mysql-book/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.BookRoutes(r)
	http.Handle("/", r)
	fmt.Println("go")
	
	log.Fatal(http.ListenAndServe(":8080", r))
}