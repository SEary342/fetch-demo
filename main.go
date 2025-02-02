package main

import (
	"fmt"
	"log"
	"net/http"

	"fetch-demo/internal/api"
	"fetch-demo/internal/restapi"
)

func main() {
	server := restapi.NewServer()

	r := http.NewServeMux()

	h := api.HandlerFromMux(&server, r)
	addr := "0.0.0.0:8080"
	s := &http.Server{
		Handler: h,
		Addr:    addr,
	}
	fmt.Printf("App is running at http://%s", addr)

	log.Fatal(s.ListenAndServe())
}
