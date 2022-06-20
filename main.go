package main

import (
	"log"
	"movies-crud-api/director"
	"movies-crud-api/movies"
	"net/http"

	"github.com/gorilla/mux"
	// log "github.com/sirupsen/logrus"
)

func main() {
	r := mux.NewRouter()
	// moviesList := movies.Movies{}
	movies.Movies = append(movies.Movies, []movies.Movie{
		{
			ID:    "1",
			Isbn:  "12341",
			Title: "Gravity",
			Director: &director.Director{
				FirstName: "Alfonso",
				LastName:  "Cuarón",
			},
		},
		{
			ID:    "2",
			Isbn:  "12342",
			Title: "Hold the Dark",
			Director: &director.Director{
				FirstName: "Jeremy",
				LastName:  "Saulnier",
			},
		},
	}...)
	r.HandleFunc("/movies", movies.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", movies.GetMovie).Methods("GET")
	r.HandleFunc("/movies", movies.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", movies.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", movies.DeleteMovie).Methods("DELETE")
	log.Printf("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// zahid@Zahid-urrehman Movies-CRUD-API % curl localhost:8000/movies
// List of Movies
// [{"ID":"1","Isbn":"12341","Title":"Gravity","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}},{"ID":"2","Isbn":"12342","Title":"Hold the Dark","Director":{"FirstName":"Jeremy","LastName":"Saulnier"}}]
// zahid@Zahid-urrehman Movies-CRUD-API % curl localhost:8000/movies
// List of Movies
// [{"ID":"1","Isbn":"12341","Title":"Gravity","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}},{"ID":"2","Isbn":"12342","Title":"Hold the Dark","Director":{"FirstName":"Jeremy","LastName":"Saulnier"}}]
// zahid@Zahid-urrehman Movies-CRUD-API % curl localhost:8000/movies/1
// {"ID":"1","Isbn":"12341","Title":"Gravity","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}}
// zahid@Zahid-urrehman Movies-CRUD-API % curl localhost:8000/movies/1
// {"ID":"1","Isbn":"12341","Title":"Gravity","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}}
// zahid@Zahid-urrehman Movies-CRUD-API % curl -X POST -H "Content-Type: application/json" -d '{"ID":"3","Isbn":"123415","Title":"Gravity1","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}}' localhost:8000/movies
// {"ID":"3","Isbn":"123415","Title":"Gravity1","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}}
// zahid@Zahid-urrehman Movies-CRUD-API % curl localhost:8000/movies
// List of Movies
// [{"ID":"1","Isbn":"12341","Title":"Gravity","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}},{"ID":"2","Isbn":"12342","Title":"Hold the Dark","Director":{"FirstName":"Jeremy","LastName":"Saulnier"}},{"ID":"3","Isbn":"123415","Title":"Gravity1","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}}]
// zahid@Zahid-urrehman Movies-CRUD-API % curl -X PUT -H "Content-Type: application/json" -d '{"ID":"3","Isbn":"123415","Title":"Gravity2","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}}' localhost:8000/movies
// zahid@Zahid-urrehman Movies-CRUD-API % curl -X PUT -H "Content-Type: application/json" -d '{"ID":"3","Isbn":"123415","Title":"Gravity2","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}}' localhost:8000/movies/3
// zahid@Zahid-urrehman Movies-CRUD-API % curl localhost:8000/movies
// List of Movies
// [{"ID":"1","Isbn":"12341","Title":"Gravity","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}},{"ID":"2","Isbn":"12342","Title":"Hold the Dark","Director":{"FirstName":"Jeremy","LastName":"Saulnier"}},{"ID":"3","Isbn":"123415","Title":"Gravity1","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}}]
// zahid@Zahid-urrehman Movies-CRUD-API % curl -X DELETE localhost:8000/movies/3
// [{"ID":"1","Isbn":"12341","Title":"Gravity","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}},{"ID":"2","Isbn":"12342","Title":"Hold the Dark","Director":{"FirstName":"Jeremy","LastName":"Saulnier"}}]
// zahid@Zahid-urrehman Movies-CRUD-API % curl localhost:8000/movies
// List of Movies
// [{"ID":"1","Isbn":"12341","Title":"Gravity","Director":{"FirstName":"Alfonso","LastName":"Cuarón"}},{"ID":"2","Isbn":"12342","Title":"Hold the Dark","Director":{"FirstName":"Jeremy","LastName":"Saulnier"}}]
// zahid@Zahid-urrehman Movies-CRUD-API %
