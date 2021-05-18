package main

import (
	"log"
	"net/http"

	"link-logger/db"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatalln("cannot init db")
	}

	err = http.ListenAndServe("127.0.0.1:8080", controllers())
	if err != nil {
		log.Fatalf("\nserve error: %s\n", err.Error())
	}
}
