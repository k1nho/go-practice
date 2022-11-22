package main

import (
	"context"
	"fmt"
	"time"
)

type DTO struct {
  user string
  thirdPartyRes int
}

type ThirdPartyAns struct {
  val int
  err error
}


func main() {
  start := time.Now() 
  ctx := context.Background()
  res , err := restAPIFetch(ctx)
  if err != nil {
    fmt.Println("could not process data")
  }
  fmt.Printf("data: {user : %s, thirdPartyRes: %d}\n", res.user, res.thirdPartyRes)
  fmt.Printf("Roundtrip total time: %v", time.Since(start))
}


func restAPIFetch(c context.Context) (*DTO, error) {
  data := &DTO{}
  data.user = "Gary"
  ctx , cancel := context.WithTimeout(c, time.Millisecond*100)
  defer cancel()
  
  ch := make(chan ThirdPartyAns, 1)

  go func() {
    thirdPres, err := thirdPartyFetch()
    ch <- ThirdPartyAns{ val: thirdPres, err: err}
  }() 
 
  for {
    select {
      case <-ctx.Done():
        return data, fmt.Errorf("took longer than 100 secs to process third party service")

      case res := <-ch :
        data.thirdPartyRes = res.val
        return data, nil
    }
  }
  

}

func thirdPartyFetch() (int, error) {
  time.Sleep(time.Millisecond * 1000) // if the time to process this request is greater than 100 ms then we do not process
  return 1, nil
}
