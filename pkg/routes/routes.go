package routes

import (
	"github.com/gorilla/mux"
	"example/03-mux-gorm-mysql-book/pkg/controllers"
)


func BookRoutes(r *mux.Router){
	r.HandleFunc("/", controllers.CreateB).Methods("POST")
	r.HandleFunc("/", controllers.GetAllB).Methods("GET")
	r.HandleFunc("/{bookid}", controllers.GetOneB).Methods("GET")
	r.HandleFunc("/{bookid}", controllers.DeleteB).Methods("DELETE")
	r.HandleFunc("/{bookid}", controllers.UpdateB).Methods("PATCH")
}