package main

import (
	"context"
	"fmt"
	"time"
)

type Logger struct {
  next Service
}

func CreateLoggerService(next Service) Service {
  return &Logger{
    next: next,
  }
}

func (service *Logger) GetDogFact(ctx context.Context) (fact *DogFact, err error) {
  defer func(start time.Time){
    fmt.Printf("dog_fact= %v, err = %v, time=%v", fact, err, time.Since(start))
  }(time.Now())
  
  return service.next.GetDogFact(ctx)
}


