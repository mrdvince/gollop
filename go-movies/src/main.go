package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Movie struct {
	ID       string    `json:"ID,"`
	ISBN     string    `json:"ISBN"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

var movies []Movie

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{ID}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{ID}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{ID}", deleteMovie).Methods("DELETE")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for idx, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			break
		}
	}
	err := json.NewEncoder(writer).Encode(movies)
	if err != nil {
		return
	}
}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for idx, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			var movie Movie
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = params["ID"]
			movies = append(movies, movie)
			json.NewEncoder(writer).Encode(movie)
			return
		}
	}
}

func createMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movie)
}

func getMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range movies {
		if item.ID == params["ID"] {
			json.NewEncoder(writer).Encode(item)
			return

		}
	}
}

func getMovies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(movies)
	if err != nil {
		return
	}
}
