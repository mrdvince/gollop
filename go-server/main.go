package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(writer, "method not supported", http.StatusNotFound)
		return
	}
	_, err := fmt.Fprintf(writer, "hello!")
	if err != nil {
		return
	}
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "Parse form error: %v", err)
		return
	}
	_, _ = fmt.Fprintf(writer, "POST request successful")
	name := request.FormValue("name")
	address := request.FormValue("address")

	_, _ = fmt.Fprintf(writer, "Name = %s\n", name)
	_, _ = fmt.Fprintf(writer, "Adress = %s\n", address)
}
