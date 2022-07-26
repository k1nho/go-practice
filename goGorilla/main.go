package main

import "net/http"

func main() {
	r := mux.newRouter()

	r.handleFunc(
		"/series/{title}/episode/{ep}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
		})

}
