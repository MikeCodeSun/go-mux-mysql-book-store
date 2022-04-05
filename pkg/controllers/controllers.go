package controllers

import (
	"encoding/json"
	"example/03-mux-gorm-mysql-book/pkg/models"
	"example/03-mux-gorm-mysql-book/pkg/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	
)



func CreateB(w http.ResponseWriter, r *http.Request) {
	var Book models.Book
	utils.ParseBody(r, &Book)
	b := Book.CreateBook()
	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func GetAllB(w http.ResponseWriter, r *http.Request) {
	b := models.GetAllBooks()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func GetOneB(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookid := vars["bookid"]
	id, err := strconv.ParseInt(bookid, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	book, _ := models.GetOneBook(id)
	b, _ := json.Marshal(book)

w.Header().Set("Content-Type", "Application/json")
w.WriteHeader(http.StatusAccepted)
w.Write(b)
}

func DeleteB(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookid := vars["bookid"]
	id,err := strconv.ParseInt(bookid, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	book := models.DeleteBook(id)
	b, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "Application/json")
w.WriteHeader(http.StatusAccepted)
w.Write(b)
}

func UpdateB(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	vars := mux.Vars(r)
	bookid := vars["bookid"]
	id, err := strconv.ParseInt(bookid, 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	book, db := models.GetOneBook(id)

	utils.ParseBody(r, &newBook)

	if newBook.Name != "" {
		book.Name = newBook.Name
	}
	if newBook.Author != "" {
		book.Author = newBook.Author
	}
	if newBook.Publication != "" {
		book.Publication = newBook.Publication
	}

	db.Save(book)
	b, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "Application/json")
w.WriteHeader(http.StatusAccepted)
w.Write(b)

}