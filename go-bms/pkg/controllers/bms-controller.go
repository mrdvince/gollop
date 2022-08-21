package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mrdvince/gollop/bms/pkg/models"
	"github.com/mrdvince/gollop/bms/pkg/utils"
	"net/http"
	"strconv"
)

func CreateBooks(writer http.ResponseWriter, request *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(request, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	response(writer, res)
}

func response(writer http.ResponseWriter, res []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(res)
}

func GetBooks(writer http.ResponseWriter, _ *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	response(writer, res)
}
func GetBookByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookID := vars["ID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	response(writer, res)

}
func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookID := vars["ID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error parsing book")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	response(writer, res)
}
func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(request, updateBook)
	vars := mux.Vars(request)
	bookID := vars["ID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error occurred while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	response(writer, res)
}
