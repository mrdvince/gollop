package main

import (
	"github.com/gorilla/mux"
	"github.com/mrdvince/gollop/bms/pkg/routes"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBMSRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9010", router))
}
