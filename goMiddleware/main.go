package main

import(
	"fmt"
	"net/http"
	"log"
	"time"
)

type Middleware(http.HandleFunc) http.HandleFunc

// Function to create a middleware 
func createNewMiddleware() Middleware{
	middleware := func(next http.HandleFunc) http.HandleFunc{

		//Function is called by the server eventually
		handler := func(w http.ResponseWriter, r *http.Request){

			// ... processing middleware

			//Call the next middleware in chain
			next(w, r)
		}

		return handler;
	}

	return middleware;
}


func Logging() Middleware{
	return func(f http.HandleFunc) http.HandleFunc{
		return func(w http.ResponseWriter, r * http.Request){

			start := time.Now();

			defer func(){log.Println(r.URL.Path, time.Since(start))}()
			f(w,r)
		}
	}
}

func Method(s string) Middleware{
	return func(f http.HandleFunc) http.HandleFunc{
		return func(w http.ResponseWriter, r * http.Request){

			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			f(w,r)
		}
	}
}

func Hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello!")
}


// applies all middlewares to an http.HandleFunc 
func Chain(f http.HandleFunc, middlewares ...Middleware) http.HandleFunc{
	for _, m := range middlewares{
		f = m(f)
	}
	return f
}

func main(){
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":80", nil)
}