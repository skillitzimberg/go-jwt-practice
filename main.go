package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	router := CreateRouter()
	InitializeRoutes(router)
		
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on :3000...")
	log.Fatal(srv.ListenAndServe())
}
