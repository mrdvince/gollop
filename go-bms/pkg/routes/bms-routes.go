package routes

import (
	"github.com/gorilla/mux"
	"github.com/mrdvince/gollop/bms/pkg/controllers"
)

var RegisterBMSRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.CreateBooks).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{ID}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{ID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{ID}", controllers.DeleteBook).Methods("DELETE")
}
