package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type DogFact struct {
  Fact []string  `json: "facts"`
}

type Service interface {
  GetDogFact(context.Context) (*DogFact, error)
}

type DogFactService struct{
  url string
}

func CreateDogService(url string) Service {
  return &DogFactService{
    url : url,
  }
}


func (service *DogFactService) GetDogFact(ctx context.Context) (*DogFact, error) {
  res, err := http.Get(service.url)
  if err != nil {
    return nil, err
  }
  
  defer res.Body.Close()

  fact := &DogFact{}
  if err := json.NewDecoder(res.Body).Decode(fact); err != nil {
    return nil, err
  }

  return fact, nil

}
