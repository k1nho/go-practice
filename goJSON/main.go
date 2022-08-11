package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func DecodeJson(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Fprintln(w, "%s %s is %d years old!", user.FirstName, user.LastName, user.Age)
}

func EncodeJson(w http.ResponseWriter, r *http.Request) {
	user := User{
		FirstName: "john",
		LastName:  "Doe",
		Age:       200,
	}

	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/decode", DecodeJson)
	http.HandleFunc("/encode", EncodeJson)

	http.ListenAndServe(":8000", nil)
}
