package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //takes variable like {id} from the url request
	for i, v := range movies {
		if v.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, v := range movies {
		if v.ID == params["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = params["id"]
	for i, v := range movies {
		if v.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	movies = append(movies, movie)
}

func main() {
	movie1 := Movie{ID: "1", Isbn: "456789", Title: "Movie 1", Director: &Director{Firstname: "Tom", Lastname: "Nguyen"}}
	movie2 := Movie{ID: "2", Isbn: "123456", Title: "Movie 2", Director: &Director{Firstname: "Hank", Lastname: "David"}}
	movies = append(movies, movie1, movie2)
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("Listen to port :8080")
    if err:=http.ListenAndServe(":8080", nil);err!=nil{
		log.Fatal(err)
		fmt.Printf("Catched error: %v",err)
	}
}
