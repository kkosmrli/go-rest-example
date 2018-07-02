package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/kkosmrli/rest-example/daos"
	. "github.com/kkosmrli/rest-example/models"
)

var dao MoviesDAO

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Function not yet implemented...")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var movie Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	movie.ID = bson.NewObjectId()

	if err := dao.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Function not yet implemented...")
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Function not yet implemented...")
}

func FindMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Function not yet implemented...")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovieEndPoint).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
