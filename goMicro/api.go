package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type  ApiServer struct {
  mService Service
}

func CreateApiServer(service Service) *ApiServer {
  return &ApiServer{ 
    mService: service, 
  }
}

func (service *ApiServer) Start(Addr string) error {
  http.HandleFunc("/", service.handlerGetDogFact)
  return http.ListenAndServe(Addr, nil)
}

func (service *ApiServer) handlerGetDogFact(w http.ResponseWriter, r *http.Request) {
  data, err := service.mService.GetDogFact(context.Background())
  if err != nil {
    writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
  }

  writeJSON(w, http.StatusOK, data)
}

func writeJSON(w http.ResponseWriter , s int, data any) error {
  w.WriteHeader(s)
  w.Header().Add("Content-Type", "application/json")
  return json.NewEncoder(w).Encode(data)
}
