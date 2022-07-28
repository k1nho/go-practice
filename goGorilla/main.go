package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)




func main() {
	r := mux.NewRouter()

	r.HandleFunc(
		"/series/{title}/episode/{ep}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r) // getting all params using Vars method
			fmt.Printf("The title of the series is %v\n", vars["title"])
			fmt.Printf("The episode of the series is %v\n", vars["ep"])
		})

	// since we are using a mux and not the default we need to pass it to the listen function
	http.ListenAndServe(":80", r)

	// Gorilla mux also allows us to restrict the methods used in this case we can use Methods() to specify which routes should respond
	// only to an specific method for ex
	r.HandleFunc("api/bookmark/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Printf(" Bookmark the following title %v", vars["title"])
	}).Methods("POST")
	
	// we can also restrict hosts i.e (www.bookstore.com) or restrict schmes (http/https)
	func getAllCharacters(w http.http.ResponseWriter, r *http.Request) int {
		return 1;	
	}

	func getCharacter(w http.ResponseWriter, r *http.Request){
		
	}
	// To organize routes better we can have a common prefix 
	characterRouter := r.PathPrefix("/character").Subrouter()
	characterRouter.HandleFunc("/", getAllCharacters);
	characterRouter.HandleFunc("/{id}", getCharacter);

	// DB
	// db, err:= sql.Open("mysql", "root:pw@(localhost:3306)/dbname")

}
