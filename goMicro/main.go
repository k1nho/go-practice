package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
  mService := CreateDogService("https://dog-api.kinduff.com/facts")  
  mService = CreateLoggerService(mService) 

  data, err := mService.GetDogFact(context.TODO())
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%+v\n", data)
}
