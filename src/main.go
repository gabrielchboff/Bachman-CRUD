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

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	First    string `json:"firstname"`
	Lastname string `jason:"lastname"`
}

var movies []Movie


func getMovies(w http.ResponseWriter, r http.Request)  {
					
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)

}


func main() {
	r := mux.NewRouter()
  
  movies = append(movies, Movie{
  	ID:"1", 
  	Isbn:"438227", 
  	Title:"Movie One", 
  	Director: &Director{
  		Firstname: "Jhon", Lastname: "Doe"
  	},
  } )
  movies = append(movies, Movie{
  	ID:"2", 
  	Isbn: "73565", 
  	Title: "Movie Two",
  	Director: &Director{
  		First: "Steve", 
  		Lastname:"Smith"
  	},
  } )
  
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("# Starting Server at port 8000 [* * *]")
	log.Fatal(http.ListenAndServe(":8000", r))

}
