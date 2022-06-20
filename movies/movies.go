package movies

import (
	"encoding/json"
	"io"
	"log"
	"movies-crud-api/director"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string             `json: "id"`
	Isbn     string             `json: "isbn"`
	Title    string             `json: "title"`
	Director *director.Director `json: "director"`
}

var Movies []Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Default().Println("URL:", r.URL.Path)
	log.Default().Println("HEADER:", r.Header)
	io.WriteString(w, "List of Movies\n")
	json.NewEncoder(w).Encode(Movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range Movies {
		if item.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Movies)
	return
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range Movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	log.Println("New Movie:", movie)
	Movies = append(Movies, movie)
	json.NewEncoder(w).Encode(movie)
	return
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Movies {
		if item.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
			break
		}
	}
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	log.Println("Update Movie:", movie)
	Movies = append(Movies, movie)
	json.NewEncoder(w).Encode(Movies)
	return
}
