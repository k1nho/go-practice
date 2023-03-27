package caching

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerCaching(t *testing.T) {
	server := CreateServer()
	testServer := httptest.CreateServer(http.HandlerFunc(server.Getuser()))

	reqs := 1000

	for i := 0; i < 1000; i++ {
		id := i%100 + 1
		reqURL := fmt.Sprintf("%s/?id=%d", testServer.URL, id)

		res, err := http.Get(reqURL)
		if err != nil {
			t.Error(err)
		}

		usr := &User
		if err := json.NewDecoder(res.Body).Decode(&usr); err != nil {
			t.Error(err)
		}

		fmt.Printf("%+v\n", usr)
		fmt.Println("Total DB requests: %d", server.DBcalls)
	}

}
