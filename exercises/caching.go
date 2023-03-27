package caching

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID       int
	Username string
	Likes    int
}

type Server struct {
	DB      map[int]*User
	Cache   map[int]*User
	DBcalls int
}

func CreateServer() *Server {
	database := make(map[int]*User)

	for i := 0; i < 100; i++ {
		database[i+1] = &User{
			ID:       i + 1,
			Username: fmt.Sprintf("f_user_%d", i+1),
			Likes:    i + 10,
		}
	}

	return &Server{
		DB:      database,
		Cache:   make(map[int]*User),
		DBcalls: 0,
	}
}

func (s *Server) RetrieveFromCache(userID int) (*User, bool) {
	user, ok := s.Cache[userID]
	return user, ok
}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	paramID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Parameter given is not valid"))
		return
	}

	// hit Cache
	data, ok := s.RetrieveFromCache(id)
	if ok {
		json.NewEncoder(w).Encode(data)
		return
	}

	data, ok = s.DB[id]

	if !ok {
		panic("could not find the user")
	}
	s.DBcalls++

	s.Cache[id] = data

	err = json.NewEncoder(w).Encode(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Parameter given is not valid"))
		return
	}

}

func main() {

}
