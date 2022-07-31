package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserDetails struct {
	Name    string
	Email   string
	Content string
}

func main() {
	templ := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			templ.Execute(w, nil)
			return
		}

		userDetails := UserDetails{
			Name:    r.FormValue("name"),
			Email:   r.FormValue("email"),
			Content: r.FormValue("content"),
		}

		_ = userDetails
		fmt.Printf("Log details: %v user, %v email, %v Content", userDetails.Name, userDetails.Email, userDetails.Content)

		templ.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":80", nil)
}
