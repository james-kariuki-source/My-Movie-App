package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(rw http.ResponseWriter, r *http.Request){
	rw.Header().Set("Content-Type","application/json")
	json.NewEncoder(rw).Encode(movies)
}

func deleteMovie(rw http.ResponseWriter, r *http.Request){
	rw.Header().Set("Content-Type","application/json")

	param := mux.Vars(r)

	for index, items := range movies{
		if items.ID == param["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(rw).Encode(movies)
}

func getMovie(rw http.ResponseWriter, r *http.Request){
	rw.Header().Set("Content-Type","application/json")

	param := mux.Vars(r)

	for _, items := range movies{
		if items.ID == param["id"]{
			json.NewEncoder(rw).Encode(items)
			return
		}
	}
}

func createMovie(rw http.ResponseWriter, r *http.Request){
	rw.Header().Set("Content-Type","application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	json.NewEncoder(rw).Encode(movie)
}

func updateMovie(rw http.ResponseWriter, r *http.Request){
	rw.Header().Set("Content-Type","application/json")

	param := mux.Vars(r)

	for index, items := range movies{
		if items.ID == param["id"]{

			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = param["id"]
			movies = append(movies, movie)

			json.NewEncoder(rw).Encode(movie)
		}
	}
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "2343534", Title: "Drag Me To Hell", Director: &Director{Firstname: "Sam",Lastname: "Raimi"}})

	movies = append(movies, Movie{ID: "2", Isbn: "94646458", Title: "The Matrix", Director: &Director{Firstname: "Wachowski",Lastname: "Brothers"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}